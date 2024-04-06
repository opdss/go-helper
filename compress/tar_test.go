package compress

import (
	"os"
	"testing"
)

func TestPack(t *testing.T) {
	src := "../"
	dest := "./test.tar.gz"
	err := Pack(dest, src)
	if err != nil {
		t.Error(err)
	}
	os.Remove(dest)
}

func TestPackMatch(t *testing.T) {
	src := "../"
	dest := "./test.tar.gz"
	err := PackMatch(dest, src, FileMatch("unit/*"))
	if err != nil {
		t.Error(err)
	}
	os.Remove(dest)
}

func TestUnpack(t *testing.T) {
	src := "../"
	dest := "./test.tar.gz"
	err := PackMatch(dest, src, nil)
	if err != nil {
		t.Error(err)
	}
	os.Remove(dest)
	err = Unpack(dest, "./tt")
	if err != nil {
		t.Error(err)
	}
	os.RemoveAll("./tt")
}
