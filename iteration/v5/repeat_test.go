package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	fmt.Println(repeated)
	// output "aaaaaa"
}

func TestRepeatString(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("RepeatString", func(t *testing.T) {
		got := RepeatString("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})

	t.Run("RepeatString", func(t *testing.T) {
		got := RepeatString1("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})

	t.Run("RepeatString", func(t *testing.T) {
		got := RepeatString2("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})

	t.Run("RepeatString", func(t *testing.T) {
		got := RepeatString3("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})

	t.Run("RepeatString", func(t *testing.T) {
		got := RepeatString4("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})
}
