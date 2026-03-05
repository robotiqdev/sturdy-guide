package strutil

import "testing"

func TestReverse_SimpleWord(t *testing.T) {
	got := Reverse("hello")
	want := "olleh"
	if got != want {
		t.Errorf("Reverse(%q) = %q; want %q", "hello", got, want)
	}
}

func TestReverse_EmptyString(t *testing.T) {
	got := Reverse("")
	want := ""
	if got != want {
		t.Errorf("Reverse(%q) = %q; want %q", "", got, want)
	}
}

func TestReverse_SingleCharacter(t *testing.T) {
	got := Reverse("a")
	want := "a"
	if got != want {
		t.Errorf("Reverse(%q) = %q; want %q", "a", got, want)
	}
}

func TestReverse_MultibyteUTF8(t *testing.T) {
	got := Reverse("Hello, 世界")
	want := "界世 ,olleH"
	if got != want {
		t.Errorf("Reverse(%q) = %q; want %q", "Hello, 世界", got, want)
	}
}

func TestReverse_Emoji(t *testing.T) {
	got := Reverse("Go 🚀")
	want := "🚀 oG"
	if got != want {
		t.Errorf("Reverse(%q) = %q; want %q", "Go 🚀", got, want)
	}
}

func TestReverse_Palindrome(t *testing.T) {
	got := Reverse("racecar")
	want := "racecar"
	if got != want {
		t.Errorf("Reverse(%q) = %q; want %q", "racecar", got, want)
	}
}

func TestReverse_NumbersAndSymbols(t *testing.T) {
	got := Reverse("12345")
	want := "54321"
	if got != want {
		t.Errorf("Reverse(%q) = %q; want %q", "12345", got, want)
	}
}

func TestReverse_TableDriven(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple word", "hello", "olleh"},
		{"empty string", "", ""},
		{"single char", "a", "a"},
		{"multibyte CJK", "Hello, 世界", "界世 ,olleH"},
		{"emoji", "Go 🚀", "🚀 oG"},
		{"palindrome", "racecar", "racecar"},
		{"digits", "12345", "54321"},
		{"space only", " ", " "},
		{"mixed ascii and unicode", "abc界", "界cba"},
		{"multiple emoji", "🎉🎊🎈", "🎈🎊🎉"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.input)
			if got != tt.want {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
