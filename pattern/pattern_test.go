package pattern

import (
	"testing"
)

func Test_isStartChar(t *testing.T) {
	type args struct {
		look rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid start character `{`", args{[]rune("{")[0]}, true},
		{"valid start character `[`", args{[]rune("[")[0]}, true},
		{"valid start character `(`", args{[]rune("(")[0]}, true},
		{"ending character", args{[]rune("}")[0]}, false},
		{"Unwanted character", args{[]rune("A")[0]}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStartChar(tt.args.look); got != tt.want {
				t.Errorf("isStartChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEndChar(t *testing.T) {
	type args struct {
		look rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid end character `}`", args{[]rune("}")[0]}, true},
		{"valid end character `]`", args{[]rune("]")[0]}, true},
		{"valid end character `)`", args{[]rune(")")[0]}, true},
		{"start character", args{[]rune("{")[0]}, false},
		{"Unwanted character", args{[]rune("A")[0]}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEndChar(tt.args.look); got != tt.want {
				t.Errorf("isEndChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_endingChar(t *testing.T) {
	type args struct {
		start rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"valid start character `{`", args{[]rune("{")[0]}, []rune("}")[0]},
		{"valid start character `[`", args{[]rune("[")[0]}, []rune("]")[0]},
		{"valid start character `(`", args{[]rune("(")[0]}, []rune(")")[0]},
		{"ending character", args{[]rune("}")[0]}, invalidChar},
		{"Unwanted character", args{[]rune("A")[0]}, invalidChar},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := endingChar(tt.args.start); got != tt.want {
				t.Errorf("endingChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
