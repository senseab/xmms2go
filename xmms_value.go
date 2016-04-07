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
	"reflect"
	"regexp"
	"unsafe"
)

// Value type, INT32 is same to INT64
const (
	TypeNone = iota
	TypeError
	TypeInt64
	TypeString
	TypeColl
	TypeBin
	TypeList
	TypeDict
	TypeBitBuffer
	TypeFloat
	TypeEnd
)

// In 32bit system.
const TypeInt32 = TypeInt64

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

// Since Go 1.6, some Go pointer based type
// (like map[int]) may cause cgo panic.
// So return ValueNone.
func NewValueFromAny(any interface{}) ValueAny {
	x := new(Value)
	// OK, <nil> may cause panic, so we need to make a none value.
	if any == nil {
		x = NewValueFromNone().ToValue()
	} else {
		_value := reflect.ValueOf(any)
		switch _value.Type().String() {
		case "int", "int8", "int16", "int32":
			x = NewValueFromInt32(int32(_value.Int())).ToValue()
		case "int64":
			x = NewValueFromInt64(_value.Int()).ToValue()
		case "float32":
			x = NewValueFromFloat32(float32(_value.Float())).ToValue()
		case "float64":
			x = NewValueFromFloat64(_value.Float()).ToValue()
		case "string":
			x = NewValueFromString(_value.String()).ToValue()
		case "[]byte", "[]uint8":
			x = NewValueFromBytes(_value.Bytes()).ToValue()
		case "error":
			x = NewValueFromError(any.(error)).ToValue()
		case "byte":
			x = NewValueFromString(_value.String()).ToValue()
		case "uint", "uint8", "uint16", "uint32":
			x = NewValueFromInt32(int32(_value.Uint())).ToValue()
		case "uint64":
			x = NewValueFromInt64(int64(_value.Uint())).ToValue()
		case "complex64", "complex128":
			f := make([]interface{}, 2)
			f[0] = real(_value.Complex())
			f[1] = imag(_value.Complex())
			l := NewList()
			err := l.FromSlice(f)
			if err != nil {
				x = NewValueFromNone().ToValue()
			} else {
				x = l.ToValue()
			}
		case "uintptr":
			if any.(uintptr) == 0 {
				x = NewValueFromNone().ToValue()
			} else {
				x.data = (*C.xmmsv_t)(unsafe.Pointer(any.(uintptr)))
			}
		case "bool":
			if _value.Bool() {
				x = NewValueFromInt32(1).ToValue()
			} else {
				x = NewValueFromInt32(0).ToValue()
			}
		default:
			// Pointer and Function
			_type := _value.Type()
			matched, err := regexp.MatchString("^\\*|^func", _type.String())
			if err != nil {
				x = NewValueFromNone().ToValue()
			}
			if matched {
				x.data = (*C.xmmsv_t)(unsafe.Pointer(&any))
				break
			}

			// Slice
			matched, err = regexp.MatchString("^\\[\\]", _type.String())
			if err != nil {
				x = NewValueFromNone().ToValue()
				break
			}
			if matched {
				var s []interface{}
				for i := 0; i < _value.Len(); i++ {
					s = append(s, _value.Index(i).Interface())
				}

				l := NewList()
				err := l.FromSlice(s)
				if err != nil {
					x = NewValueFromNone().ToValue()
					break
				}
				x = l.ToValue()
				break
			}

			// Map[string]
			matched, err = regexp.MatchString("^map\\[string\\]", _type.String())
			if err != nil {
				x = NewValueFromNone().ToValue()
				break
			}
			if matched {
				m := make(map[string]interface{})
				for _, k := range _value.MapKeys() {
					m[k.String()] = _value.MapIndex(k).Interface()
				}

				d := NewDict()
				err := d.FromMap(m)
				if err != nil {
					x = NewValueFromNone().ToValue()
					break
				}
				x = d.ToValue()
				break
			}

			x = NewValueFromNone().ToValue()
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
		return nil, fmt.Errorf("Not an error type")
	}
	r := int(C.xmmsv_get_error(x.data, &cErrInfo))
	if r == 0 {
		return nil, fmt.Errorf("Parse type error failed")
	}
	return fmt.Errorf(C.GoString(cErrInfo)), nil
}

func (x *Value) GetInt32() (int32, error) {
	var i C.int32_t
	r := int(C.xmmsv_get_int32(x.data, &i))
	if r == 0 {
		return 0, fmt.Errorf("Parse type int32 failed")
	}
	return int32(i), nil
}

func (x *Value) GetInt64() (int64, error) {
	var i C.int64_t
	r := int(C.xmmsv_get_int64(x.data, &i))
	if r == 0 {
		return 0, fmt.Errorf("Parse type int64 failed")
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
		return 0, fmt.Errorf("Parse type float failed")
	}
	return f, nil
}

func (x *Value) GetBytes() ([]byte, error) {
	var b *C.uchar
	var l C.uint
	defer C.free(unsafe.Pointer(b))
	r := int(C.xmmsv_get_bin(x.data, &b, &l))
	if r == 0 {
		return nil, fmt.Errorf("Parse type bytes failed")
	}
	return C.GoBytes(unsafe.Pointer(b), C.int(l)), nil

}

func (x *Value) GetString() (string, error) {
	var s *C.char
	defer C.free(unsafe.Pointer(s))

	r := int(C.xmmsv_get_string(x.data, &s))
	if r == 0 {
		return "", fmt.Errorf("Parse type string failed")
	}
	return C.GoString(s), nil
}

func (x *Value) GetList() (List, error) {
	if !x.IsType(TypeList) {
		return nil, fmt.Errorf("Parse type list failed")
	}
	l := new(list)
	l.data = x.export()
	var L List = l
	return L, nil
}

func (x *Value) GetDict() (Dict, error) {
	if !x.IsType(TypeDict) {
		return nil, fmt.Errorf("Parse type dict failed")
	}
	d := new(dict)
	d.data = x.export()
	var D Dict = d
	return D, nil
}

func (x *Value) GetCollection() (Collection, error) {
	if !x.IsType(TypeColl) {
		return nil, fmt.Errorf("Parse type collection failed")
	}
	c := new(collection)
	c.data = x.export()
	var _C Collection = c
	return _C, nil
}

func (x *Value) GetBitBuffer() (BitBuffer, error) {
	if !x.IsType(TypeBitBuffer) {
		return nil, fmt.Errorf("Parse type bitbuffer failed")
	}
	b := new(bitBuffer)
	b.data = x.export()
	var _B BitBuffer = b
	return _B, nil
}

func (x *Value) GetAny() (interface{}, error) {
	switch x.GetType() {
	case TypeInt64:
		return x.GetInt64()
	case TypeFloat:
		return x.GetFloat64()
	case TypeString:
		return x.GetString()
	case TypeError:
		return x.GetError()
	case TypeBin:
		return x.GetBytes()
	case TypeList:
		l, err := x.GetList()
		if err != nil {
			return nil, err
		}
		return l.ToSlice()
	case TypeDict:
		d, err := x.GetDict()
		if err != nil {
			return nil, err
		}
		return d.ToMap()
	case TypeColl:
		return x.GetCollection()
	case TypeNone:
		return nil, nil
	case TypeBitBuffer:
		return x.GetBitBuffer()
	default:
		return x.getAny()
	}
	return nil, fmt.Errorf("What?")
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
