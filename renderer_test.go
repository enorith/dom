package dom_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/enorith/dom"
)

func Test_NodeRender(t *testing.T) {
	n := dom.NewNode([]byte("div"))
	n.SetContent([]byte("test")).SetAttribute([]byte("class"), []byte("content"))

	t.Logf("%s", dom.H("a", map[string]string{
		"src": "http://test.dex",
	}, "123123", n))
}

func Test_parse(t *testing.T) {
	f, e := os.OpenFile("test.html", os.O_RDONLY, 0775)
	if e != nil {
		t.Fatal(e)
	}
	f.Seek(0, 0)
	var nameBuffer []byte
	var stack []byte
	var nameStart bool
	var names [][]byte
	for {
		buffer := make([]byte, 1)
		_, e := f.Read(buffer)
		if e == nil {
			stack = append(stack, buffer...)
			// if bytes.Equal(buffer, []byte{' '}) || bytes.Equal(buffer, []byte{'\n'}) {
			// 	continue
			// }
			if bytes.Equal(buffer, []byte{'<'}) {
				nameStart = true
				continue
			}
			if bytes.Equal(buffer, []byte{'/'}) || bytes.Equal(buffer, []byte{' '}) || bytes.Equal(buffer, []byte{'>'}) {
				nameStart = false
				if len(nameBuffer) > 0 {
					names = append(names, nameBuffer)
					nameBuffer = nil
				}

				continue
			}
			if nameStart {
				nameBuffer = append(nameBuffer, buffer...)
			}

		} else {
			break
		}
	}

	for _, v := range names {
		t.Logf("%s", v)
	}
	t.Logf("%s", stack)
}
