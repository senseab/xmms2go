package xmms2go

import "testing"
import "os"

func TestA(t *testing.T) {
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
