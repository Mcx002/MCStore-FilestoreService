package utils

import "testing"

func TestStringToBool(t *testing.T) {
	val := StringToBool("true")
	if val != true {
		t.Errorf("got %v; want true", val)
	}

	val = StringToBool("1")
	if val != true {
		t.Errorf("got %v; want true", val)
	}

	val = StringToBool("false")
	if val != false {
		t.Errorf("got %v; want false", val)
	}

	val = StringToBool("0")
	if val != false {
		t.Errorf("got %v; want false", val)
	}

	val = StringToBool("not bool str")
	if val != false {
		t.Errorf("got %v; want false", val)
	}
}
