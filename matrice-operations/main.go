package main

import (
	"errors"
	"log"
	"time"
)

type matrice [][]int
type vector []int

var (
	errorDifferentInternalDimension error = errors.New("internal dimensions should be the same")
)

func create2dMatrice(x int, y int) matrice {
	result := make([][]int, x)

	for index := range result {
		result[index] = make([]int, y)
	}

	return result
}

func getDimensions(m1 matrice, m2 matrice) (m1X int, m1Y int, m2X int, m2Y int) {
	m1X = len(m1)
	m1Y = len(m1[0])
	m2X = len(m2)
	m2Y = len(m2[0])

	return
}

func getColumn(m1 matrice, columnIndex int) []int {
	column := make([]int, 0)

	for _, row := range m1 {
		column = append(column, row[columnIndex])
	}

	return column
}

func pairedProduct(v1 vector, v2 vector) (production int) {
	production = 0

	for index, v1Item := range v1 {
		production += v1Item * v2[index]
	}

	return
}

func matriceProduction(m1 matrice, m2 matrice) (matrice, error) {
	m1X, m1Y, m2X, m2Y := getDimensions(m1, m2)
	resultMatrice := create2dMatrice(m1X, m2Y)

	if m1Y != m2X {
		return nil, errorDifferentInternalDimension
	}

	for xIndex, x := range resultMatrice {
		for yIndex := range x {
			selectXRow := m1[xIndex]
			selectYRow := getColumn(m2, yIndex)
			pairedProduction := pairedProduct(selectXRow, selectYRow)
			resultMatrice[xIndex][yIndex] = pairedProduction
		}
	}

	return resultMatrice, nil
}

func main() {
	matrice1 := matrice{
		{0, 3, 5},
		{5, 5, 2},
	} // 2 x 3

	matrice2 := matrice{
		{3, 4},
		{3, -2},
		{4, -2},
	} // 3 x 2

	a1Time := time.Now()
	a, err := matriceProduction(matrice1, matrice2)
	if err == nil {
		log.Println(a)
	}
	a2Time := time.Now()
	log.Println("Zaman farki:", a2Time.Sub(a1Time))
}
