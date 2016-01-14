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
	"errors"
	"unsafe"
)

// *C.xmmsv_t
type Value struct {
	data *C.xmmsv_t
}

func NewValueFromNone() ValueNone {
	x := new(Value)
	x.data = C.xmmsv_new_none()
    var vn ValueNone = x
	return vn
}

func NewValueFromError(e error) ValueError {
	x := new(Value)
	cErrInfo := C.CString(e.Error())
	defer C.free(unsafe.Pointer(cErrInfo))

	x.data = C.xmmsv_new_error(cErrInfo)
    var ve ValueError = x
	return ve
}

func NewValueFromInt64(i int64) ValueInt64 {
	x := new(Value)
	x.data = C.xmmsv_new_int(C.int64_t(i))
    var vi ValueInt64 = x
	return vi
}

func NewValueFromInt32(i int32) ValueInt32 {
	x := new(Value)
	x.data = C.xmmsv_new_int(C.int64_t(i))
    var vi ValueInt32 = x
	return vi
}

func NewValueFromFloat64(f float64) ValueFloat64{
    x := new(Value)
	x.data = C.xmmsv_new_float(C.float(f))
    var vf ValueFloat64 = x
	return vf
}

func NewValueFromFloat32(f float32) ValueFloat32 {
	x := new(Value)
	x.data = C.xmmsv_new_float(C.float(f))
    var vf ValueFloat32 = x
	return vf
}

func NewValueFromString(s string) ValueString {
	x := new(Value)
	cString := C.CString(s)
	defer C.free(unsafe.Pointer(cString))

	x.data = C.xmmsv_new_string(cString)
    var vs ValueString = x
	return vs
}

func NewValueFromBytes(b []byte) ValueBytes {
    length := len(b)
	x := new(Value)
	d := (*C.uchar)(unsafe.Pointer(&b[0]))
	//defer C.free(unsafe.Pointer(d)) // convert from Go, free memory necessary?

	x.data = C.xmmsv_new_bin(d, C.uint(length))
    var vb ValueBytes = x
	return vb
}

func NewValueFromCopyValue(v *Value) *Value {
	x := new(Value)
	xmmsvT := v.Export()
	defer C.free(unsafe.Pointer(xmmsvT))

	x.data = C.xmmsv_copy(xmmsvT)
	return x
}

func NewValueFromRef(v *Value) *Value {
	x := new(Value)
	xmmsvT := v.Export()
	defer C.free(unsafe.Pointer(xmmsvT))

	x.data = C.xmmsv_ref(xmmsvT)
	return x
}

func (x *Value) Export() *C.xmmsv_t {
	return x.data
}

func (x *Value) Unref() {
	C.xmmsv_unref(x.data)
}

func (x *Value) GetError() (error, error) {
	var cErrInfo *C.char
	defer C.free(unsafe.Pointer(cErrInfo))
    if x.IsError() == false {
        return nil, errors.New("Not an error type")
    }
	r := int(C.xmmsv_get_error(x.data, &cErrInfo))
	if r == 0 {
		return nil, errors.New("Parse type error failed")
	}
	return errors.New(C.GoString(cErrInfo)), nil
}

func (x *Value) GetInt32() (int32, error) {
	var i C.int32_t
	r := int(C.xmmsv_get_int32(x.data, &i))
	if r == 0 {
		return 0, errors.New("Parse type int32 failed")
	}
	return int32(i), nil
}

func (x *Value) GetInt64() (int64, error) {
	var i C.int64_t
	r := int(C.xmmsv_get_int64(x.data, &i))
	if r == 0 {
		return 0, errors.New("Parse type int64 failed")
	}
	return int64(i), nil
}

func (x *Value) GetFloat32() (float32, error) {
    f, err := x.getFloat()
	return float32(f), err
}

func (x *Value) GetFloat64() (float64, error) {
    f, err := x.getFloat()
	return float64(f), err
}

func (x *Value) getFloat() (C.float, error) {
	var f C.float
	r := int(C.xmmsv_get_float(x.data, &f))
	if r == 0 {
		return 0, errors.New("Parse type float failed")
	}
    return f, nil
}

func (x *Value) GetBytes() ([]byte, error) {
	var b *C.uchar
	var l C.uint
	defer C.free(unsafe.Pointer(b))
	r := int(C.xmmsv_get_bin(x.data, &b, &l))
	if r == 0 {
		return nil, errors.New("Parse type bytes failed")
	}
	return C.GoBytes(unsafe.Pointer(b), C.int(l)), nil

}

func (x *Value) GetString() (string, error){
    var s *C.char
    defer C.free(unsafe.Pointer(s))

    r := int(C.xmmsv_get_string(x.data, &s))
    if r == 0 {
        return "", errors.New("Parse type string failed")
    }
    return C.GoString(s), nil
}

// Okay, we need to implement the collection type.
//func (x *Value) GetCollection() (*Collection, error)

func (x *Value) IsError() bool {
	r := int(C.xmmsv_is_error(x.data))
	if r == 1 {
		return true
	}
	return false
}

type ValueNone interface{
    Unref()
}

type ValueError interface{
    Unref()
    IsError() bool
    GetError() (error, error)
}

type ValueInt64 interface {
    Unref()
    GetInt64() (int64, error)
}

type ValueInt32 interface {
    Unref()
    GetInt32()(int32, error)
}

type ValueFloat64 interface {
    Unref()
    GetFloat64()(float64, error)
}

type ValueFloat32 interface {
    Unref()
    GetFloat32()(float32, error)
}

type ValueString interface {
    Unref()
    GetString() (string, error)
}

type ValueBytes interface {
    Unref()
    GetBytes() ([]byte, error)
}
