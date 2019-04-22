package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	}

	t.Run("basic test", func(t *testing.T) {
		count := 5
		got := Repeat("a", count)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, got)
	})

	t.Run("test 2 character 3 repetitions", func(t *testing.T) {
		count := 3
		got := Repeat("a1", count)
		expected := "a1a1a1"
		assertCorrectMessage(t, expected, got)
	})

	t.Run("test no character 6 repetitions", func(t *testing.T) {
		count := 6
		got := Repeat("", count)
		expected := ""
		assertCorrectMessage(t, expected, got)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// Output: bbb
}

func ExampleRepeat2() {
	repeated := Repeat("d ", 7)
	fmt.Println(repeated)
	// Output: d d d d d d d
}
