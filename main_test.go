package main

import (
	"bytes"
	"testing"
)

func Test_processLine(t *testing.T) {
	tests := []struct {
		name string
		text string
		out  *bytes.Buffer
		want string
	}{
		{
			name: "empty",
			text: "",
			out:  &bytes.Buffer{},
			want: `
`,
		},
		{
			name: "whitespace",
			text: "  ",
			out:  &bytes.Buffer{},
			want: "  " + `
`,
		},
		{
			name: "new resource",
			text: " # module.example ",
			out:  &bytes.Buffer{},
			want: `# module.example
`,
		},
		{
			name: "no change",
			text: "    id  = \"some string\"",
			out:  &bytes.Buffer{},
			want: `    id  = "some string"
`,
		},
		{
			name: "added line",
			text: "   + id  = \"some string\"",
			out:  &bytes.Buffer{},
			want: `[32m+    id  = "some string"[0m
`,
		},
		{
			name: "removed line",
			text: "     - item = 0 -> null",
			out:  &bytes.Buffer{},
			want: `[31m-      item = 0 -> null[0m
`,
		},
		{
			name: "changed line",
			text: "  ~ \"new/version\"    = \"latest\" -> \"1.0.1\"",
			out:  &bytes.Buffer{},
			want: `[33m~   "new/version"    = [0m[31m"latest" -> [0m[32m"1.0.1"[0m
`,
		},
		{
			name: "complex changed line",
			text: "  ~ \"new/version = some -> thing\"    = \"latest\" -> \"1.0.1\"",
			out:  &bytes.Buffer{},
			want: `[33m~   "new/version = some -> thing"    = "latest" -> "1.0.1"[0m
`,
		},
		{
			name: "cannot edit in place",
			text: "     - item = 0 -> null",
			out:  &bytes.Buffer{},
			want: `[31m-      item = 0 -> null[0m
`,
		},
		{
			name: "pre-existing color",
			text: "   [32m+[0m id  = \"some string\"",
			out: &bytes.Buffer{},
			want: `[32m+    id  = "some string"[0m
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processLine(tt.text, tt.out)
			got := tt.out.String()
			if got != tt.want {
				t.Errorf("output is now aligned. \nGot:  %v, \nwant: %v", got, tt.want)
				// t.Errorf("length  \nGot:  %v, \nwant: %v", len(got), len(tt.want))
			}
		})
	}
}

func Test_cleanRawInput(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{
			name: "only whitespace",
			raw:  "     ",
			want: "     ",
		},
		{
			name: "no color with space",
			raw:  "    id  = \"some string\"   ",
			want: "    id  = \"some string\"   ",
		},
		{
			name: "partial colored",
			raw: "   [32m+[0m id  = \"some string\"",
			want: "   + id  = \"some string\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanRawInput(tt.raw); got != tt.want {
				t.Errorf("cleanRawInput(%s) \ngot:  %v, \nwant: %v", tt.name, got, tt.want)
			}
		})
	}
}
