package myint

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	data := []struct {
		title  string
		value  MyInt
		param  int
		should MyInt
		err    error
	}{
		{title: "A", value: 1, param: 1, should: 2, err: nil},
		{"B", 2, -1, 1, nil},
		{"C", 0, 0, 0, nil},
		{"D", 0, math.MaxInt32, math.MaxInt32, nil},
		{"E", 0, math.MinInt32, math.MinInt32, nil},
		{"F", 1, math.MaxInt32, -1, ErrOutOfRange},
		{"G", math.MinInt32, math.MaxInt32, -1, nil},
	}
	for _, v := range data {
		mi := v.value
		get, err := mi.Add(v.param)
		if get != v.should {
			t.Error("for", v.title, "got", mi, "should got", v.should)
		}
		if v.err != nil {
			if err == nil {
				t.Error("for", v.title, "got nil error should got", v.err.Error())
			} else if v.err.Error() != err.Error() {
				t.Error("for", v.title, "got error", err.Error(), "should got", v.err.Error())
				t.Errorf("value got %v value should %v", get, v.should)
			}
		}
	}
}
