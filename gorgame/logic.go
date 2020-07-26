package gorgame

import (
	"errors"
	"math/rand"
	"strconv"
)

//Game hold game related information
type Game struct {
	Board
	numberOfGenerations int
	numberOfGreens      int
	corX                int
	corY                int
}

//Board hold informaton about the game board. It's height and width plus the grid itself.
type Board struct {
	height int
	width  int
	grid   [][]int
}

//---------------- Helper functions --------------------

//Generates a random number for the color. 1 - for green 0 - for red
func randomColor() int {
	color := rand.Intn(2)
	return color
}

//Creates a slice of integers representing a row of random green or red colors
func makeRow(lenght int) []int {
	var row []int
	for i := 0; i < lenght; i++ {
		row = append(row, randomColor())
	}
	return row
}

//Returns a slice of integers containing all the neighbors of a cell in the grid indicated by x(row) and y(column)
func getAdjacent(grid [][]int, x int, y int) []int {
	listAdjacent := []int{}
	for direction := 0; direction < 9; direction++ {
		if direction != 4 {
			nRow := x + ((direction % 3) - 1)
			nColumn := y + ((direction / 3) - 1)
			if nRow >= 0 && nRow < len(grid) && nColumn >= 0 && nColumn < len(grid[0]) {
				listAdjacent = append(listAdjacent, grid[nRow][nColumn])
			}
		}
	}
	return listAdjacent
}

//Returns the number of either green or red neighbors depending on variable value. If it's 0 - counts red neighbors
//and if it's 1 - counts green neighbors
func countColor(adjacents []int, value int) int {
	countColors := map[int]int{}
	count := 0
	for _, color := range adjacents {
		countColors[color]++
	}
	for color := range countColors {
		if color == value {
			count += countColors[color]
		}
	}
	return count
}

//Returns either 0-red or 1-green depending on the result of the rules for each green cell in the grid.
//If a cell containing a green(1) color is surrounded by 2 , 3 or 6 green cells it returns green(0). Otherwise
//it returns red(0)
func applyGreenRules(grid [][]int, x int, y int, value int) int {
	listAdjacent := getAdjacent(grid, x, y)
	greenCount := countColor(listAdjacent, value)
	if greenCount == 2 || greenCount == 3 || greenCount == 6 {
		value = 1
	} else {
		value = 0
	}
	return value
}

//Returns either 0-red or 1-green depending on the result of the rules for each red cell in the grid
//If a cell containing a red(1) color is surrounded by 3 or 6 green cells it returns green(0). Otherwise
//it returns red(0)
func applyRedRules(grid [][]int, x int, y int, value int) int {
	listAdjacent := getAdjacent(grid, x, y)
	greenCount := countColor(listAdjacent, value)
	if greenCount == 3 || greenCount == 6 {
		value = 1
	} else {
		value = 0
	}
	return value
}

//------------------- End of helper functions -------------------

//SetHeight sets the number of rows for the grid. The value should be a number between 1 and 1000.
//It returns an error if the integer entered is not valid and nil if it is.
func (b *Board) SetHeight(h int) error {
	if h < 1 || h > 1000 {
		err := errors.New("height should be a number between 1 and 1000")
		return err
	}
	b.height = h
	return nil
}

//SetWidth sets the number of columns for the grid. The value should be a number between 1 and 1000.
//It returns an error if the integer entered is not valid and nil if it is.
func (b *Board) SetWidth(w int) error {
	if w < 1 || w > 1000 {
		err := errors.New("width should be a number between 1 and 1000")
		return err
	}
	b.width = w
	return nil
}

//Height returns the value of the number of rows set for the grid.
func (b *Board) Height() int {
	return b.height
}

//Width returns the value of the number of columns set for the grid.
func (b *Board) Width() int {
	return b.width
}

//MakeGrid creates a new grid filled with random colors with set height and width
func (g *Game) MakeGrid(height int, width int) [][]int {
	grid := [][]int{}
	for i := 0; i < height; i++ {
		grid = append(grid, makeRow(width))
	}
	return grid
}

//GridToString converts, formats and shows a grid on the screen as a string
func (g *Game) GridToString(grid [][]int) string {
	stringGrid := ""
	for _, row := range grid {
		stringGrid += "\n"
		for _, cell := range row {
			stringGrid += strconv.Itoa(cell)
		}
	}
	return stringGrid
}

//SetGrid sets a grid to board type
func (b *Board) SetGrid(providedGrid [][]int) {
	b.grid = providedGrid
}

//Grid returns the grid stored in type board
func (b *Board) Grid() [][]int {
	return b.grid
}

//GridToString converts, formats and shows a grid on the screen as a string
func (b *Board) GridToString() string {
	stringGrid := ""
	for _, row := range b.grid {
		stringGrid += "\n"
		for _, cell := range row {
			stringGrid += strconv.Itoa(cell)
		}
	}
	return stringGrid
}

//NextGeneration creates a new state of the game board after applying all the rules for each cell in the grid
func (g *Game) NextGeneration() {
	newGrid := [][]int{}
	for i, row := range g.grid {
		newRow := []int{}
		for j, cell := range row {
			if cell == 1 {
				cell = applyGreenRules(g.grid, i, j, cell)
			} else if cell == 0 {
				cell = applyRedRules(g.grid, i, j, cell+1)
			}
			newRow = append(newRow, cell)
		}
		newGrid = append(newGrid, newRow)
	}
	g.grid = newGrid
}

//Play starts n number of next generations until i = n and displays returns number of times the specific cell was green
func (g *Game) Play(n int, x int, y int) {
	g.numberOfGenerations = n
	g.corX = x
	g.corY = y
	i := 0
	for i < g.numberOfGenerations {
		g.NextGeneration()
		for i, row := range g.grid {
			for j, cell := range row {
				if i == x && j == y {
					if cell == 1 {
						g.numberOfGreens++
					}
				}
			}
		}
		i++
	}
}

//NumberOfGenerations returns number of generations stored into the gae type
func (g *Game) NumberOfGenerations() int {
	return g.numberOfGenerations
}

//NumberOfGreens returns number of times the cell hass been green
func (g *Game) NumberOfGreens() int {
	return g.numberOfGreens
}

//CorX returns the row of the specific cell
func (g *Game) CorX() int {
	return g.corX
}

//CorY returns the column of the specific cell
func (g *Game) CorY() int {
	return g.corY
}
