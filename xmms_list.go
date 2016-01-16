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

// XmmsList
type list struct {
	Value
}

func NewList() List {
	l := new(list)
	l.data = C.xmmsv_new_list()
	var L List = l
	return L
}

func (l *list) Get(pos int) (*Value, error) {
	val := new(Value)
	r := C.xmmsv_list_get(l.data, C.int(pos), &(val.data))
	if int(r) == 0 {
		return nil, errors.New("Get content failed")
	}
	return val, nil
}

func (l *list) Set(pos int, val *Value) error {
	r := C.xmmsv_list_set(l.data, C.int(pos), val.export())
	if int(r) == 0 {
		return errors.New("Set content failed")
	}
	return nil
}

func (l *list) Append(val *Value) error {
	r := C.xmmsv_list_append(l.data, val.export())
	if int(r) == 0 {
		return errors.New("Append content failed")
	}
	return nil
}

func (l *list) Insert(pos int, val *Value) error {
	r := C.xmmsv_list_insert(l.data, C.int(pos), val.export())
	if int(r) == 0 {
		return errors.New("Insert content failed")
	}
	return nil
}

func (l *list) Remove(pos int) error {
	r := C.xmmsv_list_remove(l.data, C.int(pos))
	if int(r) == 0 {
		return errors.New("Remove content failed")
	}
	return nil
}

func (l *list) Move(posOld int, posNew int) error {
	r := C.xmmsv_list_move(l.data, C.int(posOld), C.int(posNew))
	if int(r) == 0 {
		return errors.New("Move content failed")
	}
	return nil
}

func (l *list) Clear() error {
	r := C.xmmsv_list_clear(l.data)
	if int(r) == 0 {
		return errors.New("Clear content failed")
	}
	return nil
}

/*
func (l *List) Sort(sortMethod func) error {
}
*/

func (l *list) GetSize() int {
	return int(C.xmmsv_list_get_size(l.data))
}

func (l *list) RestrictType(_type int) error {
	r := C.xmmsv_list_restrict_type(l.data, C.xmmsv_type_t(_type))
	if int(r) == 0 {
		return errors.New("Restrict type failed")
	}
	return nil
}

func (l *list) HasType(_type int) bool {
	r := C.xmmsv_list_has_type(l.data, C.xmmsv_type_t(_type))
	if int(r) != 0 {
		return true
	}
	return false
}

func (l *list) GetType() (int, error) {
	var t C.xmmsv_type_t
	r := C.xmmsv_list_get_type(l.data, &t)
	if int(r) != 0 {
		return -1, errors.New("Get type failed")
	}
	return int(t), nil
}

func (l *list) IndexOf(val *Value) int {
	r := C.xmmsv_list_index_of(l.data, val.export())
	return int(r)
}

func (l *list) GetString(pos int) (string, error) {
	var s *C.char
	defer C.free(unsafe.Pointer(s))
	r := C.xmmsv_list_get_string(l.data, C.int(pos), &s)
	if int(r) == 0 {
		return "", errors.New("Get string failed")
	}
	return C.GoString(s), nil
}

func (l *list) GetInt32(pos int) (int32, error) {
	var i C.int32_t
	r := C.xmmsv_list_get_int32(l.data, C.int(pos), &i)
	if int(r) == 0 {
		return -1, errors.New("Get int32 failed")
	}
	return int32(i), nil
}

func (l *list) GetInt64(pos int) (int64, error) {
	var i C.int64_t
	r := C.xmmsv_list_get_int64(l.data, C.int(pos), &i)
	if int(r) == 0 {
		return -1, errors.New("Get int64 failed")
	}
	return int64(i), nil
}

func (l *list) getFloat(pos int) (C.float, error) {
	var f C.float
	r := C.xmmsv_list_get_float(l.data, C.int(pos), &f)
	if int(r) == 0 {
		return -1, errors.New("Get float failed")
	}
	return f, nil
}

func (l *list) GetFloat32(pos int) (float32, error) {
	f, err := l.getFloat(pos)
	return float32(f), err
}

func (l *list) GetFloat64(pos int) (float64, error) {
	f, err := l.getFloat(pos)
	return float64(f), err
}

func (l *list) SetString(pos int, val string) error {
	s := C.CString(val)
	defer C.free(unsafe.Pointer(s))
	r := C.xmmsv_list_set_string(l.data, C.int(pos), s)
	if int(r) == 0 {
		return errors.New("Set string failed")
	}
	return nil
}

func (l *list) setInt(pos int, val C.int64_t) error {
	r := C.xmmsv_list_set_int(l.data, C.int(pos), val)
	if int(r) == 0 {
		return errors.New("Set int failed")
	}
	return nil
}

func (l *list) SetInt32(pos int, val int32) error {
	i := C.int64_t(val)
	return l.setInt(pos, i)
}

func (l *list) SetInt64(pos int, val int64) error {
	i := C.int64_t(val)
	return l.setInt(pos, i)
}

func (l *list) setFloat(pos int, val C.float) error {
	r := C.xmmsv_list_set_float(l.data, C.int(pos), val)
	if int(r) == 0 {
		return errors.New("Set float failed")
	}
	return nil
}

func (l *list) SetFloat32(pos int, val float32) error {
	f := C.float(val)
	return l.setFloat(pos, f)
}

func (l *list) SetFloat64(pos int, val float64) error {
	f := C.float(val)
	return l.setFloat(pos, f)
}

func (l *list) InsertString(pos int, val string) error {
	s := C.CString(val)
	defer C.free(unsafe.Pointer(s))
	r := C.xmmsv_list_insert_string(l.data, C.int(pos), s)
	if int(r) == 0 {
		return errors.New("Insert string failed")
	}
	return nil
}

func (l *list) insertInt(pos int, val C.int64_t) error {
	r := C.xmmsv_list_insert_int(l.data, C.int(pos), val)
	if int(r) == 0 {
		return errors.New("Insert int failed")
	}
	return nil
}

func (l *list) InsertInt32(pos int, val int32) error {
	i := C.int64_t(val)
	return l.insertInt(pos, i)
}

func (l *list) InsertInt64(pos int, val int64) error {
	i := C.int64_t(val)
	return l.insertInt(pos, i)
}

func (l *list) insertFloat(pos int, val C.float) error {
	r := C.xmmsv_list_insert_float(l.data, C.int(pos), val)
	if int(r) == 0 {
		return errors.New("Insert float failed")
	}
	return nil
}

func (l *list) InsertFloat32(pos int, val float32) error {
	f := C.float(val)
	return l.insertFloat(pos, f)
}

func (l *list) InsertFloat64(pos int, val float64) error {
	f := C.float(val)
	return l.insertFloat(pos, f)
}

func (l *list) AppendString(val string) error {
	s := C.CString(val)
	defer C.free(unsafe.Pointer(s))
	r := C.xmmsv_list_append_string(l.data, s)
	if int(r) == 0 {
		return errors.New("Append string failed")
	}
	return nil
}

func (l *list) appendInt(val C.int64_t) error {
	r := C.xmmsv_list_append_int(l.data, val)
	if int(r) == 0 {
		return errors.New("Append int failed")
	}
	return nil
}

func (l *list) AppendInt32(val int32) error {
	i := C.int64_t(val)
	return l.appendInt(i)
}

func (l *list) AppendInt64(val int64) error {
	i := C.int64_t(val)
	return l.appendInt(i)
}

func (l *list) appendFloat(val C.float) error {
	r := C.xmmsv_list_append_float(l.data, val)
	if int(r) == 0 {
		return errors.New("Append float failed")
	}
	return nil
}

func (l *list) AppendFloat32(val float32) error {
	f := C.float(val)
	return l.appendFloat(f)
}

func (l *list) AppendFloat64(val float64) error {
	f := C.float(val)
	return l.appendFloat(f)
}

func (l *list) Flatten(dep int) List {
	v := new(list)
	v.data = C.xmmsv_list_flatten(l.data, C.int(dep))
	var V List = v
	return V
}

func (l *list) FromSlice(s []interface{}) error {
	for _, v := range s {
		val := NewValueFromAny(v)
		err := l.Append(val.ToValue())
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *list) ToSlice() ([]interface{}, error) {
	v := make([]interface{}, 0)

	i, err := NewListIter(l)
	defer i.Destroy()
	if err != nil {
		return nil, err
	}

	for a := 0; a < l.GetSize(); a++ {
		val, err := i.Entry()
		if err != nil {
			return nil, err
		}
		intf, err := val.GetAny()
		if err != nil {
			return nil, err
		}

		v = append(v, intf)
		i.Next()
	}

	return v, nil
}

type List interface {
	ValueNone
	Append(val *Value) error
	AppendFloat32(val float32) error
	AppendFloat64(val float64) error
	AppendInt32(val int32) error
	AppendInt64(val int64) error
	AppendString(val string) error
	Clear() error
	Flatten(dep int) List
	FromSlice(s []interface{}) error
	Get(pos int) (*Value, error)
	GetFloat32(pos int) (float32, error)
	GetFloat64(pos int) (float64, error)
	GetInt32(pos int) (int32, error)
	GetInt64(pos int) (int64, error)
	GetSize() int
	GetString(pos int) (string, error)
	GetType() (int, error)
	HasType(_type int) bool
	IndexOf(val *Value) int
	Insert(pos int, val *Value) error
	InsertFloat32(pos int, val float32) error
	InsertFloat64(pos int, val float64) error
	InsertInt32(pos int, val int32) error
	InsertInt64(pos int, val int64) error
	InsertString(pos int, val string) error
	Move(posOld int, posNew int) error
	Remove(pos int) error
	RestrictType(_type int) error
	Set(pos int, val *Value) error
	SetFloat32(pos int, val float32) error
	SetFloat64(pos int, val float64) error
	SetInt32(pos int, val int32) error
	SetInt64(pos int, val int64) error
	SetString(pos int, val string) error
	ToSlice() ([]interface{}, error)
}

// ListIter is cursor to List
type ListIter struct {
	data *C.xmmsv_list_iter_t
}

// Get a new list iter
func NewListIter(val List) (*ListIter, error) {
	l := new(ListIter)
	r := C.xmmsv_get_list_iter(val.export(), &l.data)
	if int(r) == 0 {
		return nil, errors.New("Get list iter failed")
	}
	return l, nil
}

func (l *ListIter) Destroy() {
	C.xmmsv_list_iter_explicit_destroy(l.data)
}

func (l *ListIter) Entry() (*Value, error) {
	e := new(Value)
	r := C.xmmsv_list_iter_entry(l.data, &e.data)
	if int(r) == 0 {
		return nil, errors.New("Get entry failed")
	}
	return e, nil
}

func (l *ListIter) Valid() bool {
	r := C.xmmsv_list_iter_valid(l.data)
	if int(r) == 0 {
		return false
	}
	return true
}

// Point to the first element
func (l *ListIter) First() {
	C.xmmsv_list_iter_first(l.data)
}

// Point to the last element
func (l *ListIter) Last() {
	C.xmmsv_list_iter_last(l.data)
}

// Point to the next element
func (l *ListIter) Next() {
	C.xmmsv_list_iter_next(l.data)
}

// Point to the previous element
func (l *ListIter) Prev() {
	C.xmmsv_list_iter_prev(l.data)
}

// Goto positon.
func (l *ListIter) Seek(pos int) error {
	r := C.xmmsv_list_iter_seek(l.data, C.int(pos))
	if int(r) == 0 {
		return errors.New("Seek failed")
	}
	return nil
}

// Tell the position of the ListIter
func (l *ListIter) Tell() int {
	r := C.xmmsv_list_iter_tell(l.data)
	return int(r)
}

func (l *ListIter) GetParent() *Value {
	v := new(Value)
	v.data = C.xmmsv_list_iter_get_parent(l.data)
	return v
}

func (l *ListIter) Set(val *Value) error {
	r := C.xmmsv_list_iter_set(l.data, val.export())
	if int(r) == 0 {
		return errors.New("Set value failed")
	}
	return nil
}

func (l *ListIter) Insert(val *Value) error {
	r := C.xmmsv_list_iter_insert(l.data, val.export())
	if int(r) == 0 {
		return errors.New("Set value failed")
	}
	return nil
}

func (l *ListIter) Remove() error {
	r := C.xmmsv_list_iter_remove(l.data)
	if int(r) == 0 {
		return errors.New("Remove value failed")
	}
	return nil
}

func (l *ListIter) EntryString() (string, error) {
	var s *C.char
	defer C.free(unsafe.Pointer(s))
	r := C.xmmsv_list_iter_entry_string(l.data, &s)
	if int(r) == 0 {
		return "", errors.New("Get string failed")
	}
	return C.GoString(s), nil
}

func (l *ListIter) EntryInt32() (int32, error) {
	var i C.int32_t
	r := C.xmmsv_list_iter_entry_int32(l.data, &i)
	if int(r) == 0 {
		return 0, errors.New("Get int32 failed")
	}
	return int32(i), nil
}

func (l *ListIter) EntryInt64() (int64, error) {
	var i C.int64_t
	r := C.xmmsv_list_iter_entry_int64(l.data, &i)
	if int(r) == 0 {
		return 0, errors.New("Get int64 failed")
	}
	return int64(i), nil
}

func (l *ListIter) entryFloat() (C.float, error) {
	var f C.float
	r := C.xmmsv_list_iter_entry_float(l.data, &f)
	if int(r) == 0 {
		return 0, errors.New("Get float failed")
	}
	return f, nil
}

func (l *ListIter) EntryFloat32() (float32, error) {
	f, err := l.entryFloat()
	return float32(f), err
}

func (l *ListIter) EntryFloat64() (float64, error) {
	f, err := l.entryFloat()
	return float64(f), err
}

func (l *ListIter) InsertString(s string) error {
	cS := C.CString(s)
	defer C.free(unsafe.Pointer(cS))
	r := C.xmmsv_list_iter_insert_string(l.data, cS)
	if int(r) == 0 {
		return errors.New("Insert String failed")
	}
	return nil
}

func (l *ListIter) insertInt(i C.int64_t) error {
	r := C.xmmsv_list_iter_insert_int(l.data, i)
	if int(r) == 0 {
		return errors.New("Insert int failed")
	}
	return nil
}

func (l *ListIter) InsertInt32(i int32) error {
	return l.insertInt(C.int64_t(i))
}

func (l *ListIter) InsertInt64(i int64) error {
	return l.insertInt(C.int64_t(i))
}

func (l *ListIter) insertFloat(f C.float) error {
	r := C.xmmsv_list_iter_insert_float(l.data, f)
	if int(r) == 0 {
		return errors.New("Insert float failed")
	}
	return nil
}

func (l *ListIter) InsertFloat32(f float32) error {
	return l.insertFloat(C.float(f))
}

func (l *ListIter) InsertFloat64(f float64) error {
	return l.insertFloat(C.float(f))
}
