package multiwriter

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestMultiWriter_Write(t *testing.T) {
	mw := MultiWriter{
		IO1: nil,
		IO2: nil,
	}

	t.Run("Write on both writers", func(t *testing.T) {
		b1 := bytes.NewBufferString("")
		b2 := bytes.NewBufferString("")
		mw.IO1 = b1
		mw.IO2 = b2

		n, err := mw.Write([]byte("test-write"))
		if err != nil {
			t.Fatal(err)
		}

		if n != 10 {
			t.Errorf("write wrong number of bytes, %d", n)
		}

		data, _ := ioutil.ReadAll(b1)
		if string(data) != "test-write" {
			t.Errorf("write wrong data, %s", string(data))
		}

		data, _ = ioutil.ReadAll(b2)
		if string(data) != "test-write" {
			t.Errorf("write wrong data, %s", string(data))
		}
	})
}
