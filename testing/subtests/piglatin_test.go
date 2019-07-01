package main

import "testing"

// PigLatin takes an English word and translates it to Pig Latin.
// If a word starts with a vowel sound then "ay" is added to the end.
// If a word starts with a consonant or consonant cluster (qu, ch, etc) it is
// moved to the end and "ay" is added.
func PigLatin(word string) string {

	// Special case empty words so we don't panic
	if word == "" {
		return ""
	}

	switch word[0] {
	// Vowels just get an "ay" suffix
	case 'a', 'e', 'i', 'o', 'u':
		return word + "ay"

	// ch => chay
	case 'c':
		if len(word) > 1 && word[1] == 'h' {
			return word[2:] + "chay"
		}

	// qu => quay
	case 'q':
		if len(word) > 1 && word[1] == 'u' {
			return word[2:] + "quay"
		}
	}

	// Consonants move to the end plus "ay" suffix
	return word[1:] + word[0:1] + "ay"
}

func TestPigLatinTable(t *testing.T) {
	tests := []struct {
		Name, Word, Want string
	}{
		{"Empty Input", "", ""},
		{"Simple Word", "banana", "ananabay"},
		{"Vowel Word", "apple", "appleay"},
		{"Consonant Cluster qu", "quince", "incequay"},
		{"Consonant Cluster ch", "cherry", "errychay"},
	}

	for i, test := range tests {
		t.Logf("Test %d: %s", i, test.Name)

		got := PigLatin(test.Word)
		if got != test.Want {
			t.Errorf("%d: PigLatin(%q) == %q, want %q", i, test.Word, got, test.Want)
		}
	}
}

func TestPigLatinSubs(t *testing.T) {

	// test defines the inputs needed for running a PigLatin test. It returns a
	// func that can be passed to t.Run
	test := func(word, want string) func(*testing.T) {
		return func(t *testing.T) {
			got := PigLatin(word)
			if got != want {
				t.Errorf("PigLatin(%q) == %q, want %q", word, got, want)
			}
		}
	}

	t.Run("Empty Input", test("", ""))
	t.Run("Simple Word", test("banana", "ananabay"))
	t.Run("Vowel Word", test("apple", "appleay"))
	t.Run("Consonant Cluster qu", test("quince", "incequay"))
	t.Run("Consonant Cluster ch", test("cherry", "errychay"))
}

func TestPigLatinTableSubs(t *testing.T) {
	tests := []struct {
		Name, Word, Want string
	}{
		{"Empty Input", "", ""},
		{"Simple Word", "banana", "ananabay"},
		{"Vowel Word", "apple", "appleay"},
		{"Consonant Cluster qu", "quince", "incequay"},
		{"Consonant Cluster ch", "cherry", "errychay"},
	}

	for _, test := range tests {
		tf := func(t *testing.T) {
			got := PigLatin(test.Word)
			if got != test.Want {
				t.Errorf("PigLatin(%q) == %q, want %q", test.Word, got, test.Want)
			}
		}
		t.Run(test.Name, tf)
	}
}
