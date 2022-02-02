package string

import (
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want bool
	}{
		{"empty stack", &Stack{}, true},
		{"empty stack", &Stack{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_String(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want string
	}{
		{"empty stack", &Stack{}, ""},
		{"non-empty stack", &Stack{104, 101, 108, 108, 111}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Stack.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		s    *Stack
		args args
		want string
	}{
		{"add char to empty stack", &Stack{}, args{104}, "h"},
		{"add char to non-empty stack", &Stack{104}, args{101}, "he"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.c)
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Stack.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name      string
		s         *Stack
		want      rune
		wantErr   bool
		wantStack string
	}{
		{"pop from empty stack", &Stack{}, invalidItem, true, ""},
		{"pop from stack with single item stack", &Stack{104}, 104, false, ""},
		{"pop from stack with multiple item stack", &Stack{104, 101}, 101, false, "h"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}

			if gotStack := tt.s.String(); gotStack != tt.wantStack {
				t.Errorf("Stack.String() = %v, want %v", gotStack, tt.want)
			}
		})
	}
}
