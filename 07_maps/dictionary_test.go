package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		knownWord := "test"

		got, _ := dictionary.Search(knownWord)
		want := "this is just a test"

		assertStrings(t, got, want, knownWord)
	})


	t.Run("unknown word", func(t *testing.T) {
		unknownWord := "test2"
		_, got := dictionary.Search(unknownWord)

		assertError(t, got, ErrNotFound, unknownWord)
	})

}

func assertStrings(t *testing.T, got, want string, searched string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s', want '%s', given '%s'", got, want, searched)
	}
}

func assertError(t *testing.T, got, want error, searched string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s', want '%s', given '%s'", got, want, searched)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test3", "this is test3")

	got, err := dictionary.Search("test3")
	want := "this is test3"

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want )
	}
}