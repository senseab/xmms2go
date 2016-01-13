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
    "errors"
)

// *C.xmmsv_t
type Value struct {
	data *C.xmmsv_t
}

func NewValueFromNone() *Value {
    x := new(Value)
    x.data = C.xmmsv_new_none()
    return x
}

func NewValueFromError(e error) *Value{
    x := new(Value)
    cErrInfo := C.CString(e.Error())
    defer C.free(unsafe.Pointer(cErrInfo))

    x.data = C.xmmsv_new_error(cErrInfo)
    return x
}

func NewValueFromInt64(i int64) *Value{
    x := new(Value)
    x.data = C.xmmsv_new_int(C.int64_t(i))
    return x
}

func NewValueFromFloat(f float32) *Value{
    x := new(Value)
    x.data = C.xmmsv_new_float(C.float(f))
    return x
}

func NewValueFromString(s string) *Value{
    x := new(Value)
    cString := C.CString(s)
    defer C.free(unsafe.Pointer(cString))

    x.data = C.xmmsv_new_string(cString)
    return x
}

func NewValueFromBytes(b []byte, length uint) *Value{
    x := new(Value)
    d := (*C.uchar)(unsafe.Pointer(&b[0]))
    defer C.free(unsafe.Pointer(d))

    x.data = C.xmmsv_new_bin(d, C.uint(length))
    return x
}

func NewValueFromCopyValue(v *Value) *Value{
    x := new(Value)
    xmmsvT := v.Export()
    defer C.free(unsafe.Pointer(xmmsvT))

    x.data = C.xmmsv_copy(xmmsvT)
    return x
}

func NewValueFromRef(v *Value) *Value{
    x := new(Value)
    xmmsvT := v.Export()
    defer C.free(unsafe.Pointer(xmmsvT))

    x.Import(xmmsvT)
    return x
}

func (x *Value) Import(v *C.xmmsv_t){
    x.data = C.xmmsv_ref(v)
}

func (x *Value) Export() *C.xmmsv_t{
    return x.data
}

func (x *Value) Unref() {
	C.xmmsv_unref(x.data)
}

func (x *Value) GetError() (error, error){
    var cErrInfo *C.char
    defer C.free(unsafe.Pointer(cErrInfo))
    r := int(C.xmmsv_get_error(x.data, &cErrInfo))
    if r == 0{
        return nil, errors.New("Parse type error failed")
    }
    return errors.New(C.GoString(cErrInfo)), nil
}

func (x *Value) GetInt32() (int32, error){
    var i C.int32_t
    r := int(C.xmmsv_get_int32(x.data, &i))
    if r == 0 {
        return 0, errors.New("Parse type int32 failed")
    }
    return int32(i), nil
}

func (x *Value) GetInt64() (int64, error){
    var i C.int64_t
    r := int(C.xmmsv_get_int64(x.data, &i))
    if r == 0 {
        return 0, errors.New("Parse type int64 failed")
    }
    return int64(i), nil
}

func (x *Value) GetFloat() (float32, error){
    var f C.float
    r := int(C.xmmsv_get_float(x.data, &f))
    if r == 0{
        return 0, errors.New("Parse type float failed")
    }
    return float32(f), nil
}

func (x *Value) GetBytes(length uint) ([]byte, error){
    var b *C.uchar
    defer C.free(unsafe.Pointer(b))
    cLength := C.uint(length)
    r := int(C.xmmsv_get_bin(x.data, &b, &cLength))
    if r == 0{
        return nil, errors.New("Parse type bytes failed")
    }
    return C.GoBytes(unsafe.Pointer(b), C.int(length)), nil

}

// Okay, we need to implement the collection type.
//func (x *Value) GetCollection() (*Collection, error)

func (x *Value) IsError() bool{
    r := int(C.xmmsv_is_error(x.data))
    if r == 1 {
        return true
    }
    return false
}

// XmmsDict
type Dict struct {
	Value
}

func NewDict() *Dict {
	d := new(Dict)
	d.data = C.xmmsv_new_dict()
	return d
}

// XmmsList
type List struct {
    Value
}

func NewList() *List {
    l := new(List)
    l.data = C.xmmsv_new_list()
    return l
}
