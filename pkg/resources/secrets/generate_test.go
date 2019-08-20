package secrets

import "testing"

func TestRandomAlphanumericString(t *testing.T) {
	password := RandomAlphanumericString(10)
	if len(password) != 10 {
		t.Error("Expected password of length 10, but got ", len(password))
	}
}

func TestRandomAlphanumericStringShouldBeRandom(t *testing.T) {
	x := RandomAlphanumericString(10)
	y := RandomAlphanumericString(10)
	if x == y {
		t.Errorf("Expected two unique passwords but got %s and %s", x, y)
	} else {
		t.Logf("well done")
	}
}
