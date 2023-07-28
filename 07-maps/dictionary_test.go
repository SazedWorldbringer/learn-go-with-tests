package main

import "testing"

func TestSearch(t *testing.T) {
	// instantiate the dictionary, add test word and definition
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		// instantiate the dictionary
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		// add word and definition to the dictionary
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		// add existing word with different definition to the dictionary
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		// instantiate and add word
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		// update existing definition
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		// instantiate dictionary
		dictionary := Dictionary{}

		// update word that doesn't exist in the dictionary
		word := "test"
		definition := "this is just a test"
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

/* Helper functions */

// Helper to assert definitions in the dictionary
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	// look for the word in the dictionary
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

// Helper to assert if the definition is correct
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Helper to assert if correct error messages are returned
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
