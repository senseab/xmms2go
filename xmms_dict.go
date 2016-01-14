package xmms2go

/*
#ifndef XMMS2GO
#define XMMS2GO
#cgo CFLAGS: -I/usr/include/xmms2
#cgo LDFLAGS: -lxmmsclient
#include <xmmsc/xmmsv.h>
#include <malloc.h>

#endif
*/
import "C"

// XmmsDict
type Dict struct {
	Value
}

func NewDict() *Dict {
	d := new(Dict)
	d.data = C.xmmsv_new_dict()
	return d
}
