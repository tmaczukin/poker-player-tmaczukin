package player

import (
	"fmt"
	"testing"
)

func assertStrEqual(t *testing.T, input, pattern string) {
	if input == pattern {
		return
	}

	t.Log(fmt.Sprintf("`%s` is not equal to `%s`"))
	t.Fail()
}

func TestVersion(t *testing.T) {
	assertStrEqual(t, Version(), VERSION)
}
