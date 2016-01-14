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

// XmmsList
type List struct {
	Value
}

func NewList() *List {
	l := new(List)
	l.data = C.xmmsv_new_list()
	return l
}
