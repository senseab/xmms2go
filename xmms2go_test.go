package xmms2go

import "testing"
import "os"

func TestA(t *testing.T) {
	X, err := NewXmms2Client("xmms2go-test")
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

	X.Unref()
}
