package diff

import (
	"github.com/skroczek/gohelper/interfaces"
	"reflect"
	"testing"
)

type String string

func (s String) Equal(v String) bool {
	return s == v
}

func TestDiff(t *testing.T) {
	type args[T interfaces.Equaler[T]] struct {
		a []T
		b []T
	}
	type testCase[T interfaces.Equaler[T]] struct {
		name        string
		args        args[T]
		wantAdded   []T
		wantRemoved []T
	}
	tests := []testCase[String]{
		{
			name: "Diff([a, b], [a, b]) == ([], [])",
			args: args[String]{
				a: []String{"a", "b"},
				b: []String{"a", "b"},
			},
			wantAdded:   nil,
			wantRemoved: nil,
		},
		{
			name: "Diff([a, b], [a]) == ([], [b])",
			args: args[String]{
				a: []String{"a", "b"},
				b: []String{"a"},
			},
			wantAdded:   nil,
			wantRemoved: []String{"b"},
		},
		{
			name: "Diff([a], [a, b]) == ([b], [])",
			args: args[String]{
				a: []String{"a"},
				b: []String{"a", "b"},
			},
			wantAdded:   []String{"b"},
			wantRemoved: nil,
		},
		{
			name: "Diff([a, b], [c, d]) == ([c, d], [a, b])",
			args: args[String]{
				a: []String{"a", "b"},
				b: []String{"c", "d"},
			},
			wantAdded:   []String{"c", "d"},
			wantRemoved: []String{"a", "b"},
		},
		{
			name: "Diff([a, b], [c, b, a]) == ([c], [])",
			args: args[String]{
				a: []String{"a", "b"},
				b: []String{"c", "b", "a"},
			},
			wantAdded:   []String{"c"},
			wantRemoved: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdded, gotRemoved := Diff(tt.args.a, tt.args.b)
			if !reflect.DeepEqual(gotAdded, tt.wantAdded) {
				t.Errorf("Diff() gotAdded = %v, want %v", gotAdded, tt.wantAdded)
			}
			if !reflect.DeepEqual(gotRemoved, tt.wantRemoved) {
				t.Errorf("Diff() gotRemoved = %v, want %v", gotRemoved, tt.wantRemoved)
			}
		})
	}
}
