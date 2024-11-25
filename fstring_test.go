package fstring_test

import (
	"strings"
	"testing"

	"github.com/WqyJh/go-fstring"
)

func TestFormat(t *testing.T) {
	t.Parallel()

	type args struct {
		format string
		values map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
	}{
		{"1", args{"", map[string]any{}}, "arg1", ""},
		{"1", args{"{", map[string]any{}}, "", "single '{' is not allowed"},
		{"2", args{"{{", map[string]any{}}, "{", ""},
		{"3", args{"}", map[string]any{}}, "", "single '}' is not allowed"},
		{"4", args{"}}", map[string]any{}}, "}", ""},
		{"4", args{"{}", map[string]any{}}, "", "empty expression not allowed"},
		{"4", args{"{val}", map[string]any{}}, "", "args not defined"},
		{"4", args{"a={val}", map[string]any{"val": 1}}, "a=1", ""},
		{"4", args{"a= {val}", map[string]any{"val": 1}}, "a= 1", ""},
		{"4", args{"a= { val }", map[string]any{"val": 1}}, "a= 1", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := fstring.Format(tt.args.format, tt.args.values)
			if (err != nil) != (tt.wantErr != "") {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultKeyValidator(t *testing.T) {
	t.Parallel()

	type args struct {
		format string
		values map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
	}{
		{"1", args{`{"not": "an arg"}`, map[string]any{}}, `{"not": "an arg"}`, ""},
		{"2", args{`{"not": "an {val}", "val": 1}`, map[string]any{"val": 1}}, `{"not": "an 1", "val": 1}`, ""},
		{"3", args{`{"not": "an {val}", "val": 1}`, map[string]any{}}, "", "args not defined"},
		{"4", args{`{"not": "an {{val }}", "val": 1}`, map[string]any{"val": "value"}}, `{"not": "an {value}", "val": 1}`, ""},
		{"5", args{`{"not": "an {{{val}}}", "val": 1}`, map[string]any{"val": "value"}}, `{"not": "an {{value}}", "val": 1}`, ""},
		{"6", args{`{"not": "an {{{val:}}}", "val": 1}`, map[string]any{"val": "value"}}, `{"not": "an {{{val:}}}", "val": 1}`, ""},
		{"7", args{`{"not": "an {val=}", "val": 1}`, map[string]any{"val": 1}}, `{"not": "an {val=}", "val": 1}`, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := fstring.Format(tt.args.format, tt.args.values, fstring.WithKeyValidator(fstring.BasicKeyValidator))
			if (err != nil) != (tt.wantErr != "") {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %v, want %v", got, tt.want)
			}
		})
	}
}
