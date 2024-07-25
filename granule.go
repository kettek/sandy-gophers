package main

type Granule struct {
	x, y     int
	tick     int
	behavior Behavior
}

func (gr *Granule) Update(g *Game) {
	// Bail out if we've already processed this tick.
	if gr.tick == g.tick {
		return
	}
	gr.tick++

	if gr.behavior != nil {
		gr.behavior.Update(g, gr)
	}
}
