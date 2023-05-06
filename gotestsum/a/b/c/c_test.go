package c_test

import "testing"

func TestFail(t *testing.T) {
	t.Log("some logs: fail")
	t.Error("test")
}

func TestPass(t *testing.T) {
	t.Log("some logs: pass")
}

func TestSkip(t *testing.T) {
	t.Log("some logs: skip")
	t.Skip("skip")
}

func TestNested(t *testing.T) {
	t.Run("fail", func(t *testing.T) {
		t.Log("some logs: fail")
		t.Error("nested: fail")
	})

	t.Run("pass", func(t *testing.T) {
		t.Log("some logs: pass")
	})

	t.Run("fail", func(t *testing.T) {
		t.Log("some logs: skip")
		t.Skip("skip")
	})
}

func TestParallelNested(t *testing.T) {
	t.Parallel()

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Log("some logs: fail")
		t.Error("nested: fail")
	})

	t.Run("pass", func(t *testing.T) {
		t.Parallel()
		t.Log("some logs: pass")
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Log("some logs: skip")
		t.Skip("skip")
	})
}
