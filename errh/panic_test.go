package errh

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name  string
		args  args
		panic bool
	}{
		{
			name:  "nil",
			args:  args{err: nil},
			panic: false,
		},
		{
			name:  "not nil",
			args:  args{err: fmt.Errorf("errh")},
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.panic {
						t.Errorf("Panic() = %v, want %v", r, nil)
					}
				} else {
					if tt.panic {
						t.Errorf("Panic() = %v, want %v", nil, "panic")
					}
				}
			}()
			Panic(tt.args.err)
		})
	}
}
