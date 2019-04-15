package models

import "testing"

func Test_Output(t *testing.T) {
	t.Run("convert expected word", func(t *testing.T) {
		var o *Output
		list := o.GetList()
		for _, v := range list {
			val := v.String()
			if val == UnknownStr {
				t.Fatalf("failed test %#v", v)
			}
		}
	})

	t.Run("convert unexpected word", func(t *testing.T) {
		o, err := ToOutput("unexpected")
		if err == nil {
			t.Fatalf("failed test")
		}
		if o != UnknownOutput {
			t.Fatalf("failed test %#v", o)
		}
	})
}
