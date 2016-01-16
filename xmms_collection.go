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

type collection struct {
	Value
}

func NewCollection(_type int) Collection {
	c := new(collection)
	c.data = C.xmmsv_new_coll((C.xmmsv_coll_type_t)(_type))
	var C Collection = c
	return C
}

type Collection interface {
	ValueNone
}
