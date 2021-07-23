package lwatch

import "testing"

func TestNew(t *testing.T) {
	r, err := New("../assets/demo.txt")
	if err != nil {
		t.Error(err)
	}
	defer r.F.Close()
}

func TestGetFileSize(t *testing.T) {
	r, err := New("../assets/demo.txt")
	if err != nil {
		t.Error(err)
	}
	defer r.F.Close()

	_, err = r.GetFileSize()
	if err != nil {
		t.Error(err)
	}
}

func TestGetTail(t *testing.T) {
	r, err := New("../assets/demo.txt")
	if err != nil {
		t.Error(err)
	}

	defer r.F.Close()
	_, err = r.GetTail(0)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAppendedString(t *testing.T) {
	r, err := New("../assets/demo.txt")
	if err != nil {
		t.Error(err)
	}

	_, err = r.GetAppendedString(0, 0)
	if err != nil {
		t.Error(err)
	}
}

func TestNeedleExists(t *testing.T) {
	r, err := New("../assets/demo.txt")
	if err != nil {
		t.Error(err)
	}
	defer r.F.Close()

	ok := r.NeedleExists("123", "abc123")
	if !ok {
		t.Error("needle should exist")
	}
}

func TestNeedleNotExists(t *testing.T) {
	r, err := New("../assets/demo.txt")
	if err != nil {
		t.Error(err)
	}
	defer r.F.Close()

	ok := r.NeedleExists("XYZ", "abc123")
	if ok {
		t.Error("needle should exist")
	}
}
