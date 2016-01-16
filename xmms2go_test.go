package xmms2go

import (
	"errors"
	"os"
	"testing"
)

func TestValue(t *testing.T) {
	// Start from None.
	vn := NewValueFromNone()
	defer vn.Unref()

	ve := NewValueFromError(errors.New("ValueError Test"))
	defer ve.Unref()
	t.Log("ve is error:", ve.IsError())
	veo, err := ve.GetError()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test error value:", veo)

	vi64 := NewValueFromInt64(17)
	defer vi64.Unref()
	vi64o, err := vi64.GetInt64()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test int64 value:", vi64o)

	vi32 := NewValueFromInt32(23)
	defer vi32.Unref()
	vi32o, err := vi32.GetInt32()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test int32 value:", vi32o)

	vf64 := NewValueFromFloat64(1.4)
	defer vf64.Unref()
	vf64o, err := vf64.GetFloat64()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test float64 value:", vf64o)

	vf32 := NewValueFromFloat32(1.5)
	defer vf32.Unref()
	vf32o, err := vf32.GetFloat32()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test float32 value:", vf32o)

	vs := NewValueFromString("Test string")
	defer vs.Unref()
	vso, err := vs.GetString()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test string value:", vso)

	vb := NewValueFromBytes([]byte("Test\tTest"))
	defer vb.Unref()
	vbo, err := vb.GetBytes()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test bytes value:", vbo, string(vbo))

	va := NewValueFromAny(func() string { return "Okay" })
	defer va.Unref()
	vao, err := va.GetAny()
	if err != nil {
		t.Error(err)
	}
	// Holy! Go type interface{} is Okay!
	t.Log("Got test anytype value: func() ->", vao.(func() string)())

}

func TestList(t *testing.T) {
	slice1 := make([]interface{}, 0)
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, int64(i))
	}
	t.Log("Slice1=", slice1)

	li64 := NewList()
	defer li64.Unref()
	err := li64.FromSlice(slice1)
	if err != nil {
		t.Error(err)
	}
	t.Log("Size=", li64.GetSize())

	slice2, err := li64.ToSlice()
	if err != nil {
		t.Error(err)
	}
	t.Log("Slice2=", slice2)

	slice3 := make([]interface{}, 0)
	var empty interface{}
	t.Log("empty=", empty)
	for i := 0; i < 10; i++ {
		slice3 = append(slice3, empty)
	}
	t.Log("Slice3=", slice3)
	lempty := NewList()
	defer lempty.Unref()
	err = lempty.FromSlice(slice3)
	if err != nil {
		t.Error(err)
	}
	t.Log("Size=", lempty.GetSize())

	slice4, err := lempty.ToSlice()
	if err != nil {
		t.Error(err)
	}
	t.Log("Slice4=", slice4)
}

func TestClient(t *testing.T) {
	X, err := NewXmms2Client("xmms2go-test")
	defer X.Unref()
	if err != nil {
		t.Error(err)
	}

	err = X.Connect(os.Getenv("XMMS_PATH"))
	if err != nil {
		t.Error(err)
	}

	err = X.Play()
	if err != nil {
		t.Error(err)
	}

	i, err := X.CurrentID()
	if err != nil {
		t.Error(err)
	}
	t.Log("Current Playing ID:", i)
}
