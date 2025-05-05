package util

import (
	"bytes"
	"testing"
)

func TestTemplate(t *testing.T) {
	page := &BasePage{
		HeadTitle: "head",
		MainBody:  "Main",
	}

	buf := new(bytes.Buffer)
	WritePageTemplate(buf, page)

	s := buf.String()
	t.Logf("result %s", s)
}
