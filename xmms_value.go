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

// Value type, INT32 is same to INT64
const (
	XMMSV_TYPE_NONE = iota
	XMMSV_TYPE_ERROR
	XMMSV_TYPE_INT64
	XMMSV_TYPE_STRING
	XMMSV_TYPE_COLL
	XMMSV_TYPE_BIN
	XMMSV_TYPE_LIST
	XMMSV_TYPE_DICT
	XMMSV_TYPE_BITBUFFER
	XMMSV_TYPE_FLOAT
	XMMSV_TYPE_END
)

const XMMSV_TYPE_INT32 = XMMSV_TYPE_INT64

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

func NewValueFromAny(any interface{}) ValueAny {
	x := new(Value)
	switch any.(type) {
	case int:
		x = NewValueFromInt32(any.(int32)).ToValue()
	case int32:
		x = NewValueFromInt32(any.(int32)).ToValue()
	case int64:
		x = NewValueFromInt64(any.(int64)).ToValue()
	case float32:
		x = NewValueFromFloat32(any.(float32)).ToValue()
	case float64:
		x = NewValueFromFloat64(any.(float64)).ToValue()
	case string:
		x = NewValueFromString(any.(string)).ToValue()
	case []byte:
		x = NewValueFromBytes(any.([]byte)).ToValue()
	case error:
		x = NewValueFromError(any.(error)).ToValue()
	case bool:
		if any.(bool) {
			x = NewValueFromInt32(1).ToValue()
		} else {
			x = NewValueFromInt32(0).ToValue()
		}
	default: // Pointer?
		if any == nil {
			x = NewValueFromNone().ToValue()
		} else {
			x.data = (*C.xmmsv_t)(unsafe.Pointer(&any))
		}
	}

	var va ValueAny = x
	return va
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

func NewValueFromFloat64(f float64) ValueFloat64 {
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

	x.data = C.xmmsv_new_bin(d, C.uint(length))
	var vb ValueBytes = x
	return vb
}

func NewValueFromCopyValue(v *Value) *Value {
	x := new(Value)
	x.data = C.xmmsv_copy(v.export())
	return x
}

func NewValueFromRef(v *Value) *Value {
	x := new(Value)
	x.data = C.xmmsv_ref(v.export())
	return x
}

// Okay, any one should not use native C types.
func (x *Value) export() *C.xmmsv_t {
	return x.data
}

func (x *Value) Unref() {
	C.xmmsv_unref(x.data)
}

func (x *Value) GetType() int {
	return int(C.xmmsv_get_type(x.data))
}

func (x *Value) IsType(t int) bool {
	if int(C.xmmsv_is_type(x.data, C.xmmsv_type_t(t))) == 1 {
		return true
	}
	return false
}

func (x *Value) getAny() (interface{}, error) {
	v := new(interface{})
	v = (*interface{})(unsafe.Pointer(x.data))
	return *v, nil
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

func (x *Value) GetString() (string, error) {
	var s *C.char
	defer C.free(unsafe.Pointer(s))

	r := int(C.xmmsv_get_string(x.data, &s))
	if r == 0 {
		return "", errors.New("Parse type string failed")
	}
	return C.GoString(s), nil
}

func (x *Value) GetList() (List, error) {
	if !x.IsType(XMMSV_TYPE_LIST) {
		return nil, errors.New("Parse type list failed")
	}
	l := new(list)
	l.data = x.export()
	var L List = l
	return L, nil
}

// Dummy
func (x *Value) GetDict() (Dict, error) {
	return nil, nil
}

// Dummy
func (x *Value) GetCollection() (Collection, error) {
	return nil, nil
}

// Dummy
func (x *Value) GetBitBuffer() (BitBuffer, error) {
	return nil, nil
}

func (x *Value) GetAny() (interface{}, error) {
	switch x.GetType() {
	case XMMSV_TYPE_INT64:
		return x.GetInt64()
	case XMMSV_TYPE_FLOAT:
		return x.GetFloat64()
	case XMMSV_TYPE_STRING:
		return x.GetString()
	case XMMSV_TYPE_ERROR:
		return x.GetError()
	case XMMSV_TYPE_BIN:
		return x.GetBytes()
	case XMMSV_TYPE_LIST:
		return x.GetList()
	case XMMSV_TYPE_DICT:
		return x.GetDict()
	case XMMSV_TYPE_COLL:
		return x.GetCollection()
	case XMMSV_TYPE_NONE:
		return nil, nil
	case XMMSV_TYPE_BITBUFFER:
		return x.GetBitBuffer()
	default:
		return x.getAny()
	}
	return nil, errors.New("What?")
}

// Okay, we need to implement the collection type.

func (x *Value) IsError() bool {
	r := int(C.xmmsv_is_error(x.data))
	if r == 1 {
		return true
	}
	return false
}

func (x *Value) ToValue() *Value {
	v := new(Value)
	v.data = C.xmmsv_ref(x.export())
	return v
}

type ValueNone interface {
	export() *C.xmmsv_t
	Unref()
	ToValue() *Value
}

type ValueAny interface {
	ValueNone
	GetAny() (interface{}, error)
}

type ValueError interface {
	ValueNone
	IsError() bool
	GetError() (error, error)
}

type ValueInt64 interface {
	ValueNone
	GetInt64() (int64, error)
}

type ValueInt32 interface {
	ValueNone
	GetInt32() (int32, error)
}

type ValueFloat64 interface {
	ValueNone
	GetFloat64() (float64, error)
}

type ValueFloat32 interface {
	ValueNone
	GetFloat32() (float32, error)
}

type ValueString interface {
	ValueNone
	GetString() (string, error)
}

type ValueBytes interface {
	ValueNone
	GetBytes() ([]byte, error)
}
