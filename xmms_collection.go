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

type collection struct {
	Value
}

func NewCollection(_type int) Collection {
	c := new(collection)
	c.data = C.xmmsv_new_coll((C.xmmsv_coll_type_t)(_type))
	var C Collection = c
	return C
}

func (c *collection) SetIDList(ids []int) {
	var cIds [](C.int)
	for _, v := range ids {
		cIds = append(cIds, (C.int)(v))
	}
	C.xmmsv_coll_set_idlist(c.data, (*C.int)(unsafe.Pointer(&cIds[0])))
}

func (c *collection) AddOperand(val *Value) {
	C.xmmsv_coll_add_operand(c.data, val.export())
}

func (c *collection) RemoveOperand(val *Value) {
	C.xmmsv_coll_remove_operand(c.data, val.export())
}

func (c *collection) GetOperands() *Value {
	r := C.xmmsv_coll_operands_get(c.data)
	v := new(Value)
	v.data = r
	return v
}

func (c *collection) SetOperands(val *Value) {
	C.xmmsv_coll_operands_set(c.data, val.export())
}

func (c *collection) IDListAppend(id int) error {
	r := C.xmmsv_coll_idlist_append(c.data, (C.int64_t)(id))
	if int(r) == 0 {
		return fmt.Errorf("Append ID list failed: id=%d", id)
	}
	return nil
}

func (c *collection) IDListInsert(index int, id int) error {
	r := C.xmmsv_coll_idlist_insert(c.data, (C.int)(index), (C.int64_t)(id))
	if int(r) == 0 {
		return fmt.Errorf("Insert ID list failed: index=%d, id=%d", index, id)
	}
	return nil
}

func (c *collection) IDListMove(_old int, _new int) error {
	r := C.xmmsv_coll_idlist_move(c.data, (C.int)(_old), (C.int)(_new))
	if int(r) == 0 {
		return fmt.Errorf("Move %d to %d failed", _old, _new)
	}
	return nil
}

func (c *collection) IDListRemove(index int) error {
	r := C.xmmsv_coll_idlist_remove(c.data, (C.int)(index))
	if int(r) == 0 {
		return fmt.Errorf("Remove ID list failed: index=%d", index)
	}
	return nil
}

func (c *collection) IDListClear() error {
	r := C.xmmsv_coll_idlist_clear(c.data)
	if int(r) == 0 {
		return fmt.Errorf("Clear ID list failed")
	}
	return nil
}

func (c *collection) IDListGetIndexInt32(index int) (int32, error) {
	var v C.int32_t
	r := C.xmmsv_coll_idlist_get_index_int32(c.data, (C.int)(index), &v)
	if int(r) == 0 {
		return 0, fmt.Errorf("Get ID list index value for int32 failed: index=%d", index)
	}
	return int32(v), nil
}

func (c *collection) IDListGetIndexInt64(index int) (int64, error) {
	var v C.int64_t
	r := C.xmmsv_coll_idlist_get_index_int64(c.data, (C.int)(index), &v)
	if int(r) == 0 {
		return 0, fmt.Errorf("Get ID list index value for int64 failed: index=%d", index)
	}
	return int64(v), nil
}

func (c *collection) idListSetIndex(index int, val C.int64_t) error {
	r := C.xmmsv_coll_idlist_set_index(c.data, (C.int)(index), val)
	if int(r) == 0 {
		return fmt.Errorf("Set ID list index value failed: index=%d, value=%d", index, int64(val))
	}
	return nil
}

func (c *collection) IDListSetIndexInt32(index int, val int32) error {
	return c.idListSetIndex(index, (C.int64_t)(val))
}

func (c *collection) IDListSetIndexInt64(index int, val int64) error {
	return c.idListSetIndex(index, (C.int64_t)(val))
}

func (c *collection) IDListGetSize() int {
	return int(C.xmmsv_coll_idlist_get_size(c.data))
}

func (c *collection) IsType(_type int) bool {
	r := C.xmmsv_coll_is_type(c.data, (C.xmmsv_coll_type_t)(_type))
	if int(r) == 0 {
		return false
	}
	return true
}

func (c *collection) GetType() int {
	r := C.xmmsv_coll_get_type(c.data)
	return int(r)
}

func (c *collection) IDListGet() *Value {
	v := new(Value)
	v.data = C.xmmsv_coll_idlist_get(c.data)
	return v
}

func (c *collection) IDListSet(idlist *Value) {
	C.xmmsv_coll_idlist_set(c.data, idlist.export())
}

func (c *collection) AttributeSetString(key string, val string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))
	C.xmmsv_coll_attribute_set_string(c.data, cKey, cVal)
}

func (c *collection) attributeSetInt(key string, val C.int64_t) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	C.xmmsv_coll_attribute_set_int(c.data, cKey, val)
}

func (c *collection) AttributeSetInt32(key string, val int32) {
	c.attributeSetInt(key, (C.int64_t)(val))
}

func (c *collection) AttributeSetInt64(key string, val int64) {
	c.attributeSetInt(key, (C.int64_t)(val))
}

func (c *collection) AttributeSetValue(key string, val *Value) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	C.xmmsv_coll_attribute_set_value(c.data, cKey, val.export())
}

func (c *collection) AttributeRemove(key string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	r := C.xmmsv_coll_attribute_remove(c.data, cKey)
	if int(r) == 0 {
		return fmt.Errorf("Remove attribute '%s' failed", key)
	}
	return nil
}

func (c *collection) AttributeGetString(key string) (string, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal *C.char
	r := C.xmmsv_coll_attribute_get_string(c.data, cKey, &cVal)
	if int(r) == 0 {
		return "", fmt.Errorf("Get attribute '%s' for string failed", key)
	} else {
		defer C.free(unsafe.Pointer(cVal))
	}
	return C.GoString(cVal), nil
}

func (c *collection) AttributeGetInt32(key string) (int32, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal C.int32_t
	r := C.xmmsv_coll_attribute_get_int32(c.data, cKey, &cVal)
	if int(r) == 0 {
		return 0, fmt.Errorf("Get attribute '%s' for int32 failed", key)
	}
	return int32(cVal), nil
}

func (c *collection) AttributeGetInt64(key string) (int64, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	var cVal C.int64_t
	r := C.xmmsv_coll_attribute_get_int64(c.data, cKey, &cVal)
	if int(r) == 0 {
		return 0, fmt.Errorf("Get attribute '%s' for int64 failed", key)
	}
	return int64(cVal), nil
}

func (c *collection) AttributeGetValue(key string) (*Value, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	v := new(Value)
	r := C.xmmsv_coll_attribute_get_value(c.data, cKey, &(v.data))
	if int(r) == 0 {
		return nil, fmt.Errorf("Get attribute '%s' failed", key)
	}
	return v, nil
}

func (c *collection) AttributesGet() *Value {
	v := new(Value)
	v.data = C.xmmsv_coll_attributes_get(c.data)
	return v
}

func (c *collection) AttributesSet(val *Value) {
	C.xmmsv_coll_attributes_set(c.data, val.export())
}

func (c *collection) AddOrderOperator(order *Value) Collection {
	v := new(collection)
	v.data = C.xmmsv_coll_add_order_operator(c.data, order.export())
	var V Collection = v
	return V
}

func (c *collection) AddOrderOperators(order *Value) Collection {
	v := new(collection)
	v.data = C.xmmsv_coll_add_order_operators(c.data, order.export())
	var V Collection = v
	return V
}

func (c *collection) AddLimitOperator(start int, length int) Collection {
	v := new(collection)
	v.data = C.xmmsv_coll_add_limit_operator(c.data, (C.int)(start), (C.int)(length))
	var V Collection = v
	return V
}

type Collection interface {
	ValueNone
	SetIDList(ids []int)
	AddOperand(val *Value)
	RemoveOperand(val *Value)
	GetOperands() *Value
	SetOperands(val *Value)
	IDListAppend(id int) error
	IDListInsert(index int, id int) error
	IDListMove(_old int, _new int) error
	IDListRemove(index int) error
	IDListClear() error
	IDListGetIndexInt32(index int) (int32, error)
	IDListGetIndexInt64(index int) (int64, error)
	IDListSetIndexInt32(index int, val int32) error
	IDListSetIndexInt64(index int, val int64) error
	IDListGetSize() int
	IsType(_type int) bool
	GetType() int
	IDListGet() *Value
	IDListSet(idlist *Value)
	AttributeSetString(key string, val string)
	AttributeSetInt32(key string, val int32)
	AttributeSetInt64(key string, val int64)
	AttributeSetValue(key string, val *Value)
	AttributeRemove(key string) error
	AttributeGetString(key string) (string, error)
	AttributeGetInt32(key string) (int32, error)
	AttributeGetInt64(key string) (int64, error)
	AttributeGetValue(key string) (*Value, error)
	AttributesGet() *Value
	AttributesSet(val *Value)
	AddOrderOperator(order *Value) Collection
	AddOrderOperators(order *Value) Collection
	AddLimitOperator(start int, length int) Collection
}
