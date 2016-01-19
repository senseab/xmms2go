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

func (d *dict) Get(key string) (*Value, error) {
	val := new(Value)
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	r := C.xmmsv_dict_get(d.data, cKey, &(val.data))
	if int(r) == 0 {
		return nil, fmt.Errorf("Get content failed")
	}
	return val, nil
}

func (d *dict) Set(key string, val *Value) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	r := C.xmmsv_dict_set(d.data, cKey, val.export())
	if int(r) == 0 {
		return fmt.Errorf("Set content failed")
	}
	return nil
}

func (d *dict) Remove(key string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	r := C.xmmsv_dict_remove(d.data, cKey)
	if int(r) == 0 {
		return fmt.Errorf("Remove content key='%s' failed", key)
	}
	return nil
}

func (d *dict) Clear() error {
	r := C.xmmsv_dict_clear(d.data)
	if int(r) == 0 {
		return fmt.Errorf("Clear content failed")
	}
	return nil
}

func (d *dict) GetSize() int {
	r := C.xmmsv_dict_get_size(d.data)
	return int(r)
}

func (d *dict) HasKey(key string) bool {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	r := C.xmmsv_dict_has_key(d.data, cKey)
	if int(r) == 0 {
		return false
	}
	return true
}

func (d *dict) GetString(key string) (string, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var retStr *C.char
	defer C.free(unsafe.Pointer(retStr))

	r := C.xmmsv_dict_entry_get_string(d.data, cKey, &retStr)
	if int(r) == 0 {
		return "", fmt.Errorf("Get string content key='%s' failed", key)
	}
	return C.GoString(retStr), nil
}

func (d *dict) GetInt32(key string) (int32, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	var v C.int32_t

	r := C.xmmsv_dict_entry_get_int32(d.data, cKey, &v)
	if int(r) == 0 {
		return 0, fmt.Errorf("Get int32 content key='%s' failed", key)
	}
	return int32(v), nil
}

func (d *dict) GetInt64(key string) (int64, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	var v C.int64_t

	r := C.xmmsv_dict_entry_get_int64(d.data, cKey, &v)
	if int(r) == 0 {
		return 0, fmt.Errorf("Get int64 content key='%s' failed", key)
	}
	return int64(v), nil
}

func (d *dict) getFloat(key string) (C.float, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	var f C.float
	r := C.xmmsv_dict_entry_get_float(d.data, cKey, &f)
	if int(r) == 0 {
		return 0, fmt.Errorf("Get float content key='%s' failed", key)
	}
	return f, nil
}

func (d *dict) GetFloat32(key string) (float32, error) {
	f, err := d.getFloat(key)
	return float32(f), err
}

func (d *dict) GetFloat64(key string) (float64, error) {
	f, err := d.getFloat(key)
	return float64(f), err
}

func (d *dict) SetString(key string, val string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(val)
	defer C.free(unsafe.Pointer(cKey))

	r := C.xmmsv_dict_set_string(d.data, cKey, cValue)
	if int(r) == 0 {
		return fmt.Errorf("Set content key='%s', val='%s' failed", key, val)
	}
	return nil
}

func (d *dict) setInt(key string, val C.int64_t) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	r := C.xmmsv_dict_set_int(d.data, cKey, val)
	if int(r) == 0 {
		return fmt.Errorf("Set content key='%s', val=%d failed", key, int64(val))
	}
	return nil
}

func (d *dict) SetInt64(key string, val int64) error {
	return d.setInt(key, (C.int64_t)(val))
}

func (d *dict) SetInt32(key string, val int32) error {
	return d.setInt(key, (C.int64_t)(val))
}

func (d *dict) setFloat(key string, val C.float) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	r := C.xmmsv_dict_set_float(d.data, cKey, val)
	if int(r) == 0 {
		return fmt.Errorf("Set content key='%s', val=%f failed.", key, float32(val))
	}
	return nil
}

func (d *dict) SetFloat64(key string, val float64) error {
	return d.setFloat(key, (C.float)(val))
}

func (d *dict) SetFloat32(key string, val float32) error {
	return d.setFloat(key, (C.float)(val))
}

func (d *dict) GetType(key string) int {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	r := C.xmmsv_dict_entry_get_type(d.data, cKey)
	return int(r)
}

func (d *dict) FromMap(val map[string]interface{}) error {
	for k, v := range val {
		val := NewValueFromAny(v)
		r := d.Set(k, val)
		if int(r) == 0 {
			return fmt.Errorf("Convert from map[string]interface{} failed, key='%s' value=%v", k, v)
		}
	}
	return nil
}

func (d *dict) ToMap() (map[string]interface{}, error) {
	r := make(map[string]interface{})
	di, err := NewDictIter(d)
	if err != nil {
		return nil, fmt.Errorf("Failed to create dict iter")
	}

	di.First()

	for i := 0; i < d.GetSize(); i++ {
		k, v, err := di.Pair()
		if err != nil {
			return nil, fmt.Errorf("Convert to map[string]interface{} failed")
		}

		val, err := v.GetAny()
		if err != nil {
			return nil, fmt.Errorf("Convert to map[string]interface{} failed, key='%s'", k)
		}
		r[k] = val
	}
	return r, nil
}

type Dict interface {
	ValueNone
	Get() (Value, error)
	GetSize() int
}

type DictIter struct {
	data *C.xmms_dict_iter_t
}

func NewDictIter(val Dict) (*DictIter, error) {
	iter := new(DictIter)
	r := C.xmmsv_get_dict_iter(val.export(), &(iter.data))
	if int(r) == 0 {
		return nil, fmt.Errorf("Get dict iter failed")
	}
	return iter, nil
}

func (i *DictIter) Destroy() {
	C.xmmsv_dict_iter_explicit_destroy(i.data)
}

func (i *DictIter) Pair() (string, *Value, error) {
	var key *C.char
	defer C.free(unsafe.Pointer(key))
	val := new(Value)
	r := C.xmmsv_dict_iter(i.data, &key, &(val.data))
	if int(r) == 0 {
		return "", nil, fmt.Errorf("Pair content failed")
	}
	return C.GoString(key), val, nil
}

func (i *DictIter) Valid() bool {
	r := C.xmmsv_dict_iter_valid(i.data)
	if int(r) == 0 {
		return false
	}
	return true
}

func (i *DictIter) First() {
	C.xmmsv_dict_iter_first(i.data)
}

func (i *DictIter) Next() {
	C.xmmsv_dict_iter_next(i.data)
}

func (i *DictIter) Find(key string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	r := C.xmmsv_dict_iter_find(i.data, cKey)
	if int(r) == 0 {
		return fmt.Errorf("Find key='%s' failed")
	}
	return nil
}

func (i *DictIter) Set(val *Value) error {
	r := C.xmmsv_dict_iter_find(i.data, val.export())
	if int(r) == 0 {
		return fmt.Errorf("Set content failed")
	}
	return nil
}

func (i *DictIter) Remove() error {
	r := C.xmmsv_dict_iter_remove(i.data)
	if int(r) == 0 {
		return fmt.Errorf("Remove content failed")
	}
	return nil
}

func (i *DictIter) PairString() (string, string, error) {
	var key, value *C.char
	defer C.free(unsafe.Pointer(key))
	defer C.free(unsafe.Pointer(value))

	r := C.xmmsv_dict_iter_pair_string(i.data, &key, &value)
	if int(r) == 0 {
		return "", "", fmt.Errorf("Pair string content failed")
	}
	return C.GoString(key), C.GoString(value), nil
}

func (i *DictIter) PairInt32() (string, int32, error) {
	var key *C.char
	defer C.free(unsafe.Pointer(key))

	var val C.int32_t

	r := C.xmmsv_dict_iter_pair_int32(i.data, &key, &val)
	if int(r) == 0 {
		return "", 0, fmt.Errorf("Pair int32 content failed")
	}
	return C.GoString(key), int32(val), nil
}

func (i *DictIter) PairInt64() (string, int64, error) {
	var key *C.char
	defer C.free(unsafe.Pointer(key))

	var val C.int64_t

	r := C.xmmsv_dict_iter_pair_int32(i.data, &key, &val)
	if int(r) == 0 {
		return "", 0, fmt.Errorf("Pair int64 content failed")
	}
	return C.GoString(key), int64(val), nil
}

func (i *DictIter) pairFloat() (string, C.float, error) {
	var key *C.char
	defer C.free(unsafe.Pointer(key))

	var val C.float

	r := C.xmmsv_dict_iter_pair_float(i.data, &key, &val)
	if int(r) == 0 {
		return "", 0, fmt.Errorf("Pair float content failed")
	}
	return C.GoString(key), val, nil
}

func (i *DictIter) GetFloat32(string, float32, error) {
	k, v, err := i.pairFloat()
	return k, float32(v), err
}

func (i *DictIter) GetFloat64(string, float64, error) {
	k, v, err := i.pairFloat()
	return k, float64(v), err
}

func (i *DictIter) SetString(s string) error {
	cS := C.CString(s)
	defer C.free(unsafe.Pointer(cS))

	r := C.xmmsv_dict_iter_set_string(i.data, cS)
	if int(r) == 0 {
		return fmt.Errorf("Set string '%s' failed", s)
	}
	return nil
}

func (i *DictIter) setInt(val C.int64_t) error {
	r := C.xmmsv_dict_iter_set_int(i.data, val)
	if int(r) == 0 {
		return fmt.Errorf("Set int %d failed", int64(val))
	}
	return nil
}

func (i *DictIter) SetInt32(val int32) error {
	return i.setInt((C.int64_t)(val))
}

func (i *DictIter) SetInt64(val int64) error {
	return i.setInt((C.int64_t)(val))
}

func (i *DictIter) setFloat(val C.float) error {
	r := C.xmmsv_dict_iter_set_float(i.data, val)
	if int(r) == 0 {
		return fmt.Errorf("Set float %f failed", float32(val))
	}
	return nil
}

func (i *DictIter) SetFloat32(val float32) error {
	return i.setFloat((C.float)(val))
}

func (i *DictIter) SetFloat64(val float64) error {
	return i.setFloat((C.float)(val))
}
