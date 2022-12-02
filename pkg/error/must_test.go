package error

import (
	"fmt"
	"testing"
)

func TestMust(t *testing.T) {
	type args[T any] struct {
		v   T
		err error
	}
	tests := []struct {
		name  string
		args  args[interface{}]
		panic bool
	}{
		{
			name: "Must(1, nil) == 1",
			args: args[interface{}]{
				v:   1,
				err: nil,
			},
			panic: false,
		},
		{
			name: "Must(1, err) == 1",
			args: args[interface{}]{
				v:   1,
				err: fmt.Errorf("err"),
			},
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.panic {
						t.Errorf("Must() panicked: %v", r)
					}
				} else {
					if tt.panic {
						t.Errorf("Must() did not panic")
					}
				}
			}()
			_ = Must(tt.args.v, tt.args.err)
		})
	}
}
