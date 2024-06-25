package gopherutils

import (
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	t.Run("bang prefix", func(t *testing.T) {
		got := RemovePrefix("!ping")
		want := "ping"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("dot prefix", func(t *testing.T) {
		got := RemovePrefix(".ping")
		want := "ping"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("no prefix", func(t *testing.T) {
		got := RemovePrefix("ping")
		want := "ping"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("misc prefix", func(t *testing.T) {
		got := RemovePrefix("#ping")
		want := "#ping"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("empty input", func(t *testing.T) {
		got := RemovePrefix("")
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("one space", func(t *testing.T) {
		got := RemovePrefix(" ")
		want := " "

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("multiple spaces", func(t *testing.T) {
		got := RemovePrefix("   ")
		want := "   "

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestRedify(t *testing.T) {
	got := Redify("ERROR: The process \"SonsOfTheForestDS.exe\" not found.")
	want := "```diff\n- ERROR: The process \"SonsOfTheForestDS.exe\" not found.\n```"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreenify(t *testing.T) {
	got := Greenify("SUCCESS: The process \"SonsOfTheForestDS.exe\" with PID 24824 has been terminated.")
	want := "```diff\n+ SUCCESS: The process \"SonsOfTheForestDS.exe\" with PID 24824 has been terminated.\n```"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
