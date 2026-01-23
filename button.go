package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	X, Y, Width, Height float32
	Text                string
	WasPressed          bool
}

func (button *Button) Draw(screen *ebiten.Image, clr color.Color, fontSize float64) {
	vector.FillRect(screen, button.X, button.Y, button.Width, button.Height, clr, true)

	face := &text.GoTextFace{
		Source: poppinsSource,
		Size:   fontSize,
	}

	textW, textH := text.Measure(button.Text, face, 0)

	textX := button.X + (button.Width-float32(textW))/2
	textY := button.Y + (button.Height-float32(textH))/2

	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(textX), float64(textY))
	opts.ColorScale.ScaleWithColor(color.Black)

	text.Draw(screen, button.Text, face, opts)
}

func (button *Button) IsHovered(mx, my int) bool {
	return float32(mx) >= button.X &&
		float32(mx) <= button.X+button.Width &&
		float32(my) >= button.Y &&
		float32(my) <= button.Y+button.Height
}

func (button *Button) IsClicked() bool {
	mx, my := ebiten.CursorPosition()
	hovered := button.IsHovered(mx, my)

	pressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	clicked := hovered && !pressed && button.WasPressed

	button.WasPressed = pressed
	return clicked
}
