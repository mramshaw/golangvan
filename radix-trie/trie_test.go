package trie

import "testing"

func TestInsert(t *testing.T) {
	tests := []struct {
		name          string
		value         string
		tree          Tree
		expectedCount int
		inserted      bool
	}{
		{
			name:          "insert into empty tree",
			value:         "bar",
			expectedCount: 1,
			inserted:      true,
		},
		{
			name:  "insert into tree with one element",
			value: "bar",
			tree: Tree{root: node{
				value: "foo",
			}},
			expectedCount: 2,
			inserted:      true,
		},
		{
			name:  "insert into tree existing element",
			value: "foo",
			tree: Tree{root: node{
				value: "foo",
			}},
			expectedCount: 1,
			inserted:      false,
		},
	}

	for _, test := range tests {
		inserted := test.tree.Insert(test.value)
		if inserted != test.inserted {
			t.Fatalf("test %s: expected inserted to be %t", test.name, test.inserted)
		}
		if test.tree.Count() != test.expectedCount {
			t.Fatalf("test %s: expected count to be %d, but was %d", test.name, test.expectedCount, test.tree.Count())
		}
	}
}

func TestFind(t *testing.T) {

}
