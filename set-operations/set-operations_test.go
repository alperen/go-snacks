package setoperations

import (
	"testing"

	"github.com/pjbgf/go-test/should"
)

var (
	setA = Set{1, 2, 3, 4}
	setB = Set{3, 4, 5, 6}
)

func TestFilter(t *testing.T) {
	should := should.New(t)
	setAWithoutOne := Set{2, 3, 4}

	filteredSetA := Filter(setA, func(el Element) bool {
		return el != 1
	})

	should.BeEqual(setAWithoutOne, filteredSetA, "The filter function should remove 1 key from setA")
}

func TestIntersection(t *testing.T) {
	should := should.New(t)
	intersectionedSet := Set{3, 4}

	should.BeEqual(intersectionedSet, Intersection(setA, setB), "Should intersectioned")
}

func TestUnion(t *testing.T) {
	should := should.New(t)
	unionedSet := Set{1, 2, 3, 4, 5, 6}

	should.BeEqual(unionedSet, Union(setA, setB), "Should unioned")
}
