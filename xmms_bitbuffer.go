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
import (
	"unsafe"
)

type bitBuffer struct {
	Value
}

func NewBitBuffer() BitBuffer {
	b := new(bitBuffer)
	b.data = C.xmmsv_new_bitbuffer()

	var B BitBuffer = b
	return B
}

func NewBitBufferReadonly(v []byte, length int) BitBufferReadonly {
	b := new(bitBuffer)
	b.data = C.xmmsv_new_bitbuffer_ro((*C.uchar)(unsafe.Pointer(&v[0])), (C.int)(length))

	var B BitBufferReadonly = b
	return B
}

type BitBuffer interface {
	ValueNone
}

type BitBufferReadonly interface {
	ValueNone
}
