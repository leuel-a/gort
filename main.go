package main

import (
	"bytes"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth     = 1920
	screenHeight    = 1080
	rectangleWidth  = 55
	rectangleHeight = 700
	fontSize        = 10
)

var pressStartSource *text.GoTextFaceSource
var pressStartFace *text.GoTextFace

type Game struct {
	sortingAlgorithmName       string
	numbers                    []int
	i, j                       int
	frame                      int
	speed                      int
	stopVisualization          bool
	sortingFinished            bool
	resetButton, controlButton *Button
	selectInput                *Select
}

func (game *Game) Reset() {
	game.i = 0
	game.j = 0
	game.frame = 0
	game.stopVisualization = true
	game.sortingFinished = false

	game.numbers = []int{
		12, 3, 7, 1, 15, 8, 5, 10, 2, 14,
		6, 9, 4, 11, 13, 7, 2, 5, 15, 1,
		8, 6, 3, 12, 14, 9, 11, 10, 13, 4,
	}

	game.controlButton.Text = "Start"
}

func (game *Game) Update() error {
	game.frame++
	game.selectInput.Update()

	if game.resetButton.IsClicked() {
		game.Reset()
		log.Printf("[INFO] `Reset` button pressed\n")
	}

	if game.controlButton.IsClicked() {
		if game.stopVisualization {
			game.controlButton.Text = "Stop"
		} else {
			game.controlButton.Text = "Continue"
		}

		game.stopVisualization = !game.stopVisualization
		log.Printf("[INFO] `Stop` button pressed\n")
	}

	if game.stopVisualization {
		return nil
	}

	if game.frame < game.speed {
		return nil
	}

	game.frame = 0

	bubbleSort(game)

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	if game.sortingFinished {
		VisualizeNumberArraySorted(screen, game.numbers)
	} else {
		current := game.j
		next := game.j + 1

		if next >= len(game.numbers) {
			next = len(game.numbers) - 1
		}

		VisualizeNumberArray(screen, game.numbers, current, next)
	}

	game.resetButton.Draw(screen, color.White)
	game.controlButton.Draw(screen, color.White)
	game.selectInput.Draw(screen)

	drawSortingAlgorithmName(screen, game, color.White)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func initialize(game *Game) {
	numbers := []int{
		12, 3, 7, 1, 15, 8, 5, 10, 2, 14,
		6, 9, 4, 11, 13, 7, 2, 5, 15, 1,
		8, 6, 3, 12, 14, 9, 11, 10, 13, 4,
	}

	selectOptions := []string{
		"Bubble Sort",
		"Selection Sort",
	}

	selectInput := &Select{
		X:             100,
		Y:             screenHeight - 200,
		Width:         200,
		Height:        40,
		Options:       selectOptions,
		SelectedIndex: 0,
	}

	pressStartRegularPath := "assets/fonts/PressStart2P-Regular.ttf"
	fileBytes, err := os.ReadFile(pressStartRegularPath)
	if err != nil {
		log.Fatal(err)
	}

	pressStartSource, err = text.NewGoTextFaceSource(bytes.NewReader(fileBytes))
	if err != nil {
		log.Fatal(err)
	}

	pressStartFace = &text.GoTextFace{
		Source: pressStartSource,
		Size:   fontSize,
	}

	game.numbers = numbers
	game.selectInput = selectInput
	game.sortingAlgorithmName = "Bubble Sort"
	game.resetButton = &Button{X: screenWidth - 200, Y: screenHeight - 100, Width: 100, Height: 50, Text: "Reset"}
	game.controlButton = &Button{X: screenWidth - 320, Y: screenHeight - 100, Width: 100, Height: 50, Text: "Start"}
	game.i = 0
	game.j = 0
	game.speed = 2
	game.stopVisualization = true
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Visualize Sorting Algorithms")

	game := &Game{}

	initialize(game)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
