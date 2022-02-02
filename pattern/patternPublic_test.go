// Package pattern_test will test all the public functions of `pattern` package
package pattern_test

import (
	"testing"

	"github.com/prasnitt/stack-via-golang/pattern"
)

func TestIsValid(t *testing.T) {
	type args struct {
		pat string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty patter must be valid", args{""}, true},
		{"single start char", args{"{"}, false},
		{"single end char", args{"}"}, false},
		{"end before start", args{"}{"}, false},
		{"valid single start end", args{"{}"}, true},
		{"valid multiple start end", args{"{}()[]()"}, true},
		{"valid nested start end", args{"{{[]()}}({})[]()"}, true},
		{"valid nested start end with unwanted characters", args{"{A{[B](C)}}E(F{})G[](H)"}, true},

		// Invalid examples
		{"invalid example-1 given in problem", args{"([)]"}, false},
		{"invalid example-2 given in problem", args{"{[}"}, false},

		// Valid examples
		{"valid example-1 given in problem", args{"([])"}, true},
		{"valid example-2 given in problem", args{"([]{}())"}, true},
		{"valid example-3 given in problem", args{"(){}()"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pattern.IsValid(tt.args.pat); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
