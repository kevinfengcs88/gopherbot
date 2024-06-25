package gopherutils

import (
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	var tests = []struct {
		name string
		got  string
		want string
	}{
		{"bang prefix", "!ping", "ping"},
		{"dot prefix", ".ping", "ping"},
		{"no prefix", "ping", "ping"},
		{"misc prefix", "#ping", "#ping"},
		{"one space", " ", " "},
		{"multiple spaces", "   ", "   "},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.got
			want := test.want

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
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
