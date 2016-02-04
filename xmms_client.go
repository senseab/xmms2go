package xmms2go

/*
#ifndef XMMS2GO
#define XMMS2GO
#cgo CFLAGS: -I/usr/include/xmms2
#cgo LDFLAGS: -lxmmsclient
#include <xmmsclient/xmmsclient.h>
#include <malloc.h>

static int macro_xmmsc_result_iserror(xmmsc_result_t *val) {
    return xmmsc_result_iserror(val);
}
#endif
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)
// A class of xmmsclient
type Xmms2Client struct {
	connection  *C.xmmsc_connection_t
	result      *C.xmmsc_result_t
	returnValue *Value
}

// Make new xmmsclient instance.
func NewXmms2Client(clientName string) (*Xmms2Client, error) {
	x := new(Xmms2Client)
	cClientName := C.CString(clientName)
	defer C.free(unsafe.Pointer(cClientName))
	x.connection = C.xmmsc_init(cClientName)
    x.returnValue = new(Value)
	if x.connection == nil {
		return nil, errors.New("Client init failed")
	}
	return x, nil
}

/*
Connect to xmms server, both tcp or unix socket are works.

    x = NewXmms2Client("test")
    x.Connect("unix://somewhere")
    x.Connect("tcp://somewhere")

*/
func (x *Xmms2Client) Connect(url string) error {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))
	r := C.xmmsc_connect(x.connection, cUrl)
	if r == 0 {
		errInfo := C.GoString(C.xmmsc_get_last_error(x.connection))
		return errors.New(fmt.Sprintf("Connection failed: %s", errInfo))
	}
	return nil
}

// --- Playback operations ---

// Start playback.
func (x *Xmms2Client) Play() error {
	defer x.ResultUnref()
    defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_start(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback start returned error: %s")
}

// Pause playback.
func (x *Xmms2Client) Pause() error {
	defer x.ResultUnref()
    defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_pause(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback pause returned error: %s")
}

// Stop playback.
func (x *Xmms2Client) Stop() error {
	defer x.ResultUnref()
    defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_stop(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback stop returned error: %s")
}

// Stop decoding of current song.
func (x *Xmms2Client) Tickle() error {
	defer x.ResultUnref()
    defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_tickle(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback tickle returned error: %s")

}

// Get Current ID. If failed, return -1 and error info
func (x *Xmms2Client) CurrentID() (int, error) {
	defer x.ResultUnref()
    defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_current_id(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	err := x.checkError("Get current ID failed: %s")
	if err != nil {
		return -1, err
	}
	return getInt(x.returnValue)
}

// --- Medialib operations ---

// Get medialib info
func (x *Xmms2Client) MediaLibInfo(id int) (map[string]interface{}, error) {
	defer x.ResultUnref()
    defer x.returnValue.Unref()
	m := make(map[string]interface{})
	x.result = C.xmmsc_medialib_get_info(x.connection, C.int(id))
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	err := x.checkError("Get medialib info failed: %s")
	if err != nil {
		return nil, err
	}
	// TODO: new dict func
	return m, nil
}

// --- Clean operations ---

// Every operation is done, clear memeory is needed.
func (x *Xmms2Client) ResultUnref() {
	C.xmmsc_result_unref(x.result)
}

/*
You SHOULD use this when you quit.

    x := NewXmms2Client("test")
    x.Connect("somewhere")
    defer x.Unref()
    defer x.ResultUnref()
    os.Exit(0)

*/
func (x *Xmms2Client) Unref() {
	C.xmmsc_unref(x.Connection)
}

// --- Private operations ---

func (x *Xmms2Client) checkError(hintString string) error {
	if int(C.macro_xmmsc_result_iserror(x.result)) != 0 {
        errorBuff := C.xmmsc_get_last_error(x.Connection)
	    defer C.free(unsafe.Pointer(errorBuff))
		return errors.New(fmt.Sprintf(
			hintString, C.GoString(errorBuff),
		))
	}
	return nil
}

// deprecated
// --- Data operations ---

// Get integer form xmmsv_t
func getInt(x *C.xmmsv_t) (int, error) {
	var i C.int32_t
	if int(C.xmmsv_get_int(x, &i)) == 0 {
		return -1, errors.New("Parse int failed")
	}
	return int(i), nil
}

// Get string from xmmsv_t
func getString(x *C.xmmsv_t) (string, error) {
	var s *C.char
	if int(C.xmmsv_get_string(x, &s)) == 0 {
		return "", errors.New("Parse string failed")
	}
	return C.GoString(s), nil
}
