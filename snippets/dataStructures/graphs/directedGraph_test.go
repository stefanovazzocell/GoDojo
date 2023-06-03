package graphs_test

import (
	"testing"

	graphs "github.com/stefanovazzocell/GoDojo/snippets/dataStructures/graphs"
)

func TestHasPathTo(t *testing.T) {
	d := graphs.Node{"d", []*graphs.Node{nil}}
	b := graphs.Node{"b", []*graphs.Node{}}
	f := graphs.Node{"f", []*graphs.Node{}}
	c := graphs.Node{"c", []*graphs.Node{&d, &b, &f}}
	a := graphs.Node{"a", []*graphs.Node{&b, &c}}
	e := graphs.Node{"e", []*graphs.Node{&c}}
	g := graphs.Node{"g", []*graphs.Node{&e}}
	f.Childs = append(f.Childs, &g)

	testCases := []struct {
		name    string
		start   *graphs.Node
		end     *graphs.Node
		hasPath bool
	}{
		{"Empty Nodes", nil, nil, false},
		{"Empty End", &a, nil, false},
		{"Empty Start", nil, &d, false},
		{"A -> D", &a, &d, true},
		{"C -> A", &c, &a, false},
		{"A -> G", &a, &g, true},
		{"G -> A", &g, &a, false},
		{"C -> E", &c, &e, true},
	}

	for _, test := range testCases {
		actual := test.start.HasPathTo(test.end)
		if actual != test.hasPath {
			t.Errorf("Test %q failed: expected %v but got %v", test.name, test.hasPath, actual)
		}
	}
}
