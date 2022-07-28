package png2j

import "testing"

func TestCon2jpg(t *testing.T) {
	err := Con2jpg("./image/go.png", "./image/go.png.jpg")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestReSizeImage(t *testing.T) {
	err := ReSizeImage("./image/go.png.jpg", 948, 418, "./image/go_big.png.jpg")
	if err != nil {
		t.Errorf("%v", err)
	}
}
