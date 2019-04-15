package models

import "testing"

func Test_ExportType(t *testing.T) {
	t.Run("convert expected word", func(t *testing.T) {
		var e *ExportType
		list := e.GetList()
		for _, v := range list {
			val := v.String()
			if val == UnknownStr {
				t.Fatalf("failed test %#v", v)
			}
		}
	})

	t.Run("convert unexpected word", func(t *testing.T) {
		e, err := ToType("unexpected")
		if err == nil {
			t.Fatalf("failed test")
		}
		if e != UnknownType {
			t.Fatalf("failed test %#v", e)
		}
	})
}
