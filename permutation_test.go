package itertools_test

import (
	"testing"

	"github.com/skillian/itertools"
)

var expected = [...][3]string{
	[3]string{"a", "a", "a"},
	[3]string{"b", "a", "a"},
	[3]string{"c", "a", "a"},
	[3]string{"a", "b", "a"},
	[3]string{"b", "b", "a"},
	[3]string{"c", "b", "a"},
	[3]string{"a", "c", "a"},
	[3]string{"b", "c", "a"},
	[3]string{"c", "c", "a"},
	[3]string{"a", "a", "b"},
	[3]string{"b", "a", "b"},
	[3]string{"c", "a", "b"},
	[3]string{"a", "b", "b"},
	[3]string{"b", "b", "b"},
	[3]string{"c", "b", "b"},
	[3]string{"a", "c", "b"},
	[3]string{"b", "c", "b"},
	[3]string{"c", "c", "b"},
	[3]string{"a", "a", "c"},
	[3]string{"b", "a", "c"},
	[3]string{"c", "a", "c"},
	[3]string{"a", "b", "c"},
	[3]string{"b", "b", "c"},
	[3]string{"c", "b", "c"},
	[3]string{"a", "c", "c"},
	[3]string{"b", "c", "c"},
	[3]string{"c", "c", "c"},
}

func TestPermutation(t *testing.T) {
	t.Parallel()

	abc := itertools.Strings{"a", "b", "c"}

	perm := itertools.NewPermuter(
		&itertools.ArrayIterator{Array: &abc, Index: 0},
		&itertools.ArrayIterator{Array: &abc, Index: 0},
		&itertools.ArrayIterator{Array: &abc, Index: 0})

	var ptrs [3]*string

	results := make([][]string, 0, 27)

	for perm.Next(&ptrs[0], &ptrs[1], &ptrs[2]) {
		vals := make([]string, 3)
		vals[0] = *(ptrs[0])
		vals[1] = *(ptrs[1])
		vals[2] = *(ptrs[2])

		t.Log(vals[0], vals[1], vals[2])

		results = append(results, vals)
	}

	if len(results) != len(expected) {
		t.Fatalf("Expected %d results but got %d", len(expected), len(results))
	}

	for i, result := range results {
		expect := expected[i]
		for j, column := range result {
			if column != expect[j] {
				t.Fatalf(
					"mismatch between results[%d][%d]: "+
						"%q and expect[%d][%d]: %q",
					i, j, column, i, j, expect[j])
			}
		}
	}
}
