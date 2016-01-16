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
type dict struct {
	Value
}

func NewDict() Dict {
	d := new(dict)
	d.data = C.xmmsv_new_dict()

	var D Dict = d
	return D
}

type Dict interface {
	ValueNone
}
