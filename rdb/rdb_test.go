package rdb

import "testing"

func TestROffset(t *testing.T) {
	r := New()
	defer r.cli.Close()

	_, err := r.ROffset("test")
	if err != nil {
		t.Error(err)
	}
}

func TestWOffSet(t *testing.T) {
	r := New()
	defer r.cli.Close()

	_, err := r.WOffSet("test", 10)
	if err != nil {
		t.Error(err)
	}
}

func TestWLastSent(t *testing.T) {
	r := New()
	defer r.cli.Close()

	_, err := r.WLastSent("demo@gmail.com", 123456)
	if err != nil {
		t.Error(err)
	}
}

func TestRLastSent(t *testing.T) {
	r := New()
	defer r.cli.Close()

	_, err := r.RLastSent("demo@gmail.com")
	if err != nil {
		t.Error(err)
	}
}
