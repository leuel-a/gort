package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Select struct {
	X, Y          float32
	Width, Height float32
	Options       []string
	SelectedIndex int
	Open          bool
}

func (s *Select) Update() {
	if !mouseJustPressed() {
		return
	}

	mx, my := ebiten.CursorPosition()
	x, y := float32(mx), float32(my)

	if pointInRect(x, y, s.X, s.Y, s.Width, s.Height) {
		s.Open = !s.Open
		return
	}

	if s.Open {
		for i := range s.Options {
			oy := s.Y + s.Height*float32(i+1)
			if pointInRect(x, y, s.X, oy, s.Width, s.Height) {
				s.SelectedIndex = i
				s.Open = false
				return
			}
		}
	}

	// click outside -> close the options
	s.Open = false
}

func (s *Select) Draw(screen *ebiten.Image) {
	selectFontSize := 12.0
	fontFace := &text.GoTextFace{
		Source: pressStartSource,
		Size:   selectFontSize,
	}

	vector.FillRect(
		screen,
		s.X, s.Y,
		s.Width, s.Height,
		color.RGBA{40, 40, 40, 255},
		false,
	)

	textOptions := &text.DrawOptions{}
	textOptions.LayoutOptions = text.LayoutOptions{
		PrimaryAlign:   text.AlignStart,
		SecondaryAlign: text.AlignCenter,
	}
	textOptions.GeoM.Translate(float64(s.X+8), float64(s.Y+s.Height/2))

	text.Draw(
		screen,
		s.Options[s.SelectedIndex],
		fontFace,
		textOptions,
	)

	if s.Open {
		for i, opt := range s.Options {
			oy := s.Y + s.Height*float32(i+1)

			vector.FillRect(
				screen,
				s.X, oy,
				s.Width, s.Height,
				color.RGBA{60, 60, 60, 255},
				false,
			)

			textOptions := &text.DrawOptions{}
			textOptions.GeoM.Translate(
				float64(s.X+8),
				float64(oy+s.Height/2),
			)
			textOptions.LayoutOptions = text.LayoutOptions{
				PrimaryAlign:   text.AlignStart,
				SecondaryAlign: text.AlignCenter,
			}

			text.Draw(
				screen,
				opt,
				fontFace,
				textOptions,
			)
		}
	}
}
