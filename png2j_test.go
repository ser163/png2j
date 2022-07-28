package png2j_test

import (
	"github.com/ser163/png2j"
	"testing"
)

func TestCon2jpg(t *testing.T) {
	err := png2j.Con2jpg("./image/go.png", "./image/go.png.jpg")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestReSizeImage(t *testing.T) {
	err := png2j.ReSizeImage("./image/go.png.jpg", 948, 418, "./image/go_big.png.jpg")
	if err != nil {
		t.Errorf("%v", err)
	}
}
