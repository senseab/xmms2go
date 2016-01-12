/*
Xmms2Go -- A golang binding to libxmmsclient.
Copyright (C) 2016  TonyChyi <tonychee1989@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/*
Xmms2Go is a Go binding for libxmmsclient.
It's easy to use and friendly for go developers.
Just import me to use.

    package main
    import(
        "github.com/tonychee7000/xmms2go"
        "os"
        "fmt"
    )

    func main(){
        x := xmms2go.NewXmms2Client("test")
        err := x.Connect(os.Getenv("XMMS_PATH"))
        // According to the documents of xmms2, some resources
        // should be released. So Unref() is necessary.
        defer x.Unref()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        err := x.Play()
        if err != nil {
            fmt.Println(err)
        }
    }

*/
package xmms2go

/*
#cgo CFLAGS: -I/usr/include/xmms2
#cgo LDFLAGS: -L/usr/lib -lxmmsclient
#include <xmmsclient/xmmsclient.h>
#include <xmmsc/xmmsv.h>
*/
import "C"
import (
	"errors"
	"fmt"
)

// A class of xmmsclient
type Xmms2Client struct {
	Connection  *C.xmmsc_connection_t
	result      *C.xmmsc_result_t
	returnValue *C.xmmsv_t
	errorBuff   *C.char
}

// Make new xmmsclient instance.
func NewXmms2Client(clientName string) (*Xmms2Client, error) {
	x := new(Xmms2Client)
	x.Connection = C.xmmsc_init(C.CString(clientName))
	if x.Connection == nil {
		return nil, errors.New("Client init failed")
	}
	return x, nil
}

/*
Connect to xmms server, both tcp or unix socket are works.
*/
func (x *Xmms2Client) Connect(url string) error {
	r := C.xmmsc_connect(x.Connection, C.CString(url))
	if r == 0 {
		errInfo := C.GoString(C.xmmsc_get_last_error(x.Connection))
		return errors.New(fmt.Sprintf("Connection failed: %s", errInfo))
	}
	return nil
}

// Start playback.
func (x *Xmms2Client) Play() error {
	x.result = C.xmmsc_playback_start(x.Connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback start returned error: %s")
}

// Pause playback.
func (x *Xmms2Client) Pause() error {
	x.result = C.xmmsc_playback_pause(x.Connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback pause returned error: %s")
}

// Stop playback.
func (x *Xmms2Client) Stop() error {
	x.result = C.xmmsc_playback_stop(x.Connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback stop returned error: %s")
}

// Stop decoding of current song.
func (x *Xmms2Client) Tickle() error {
	x.result = C.xmmsc_playback_tickle(x.Connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue = C.xmmsc_result_get_value(x.result)
	return x.checkError("Playback tickle returned error: %s")

}

// Get Current ID
func (x *Xmms2Client) CurrentID() (int, error) {
	x.result = C.xmmsc_playback_current_id(x.Connection)
	C.xmmsc_result_wait(x.result)
	x.returnValue = C.xmmsc_result_get_value(x.result)
	err := x.checkError("Get Current ID failed: %s")
	if err != nil {
		return -1, err
	}
	return x.GetInt()
}

// Get integer form return value
func (x *Xmms2Client) GetInt() (int, error) {
	var i C.int32_t
	if int(C.xmmsv_get_int(x.returnValue, &i)) == 0 {
		return -1, errors.New("Get Current ID failed")
	}
	return int(i), nil
}

/*
You SHOULD use this when you quit.

    x := NewXmms2Client("test")
    x.Connect("somewhere")
    defer x.Unref()
    os.Exit(0)

*/
func (x *Xmms2Client) Unref() {
	C.xmmsc_result_unref(x.result)
	C.xmmsc_unref(x.Connection)
}

func (x *Xmms2Client) checkError(hintString string) error {
	if int(C.xmmsv_is_error(x.returnValue)) != 0 &&
		int(C.xmmsv_get_error(x.returnValue, &x.errorBuff)) != 0 {
		return errors.New(fmt.Sprintf(
			hintString, C.GoString(x.errorBuff),
		))
	}
	return nil
}
