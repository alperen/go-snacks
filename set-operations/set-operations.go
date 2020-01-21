package setoperations

// Set ...
type Set []interface{}

// Element ...
type Element interface{}

// Filter ...
func Filter(s1 Set, criteriaFunction func(element Element) bool) (result Set) {
	for _, element := range s1 {
		if criteriaFunction(element) {
			result = append(result, element)
		}
	}

	return
}

// RemoveIndex ...
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// Has ...
func Has(s1 Set, el Element) bool {
	for _, element := range s1 {
		if element == el {
			return true
		}
	}

	return false
}

// Intersection ...
func Intersection(s1 Set, s2 Set) (result Set) {
	for _, element := range s1 {
		if Has(s2, element) {
			result = append(result, element)
		}
	}

	return
}

// Union ...
func Union(s1 Set, s2 Set) (result Set) {
	result = s1

	for _, element := range s2 {
		if !Has(s1, element) {
			result = append(result, element)
		}
	}

	return
}

// Difference ... [not-implements usendim]
func Difference(s1 Set, s2 Set) (result Set) {
	result = s1

	for _, element := range s2 {
		if Has(s1, element) {

		}
	}

	return
}
