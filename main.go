package main

import (
	"bytes"
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
	numbers []int
	i, j    int
	frame   int
	speed   int
}

func (game *Game) Update() error {
	game.frame++

	if game.frame < game.speed {
		return nil
	}

	game.frame = 0

	if game.i < len(game.numbers) {
		if game.j < len(game.numbers)-1 {
			if game.numbers[game.j] > game.numbers[game.j+1] {
				game.numbers[game.j], game.numbers[game.j+1] = game.numbers[game.j+1], game.numbers[game.j]
			}
			game.j++
		} else {
			game.j = 0
			game.i++
		}
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	current := game.j
	next := game.j + 1

	if next >= len(game.numbers) {
		next = len(game.numbers) - 1
	}

	VisualizeNumberArray(screen, game.numbers, current, next)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func initialize() {
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
}

func main() {
	initialize()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Visualize Sorting Algorithms")

	numbers := []int{
		12, 3, 7, 1, 15, 8, 5, 10, 2, 14,
		6, 9, 4, 11, 13, 7, 2, 5, 15, 1,
		8, 6, 3, 12, 14, 9, 11, 10, 13, 4,
	}
	if err := ebiten.RunGame(&Game{numbers: numbers, i: 0, j: 0, speed: 15}); err != nil {
		log.Fatal(err)
	}
}
