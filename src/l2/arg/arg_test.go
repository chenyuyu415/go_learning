package arg

import (
	"testing"
)

func TestFirstArg(t *testing.T) {
	t.Log("my test")
}

func TestSecArg(t *testing.T) {
	var a, b, c int = 1, 2, 3
	t.Log(a, b, c)

}

func TestThrArg(t *testing.T) {
	var (
		a int = 1
		b int = 1
		c int = 1
	)
	t.Log(a, b, c)
}

func TestForArg(t *testing.T) {
	var a, b, c = 1, 2, 4
	t.Log(a, b, c)
}

func TestFriArg(t *testing.T) {
	var (
		a = 1
		b = 1
		c = 1
	)
	t.Log(a, b, c)
}

func TestSixArg(t *testing.T) {
	a, b, c := 2, 4, 5
	t.Log(a, b, c)
}

func TestSevArg(t *testing.T) {
	a := 1
	b := 4
	c := "ddd"
	t.Log(a, b, c)
}

func TestFib(t *testing.T) {
	var a, b, c int
	a, b = 1, 1
	for i := 0; i <= 10; i++ {
		t.Log(a, "")
		c = a + b
		a = b
		b = c
	}
	t.Log(c, " ")
}
