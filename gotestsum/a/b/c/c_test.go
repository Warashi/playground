package c_test

import "testing"

func TestFail(t *testing.T) {
	t.Log("some logs: fail")
	t.Error("test")
}
