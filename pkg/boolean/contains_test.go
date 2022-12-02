package boolean

import (
	"github.com/skroczek/gohelper/pkg/interfaces"
	"testing"
)

type String string

func (s String) Equal(v String) bool {
	return s == v
}

func TestContains(t *testing.T) {
	type args[T interfaces.Equaler[T]] struct {
		a []T
		v T
	}
	type testCase[T interfaces.Equaler[T]] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[String]{
		{
			name: "Contains([a, b], a) == true",
			args: args[String]{
				a: []String{"a", "b"},
				v: "a",
			},
			want: true,
		},
		{
			name: "Contains([a, b], b) == true",
			args: args[String]{
				a: []String{"a", "b"},
				v: "b",
			},
			want: true,
		},
		{
			name: "Contains([a, b], c) == false",
			args: args[String]{
				a: []String{"a", "b"},
				v: "c",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
