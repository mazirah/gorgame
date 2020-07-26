package main

import (
	"bufio"
	"fmt"
	"gorgame"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Short way to error check :)
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//getIntInput reads a string from the keyboard on press of return key and converts it to an integer to be used in calculations
func getIntInput() int {
	reader := bufio.NewReader(os.Stdin)
	numberStr, err := reader.ReadString('\r')
	check(err)
	numberStr = strings.TrimSpace(numberStr)
	number, err := strconv.Atoi(numberStr)
	check(err)
	return number
}

//runWithRandomSlice allows you to customize every variable used by the gorgame logic to achieve the wanted results.
//You can choose size of the grid and it gets populated by random 0 or 1. You can chose number of generations and which
//cell to watch for the result.
func runWithRandomSlice() {
	game := gorgame.Game{}
	fmt.Print("Enter number of rows: ")
	height := getIntInput()
	game.SetHeight(height)
	fmt.Print("Enter number of columns: ")
	width := getIntInput()
	game.SetWidth(width)
	grid := game.MakeGrid(game.Height(), game.Width())
	game.SetGrid(grid)
	fmt.Print("Grid size:\n")
	fmt.Print(game.Height(), ", ", game.Width(), "\n")
	fmt.Print("Generation Zero:")
	fmt.Print(game.Board.GridToString(), "\n")
	fmt.Print("Enter number of generations: ")
	generations := getIntInput()
	fmt.Print("Enter x coordinate of the cell ")
	x := getIntInput()
	fmt.Print("Enter y coordinate of the cell: ")
	y := getIntInput()
	game.Play(generations, x, y)
	fmt.Print("Generation ", game.NumberOfGenerations(), ":")
	fmt.Print(game.Board.GridToString(), "\n")
	fmt.Print(game.CorX(), ", ", game.CorY(), ", ", game.NumberOfGenerations(), "\n")
	fmt.Print("Result: ", game.NumberOfGreens())
}

//runWithStaticSlice require you to hardcode the grid in order to use it.
//The size of the grid is derived from the grid's lenght and the the lenght of the first row.
//The function still allows you to inspect the generation zero and choose number of generations to run and which cell to watch
func runWithStaticSlice() {
	game := gorgame.Game{}
	grid := [][]int{
		{1, 0, 0, 1},
		{1, 1, 1, 1},
		{0, 1, 0, 0},
		{1, 0, 1, 0},
	}
	game.SetGrid(grid)
	height := len(grid)
	game.SetHeight(height)
	width := len(grid[0])
	game.SetWidth(width)
	fmt.Print("Grid size:\n")
	fmt.Print(game.Height(), ", ", game.Width(), "\n")
	fmt.Print("Generation Zero:")
	fmt.Print(game.Board.GridToString(), "\n")
	fmt.Print("Enter number of generations: ")
	generations := getIntInput()
	fmt.Print("Enter x coordinate of the cell ")
	x := getIntInput()
	fmt.Print("Enter y coordinate of the cell: ")
	y := getIntInput()
	game.Play(generations, x, y)
	fmt.Print("Generation ", game.NumberOfGenerations(), ":")
	fmt.Print(game.Board.GridToString(), "\n")
	fmt.Print(game.CorX(), ", ", game.CorY(), ", ", game.NumberOfGenerations(), "\n")
	fmt.Print("Result: ", game.NumberOfGreens())
}

func main() {
	//helps with generation of random number. Used by the runWithRandomSlice function.
	//Without it the number won't be random and you will end up with a slice of slices containing the same color
	rand.Seed(time.Now().UnixNano())

	//comment or uncomment the following lines in order to run the game with a static or a random set of data
	//it is prefered only one of those to be uncommented to avoid confusion when entering input data
	runWithRandomSlice()
	//runWithStaticSlice()
}
