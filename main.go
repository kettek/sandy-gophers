package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetTPS(200)
	ebiten.SetWindowTitle("Snandy Gofers")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
