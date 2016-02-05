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
	"fmt"
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

func (b *bitBuffer) GetBits(bits int) (int64, error) {
	var res C.int64_t
	r := C.xmmsv_bitbuffer_get_bits(b.data, C.int(bits), &res)
	if int(r) == 0 {
		return 0, fmt.Errorf("Get bits for length=%d failed", bits)
	}
	return int64(res), nil
}

func (b *bitBuffer) GetData(length int) ([]byte, error) {
	var res *C.uchar
	res = (*C.uchar)(C.malloc(C.size_t(length)))
	defer C.free(unsafe.Pointer(res))
	r := C.xmmsv_bitbuffer_get_data(b.data, res, C.int(length))
	if int(r) == 0 {
		return nil, fmt.Errorf("Get data for length=%d failed", length)
	}
	return C.GoBytes(unsafe.Pointer(res), C.int(length)), nil

}

func (b *bitBuffer) PutBits(bits int, data int64) error {
	r := C.xmmsv_bitbuffer_put_bits(b.data, C.int(bits), C.int64_t(data))
	if int(r) == 0 {
		return fmt.Errorf("Put %d bits with %d failed", bits, data)
	}
	return nil
}

func (b *bitBuffer) PutBitsAt(bits int, data int64, offset int) error {
	r := C.xmmsv_bitbuffer_put_bits_at(b.data, C.int(bits), C.int64_t(data), C.int(offset))
	if int(r) == 0 {
		return fmt.Errorf("Put %d bits with %d to 0x%h, failed", bits, data, offset)
	}
	return nil
}

func (b *bitBuffer) PutData(bb []byte) error {
	length := len(bb)
	cB := (*C.uchar)(unsafe.Pointer(&bb[0]))
	r := C.xmmsv_bitbuffer_put_data(b.data, cB, C.int(length))
	if int(r) == 0 {
		return fmt.Errorf("Put data with %d bytes failed", length)
	}
	return nil
}

func (b *bitBuffer) Align() error {
	r := C.xmmsv_bitbuffer_align(b.data)
	if int(r) == 0 {
		return fmt.Errorf("Align buffer failed")
	}
	return nil
}

func (b *bitBuffer) Goto(pos int) error {
	r := C.xmmsv_bitbuffer_goto(b.data, C.int(pos))
	if int(r) == 0 {
		return fmt.Errorf("Goto pos 0x%h failed", pos)
	}
	return nil
}

func (b *bitBuffer) Pos() int {
	return int(C.xmmsv_bitbuffer_pos(b.data))
}

func (b *bitBuffer) Rewind() error {
	r := C.xmmsv_bitbuffer_rewind(b.data)
	if int(r) == 0 {
		return fmt.Errorf("Rewind buffer failed")
	}
	return nil
}

func (b *bitBuffer) End() error {
	r := C.xmmsv_bitbuffer_end(b.data)
	if int(r) == 0 {
		return fmt.Errorf("End buffer failed")
	}
	return nil
}

func (b *bitBuffer) Len() int {
	return int(C.xmmsv_bitbuffer_len(b.data))
}

func (b *bitBuffer) GetBuffer() ([]byte, error) {
	var bb *C.uchar
	var length C.uint
	r := C.xmmsv_get_bitbuffer(b.data, &bb, &length)
	if int(r) == 0 {
		return nil, fmt.Errorf("Get buffer failed")
	} else {
		defer C.free(unsafe.Pointer(bb))
	}
	return C.GoBytes(unsafe.Pointer(bb), C.int(length)), nil
}


// Use this BitBuffer interface to avoid native value methods.
type BitBuffer interface {
	ValueNone
	GetBits(bits int) (int64, error)
	GetData(length int) ([]byte, error)
	PutBits(bits int, data int64) error
	PutBitsAt(bits int, data int64, offset int) error
	PutData(b []byte) error
	Align() error
	Goto(pos int) error
	Pos() int
	Rewind() error
	End() error
	Len() int
	GetBuffer() ([]byte, error)
}

type BitBufferReadonly interface {
	ValueNone
	GetBits(bits int) (int64, error)
	GetData(length int) ([]byte, error)
	Goto(pos int) error
	Pos() int
	Rewind() error
	End() error
	Len() int
	GetBuffer() ([]byte, error)
}
