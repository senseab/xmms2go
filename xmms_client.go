package xmms2go

/*
#ifndef XMMS2GO
#define XMMS2GO
#cgo CFLAGS: -I/usr/include/xmms2
#cgo LDFLAGS: -lxmmsclient
#include <xmmsclient/xmmsclient.h>
#include <malloc.h>

void callLockFunc(void*);
void callUnlockFunc(void*);
void callDisconnectFunc(void*);
void callUserDataFreeFunc(void*);
void callIONeedOutCallbackFunc(int, void*);

static int macro_xmmsc_result_iserror(xmmsc_result_t *val) {
    return xmmsc_result_iserror(val);
}

static void xmmsc_lock_set_wrapper(
    xmmsc_connection_t *conn,
    void *lock
) {
    xmmsc_lock_set(conn, lock, callLockFunc, callUnlockFunc);
}

static void xmmsc_disconnect_callback_set_wrapper(
    xmmsc_connection_t *conn,
    void *userdata
) {
    xmmsc_disconnect_callback_set(conn, callDisconnectFunc, userdata);
}

static void xmmsc_disconnect_callback_set_full_wrapper(
    xmmsc_connection_t *conn,
    void *userdata
) {
    xmmsc_disconnect_callback_set_full(
        conn, callDisconnectFunc,
        userdata,
        callUserDataFreeFunc
    );
}

static void xmmsc_io_need_out_callback_set_wrapper(
    xmmsc_connection_t *conn,
    void *userdata
){
    xmmsc_io_need_out_callback_set(conn, callIONeedOutCallbackFunc, userdata);
}

static void xmmsc_io_need_out_callback_set_full_wrapper(
    xmmsc_connection_t *conn,
    void *userdata
){
    xmmsc_io_need_out_callback_set_full(
        conn, callIONeedOutCallbackFunc,
        userdata, callUserDataFreeFunc
    );
}


#endif
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
	//"sync"
)

const (
	ResultClassDefault = iota
	ResultClassSignal
	ResultClassBroadcast
)

/*
LockFunc means origin
    void (*lockfunc)(void *)
at xmms_client.h: 50
*/
type LockFunc func(interface{})

/*
UnlockFunc means origin
    void (*unlockfunc)(void *)
at xmms_client.h: 50
*/
type UnlockFunc func(interface{})

/*
DisconnectFunc means origin
    typedef void (*xmmsc_disconnect_func_t) (void *user_data);

To convert interface{} to a struct, use:
    foo := (*Foo)(v.(unsafe.Pointer))
*/
type DisconnectFunc func(interface{})

/*
UserDataFreeFunc means origin
    typedef void (*xmmsc_user_data_free_func_t) (void *user_data);

To convert interface{} to a struct, use:
    foo := (*Foo)(v.(unsafe.Pointer))
*/
type UserDataFreeFunc func(interface{})

/*
IONeedOutCallbackFunc means origin
    typedef void (*xmmsc_io_need_out_call_back_func_t) (int, void*);

To convert interface{} to a struct, use:
    foo := (*Foo)(v.(unsafe.Pointer))
*/
type IONeedOutCallbackFunc func(int, interface{})

var lockFunc LockFunc
var unlockFunc UnlockFunc
var disconnectFunc DisconnectFunc
var userDataFreeFunc UserDataFreeFunc
var ioNeedOutCallbackFunc IONeedOutCallbackFunc

//export callLockFunc
func callLockFunc(p unsafe.Pointer) {
	v := (interface{})(p)
	lockFunc(v)
}

//export callUnlockFunc
func callUnlockFunc(p unsafe.Pointer) {
	v := (interface{})(p)
	unlockFunc(v)
}

//export callDisconnectFunc
func callDisconnectFunc(p unsafe.Pointer) {
	v := (interface{})(p)
	disconnectFunc(v)
}

//export callUserDataFreeFunc
func callUserDataFreeFunc(p unsafe.Pointer) {
	v := (interface{})(p)
	userDataFreeFunc(v)
}

//export callIONeedOutCallbackFunc
func callIONeedOutCallbackFunc(i C.int, p unsafe.Pointer) {
	v := (interface{})(p)
	ioNeedOutCallbackFunc(int(i), v)
}

func GetUserConfDir(b []byte) string {
	pb := (*C.char)((unsafe.Pointer)(&b[0]))
	length := len(b)
	c := C.xmmsc_userconfdir_get(pb, (C.int)(length))
	defer C.free((unsafe.Pointer)(c))
	return C.GoString(c)
}

type Connector struct {
	connection *C.xmmsc_connection_t
}

func NewConnector(clientName string) (*Connector, error) {
	x := new(Connector)
	cClientName := C.CString(clientName)
	defer C.free(unsafe.Pointer(cClientName))
	x.connection = C.xmmsc_init(cClientName)
	if x.connection == nil {
		return nil, fmt.Errorf("Client init failed")
	}
	return x, nil
}

func (c *Connector) Connect(url string) error {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))
	r := C.xmmsc_connect(c.connection, cUrl)
	if r == 0 {
		return fmt.Errorf("Connection failed %v", c.GetLastError())
	}
	return nil
}

func (c *Connector) UnRef() {
	C.xmmsc_unref(c.export())
}

// Actually, Use native sync.Mutex is better
func (c *Connector) LockSet(
	lock interface{},
	lockfunc LockFunc,
	unlockfunc UnlockFunc,
) {
	lockFunc = lockfunc
	unlockFunc = unlockfunc

	C.xmmsc_lock_set_wrapper(c.connection, lock)
}

// When use this in parallel, a lock is needed.
func (c *Connector) DisconnectCallBackSet(
	disconnectfunc DisconnectFunc,
	userdata interface{},
) {
	disconnectFunc = disconnectfunc
	C.xmmsc_disconnect_callback_set_wrapper(c.connection, userdata)
}

// When use this in parallel, a lock is needed.
func (c *Connector) DisconnectCallBackSetFull(
	disconnectfunc DisconnectFunc,
	userdata interface{},
	userdatafreefunc UserDataFreeFunc,
) {
	disconnectFunc = disconnectfunc
	userDataFreeFunc = userdatafreefunc
	C.xmmsc_disconnect_callback_set_full_wrapper(
		c.connection,
		userdata,
	)
}

// When use this in parallel, a lock is needed.
func (c *Connector) IONeedOutCallbackSet(
	ioneedoutcallback IONeedOutCallbackFunc,
	userdata interface{},
) {
	ioNeedOutCallbackFunc = ioneedoutcallback
	C.xmmsc_io_need_out_callback_set_wrapper(
		c.connection,
		userdata,
	)
}

// When use this in parallel, a lock is needed.
func (c *Connector) IONeedOutCallbackSetFull(
	ioneedoutcallback IONeedOutCallbackFunc,
	userdata interface{},
	userdatafreefunc UserDataFreeFunc,
) {
	ioNeedOutCallbackFunc = ioneedoutcallback
	userDataFreeFunc = userdatafreefunc
	C.xmmsc_io_need_out_callback_set_full_wrapper(
		c.connection,
		userdata,
	)
}

func (c *Connector) IODisconnect() {
	C.xmmsc_io_disconnect(c.connection)
}

func (c *Connector) IOWantOut() error {
	r := C.xmmsc_io_want_out(c.connection)
	if int(r) == 0 {
		return fmt.Errorf("IO Want Out failed")
	}
	return nil
}

func (c *Connector) IOOutHandle() error {
	r := C.xmmsc_io_out_handle(c.export())
	if int(r) == 0 {
		return fmt.Errorf("IO Out Handle failed")
	}
	return nil
}

func (c *Connector) IOInHandle() error {
	r := C.xmmsc_io_in_handle(c.export())
	if int(r) == 0 {
		return fmt.Errorf("IO In Handle failed")
	}
	return nil
}

func (c *Connector) IOFdGet() error {
	r := C.xmmsc_io_fd_get(c.export())
	if int(r) == 0 {
		return fmt.Errorf("IO Fd Get failed")
	}
	return nil
}

func (c *Connector) Quit() *Result {
	r := new(Result)
	r.result = C.xmmsc_quit(c.export())
	return r
}

func (c *Connector) BroadCastQuit() *Result {
	r := new(Result)
	r.result = C.xmmsc_broadcast_quit(c.export())
	return r
}

func (c *Connector) GetLastError() error {
	cErrInfo := C.xmmsc_get_last_error(c.connection)
	defer C.free(unsafe.Pointer(cErrInfo))
	return fmt.Errorf(C.GoString(cErrInfo))
}

func (c *Connector) export() *C.xmmsc_connection_t {
	return c.connection
}

type Result struct {
	result *C.xmmsc_result_t
}

func NewResult() *Result {
	return new(Result)
}

func (r *Result) GetClass() int {
	return int(C.xmmsc_result_get_class(r.export()))
}

func (r *Result) Disconnect() {
	C.xmmsc_result_disconnect(r.export())
}

func (r *Result) UnRef() {
	C.xmmsc_result_unref(r.export())
}

func (r *Result) Wait() {
	C.xmmsc_result_wait(r.export())
}

func (r *Result) GetValue() *Value {
	v := new(Value)
	v.data = C.xmmsc_result_get_value(r.export())
	return v
}

// Dummy
func (r *Result) NotifierSetDefault() {
}

// Dummy
func (r *Result) NotifierSetDefaultFull() {
}

// Dummy
func (r *Result) NotifierSetRaw() {
}

// Dummy
func (r *Result) NotifierSetRawFull() {
}

// Dummy
func (r *Result) NotifierSetC2C() {
}

// Dummy
func (r *Result) NotifierSetC2CFull() {
}

func (r *Result) export() *C.xmmsc_result_t {
	return r.result
}

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
	defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_start(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback start returned error: %s")
}

// Pause playback.
func (x *Xmms2Client) Pause() error {
	defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_pause(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback pause returned error: %s")
}

// Stop playback.
func (x *Xmms2Client) Stop() error {
	defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_stop(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback stop returned error: %s")
}

// Stop decoding of current song.
func (x *Xmms2Client) Tickle() error {
	defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_tickle(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback tickle returned error: %s")

}

// Get Current ID. If failed, return -1 and error info
func (x *Xmms2Client) CurrentID() (int, error) {
	defer x.returnValue.Unref()
	x.result = C.xmmsc_playback_current_id(x.connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	err := x.checkError("Get current ID failed: %s")
	if err != nil {
		return -1, err
	}
	i, err := x.returnValue.GetInt32()
	return int(i), err
}

// --- Medialib operations ---

// Get medialib info
func (x *Xmms2Client) MediaLibInfo(id int) (map[string]interface{}, error) {
	defer x.ResultUnref()
	defer x.returnValue.Unref()
	x.result = C.xmmsc_medialib_get_info(x.connection, C.int(id))
	C.xmmsc_result_wait(x.result)
	x.returnValue.data = C.xmmsc_result_get_value(x.result)
	err := x.checkError("Get medialib info failed: %s")
	if err != nil {
		return nil, err
	}

	m, err := x.returnValue.GetDict()
	defer m.Unref()
	if err != nil {
		return nil, err
	}
	return m.ToMap()
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
	C.xmmsc_unref(x.connection)
}

// --- Private operations ---

func (x *Xmms2Client) checkError(hintString string) error {
	if int(C.macro_xmmsc_result_iserror(x.result)) != 0 {
		errorBuff := C.xmmsc_get_last_error(x.connection)
		defer C.free(unsafe.Pointer(errorBuff))
		return errors.New(fmt.Sprintf(
			hintString, C.GoString(errorBuff),
		))
	}
	return nil
}
