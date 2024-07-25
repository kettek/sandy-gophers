package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	tick          int
	granules      Granules
	granulesImage *ebiten.Image
	scaleX        float64
	scaleY        float64
	raining       bool
	rain          Behavior
}

func NewGame() *Game {
	g := &Game{
		granules:      MakeGranules(320, 240),
		granulesImage: ebiten.NewImage(320, 240),
	}

	return g
}

func (g *Game) Cursor() (int, int) {
	x, y := ebiten.CursorPosition()
	x /= int(g.scaleX)
	y /= int(g.scaleY)
	return x, y
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := g.Cursor()
		g.granules.FillCircle(x, y, 3, gSandBehavior)
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		x, y := g.Cursor()
		g.granules.FillCircle(x, y, 3, gWaterBehavior)
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		x, y := g.Cursor()
		g.granules.FillCircle(x, y, 3, nil)
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		x, y := g.Cursor()
		g.granules.FillCircle(x, y, 1, gPlasticBehavior)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.raining = !g.raining
		g.rain = gWaterBehavior
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.rain = gAcidBehavior
	}

	if g.raining {
		count := g.tick % 2
		for i := 0; i < count; i++ {
			x := rand.Intn(len(g.granules[0]))
			if at := g.granules.At(x, 0); at != nil && at.behavior == nil {
				at.behavior = g.rain
			}
		}
	}

	g.tick++

	for j := len(g.granules) - 1; j >= 0; j-- {
		y := j
		/*if j%2 == 0 {
			y = len(g.granules) - 1 - j
		}*/
		if g.tick%2 == 0 {
			for x := len(g.granules[y]) - 1; x >= 0; x-- {
				(&g.granules[y][x]).Update(g)
			}
		} else {
			for x := range g.granules[y] {
				(&g.granules[y][x]).Update(g)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.granulesImage.Clear()
	for y := range g.granules {
		for x := range g.granules[y] {
			if g.granules[y][x].behavior == nil {
				continue
			}
			if g.granules[y][x].behavior == gSandBehavior {
				g.granulesImage.Set(x, y, color.NRGBA{0xff, 0xff, 0x00, 0xff})
			} else if g.granules[y][x].behavior == gWaterBehavior {
				g.granulesImage.Set(x, y, color.NRGBA{0x00, 0x00, 0xff, 0xff})
			} else if g.granules[y][x].behavior == gPlasticBehavior {
				g.granulesImage.Set(x, y, color.NRGBA{0xff, 0x00, 0xff, 0xff})
			} else if g.granules[y][x].behavior == gAcidBehavior {
				g.granulesImage.Set(x, y, color.NRGBA{0x00, 0xff, 0x00, 0xff})
			}
		}
	}

	// Draw the granules image to screen with scaling.
	op := &ebiten.DrawImageOptions{}

	// Scale the granules image to the screen size.
	op.GeoM.Scale(g.scaleX, g.scaleY)

	screen.DrawImage(g.granulesImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	g.scaleX = float64(outsideWidth) / float64(len(g.granules[0]))
	g.scaleY = float64(outsideHeight) / float64(len(g.granules))

	return outsideWidth, outsideHeight
}
