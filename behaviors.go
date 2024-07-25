package main

// Behavior is an interface that defines the behavior of a granule.
type Behavior interface {
	Update(g *Game, gr *Granule)
}

// SandBehavior do be simple tho.
type SandBehavior struct{}

// Update provides basic logic for sand.
func (b SandBehavior) Update(g *Game, gr *Granule) {
	// Bail if we're at the bottom of the screen.
	if gr.y >= len(g.granules)-1 {
		return
	}

	// Swap our kinds, first attempting downwards, then left, then right.
	if next := b.GetNextOpenGranule(g, gr); next != nil {
		gr.behavior = next.behavior
		next.behavior = b
	}
}

// GetNextOpenGranule gets the next granule that we consider available. If not oob, this will be y+1, y+1 & x-1, or y+1 & x+1.
func (b SandBehavior) GetNextOpenGranule(g *Game, gr *Granule) *Granule {
	if gr.y >= len(g.granules)-1 {
		return nil
	}

	if at := g.granules.At(gr.x, gr.y+1); at != nil && (at.behavior == nil || at.behavior == gWaterBehavior) {
		return at
	}
	if at := g.granules.At(gr.x-1, gr.y+1); at != nil && (at.behavior == nil || at.behavior == gWaterBehavior) {
		return at
	}
	if at := g.granules.At(gr.x+1, gr.y+1); at != nil && (at.behavior == nil || at.behavior == gWaterBehavior) {
		return at
	}

	return nil
}

var gSandBehavior = SandBehavior{}

type WaterBehavior struct{}

func (b WaterBehavior) Update(g *Game, gr *Granule) {
	if next := b.GetNextOpenGranule(g, gr); next != nil {
		gr.behavior = nil
		next.behavior = b
	}
}

func (b WaterBehavior) GetNextOpenGranule(g *Game, gr *Granule) *Granule {
	if gr.y >= len(g.granules)-1 {
		return nil
	}

	if at := g.granules.At(gr.x, gr.y+1); at != nil && at.behavior == nil {
		return at
	}
	dir := gr.y % 2 // This one looks better, imo.
	//dir := g.tick % 2
	if dir == 0 {
		if gr.x > 0 && g.granules[gr.y][gr.x-1].behavior == nil {
			return &g.granules[gr.y][gr.x-1]
		}
		if gr.x < len(g.granules[0])-1 && g.granules[gr.y][gr.x+1].behavior == nil {
			return &g.granules[gr.y][gr.x+1]
		}
	} else {
		if gr.x < len(g.granules[0])-1 && g.granules[gr.y][gr.x+1].behavior == nil {
			return &g.granules[gr.y][gr.x+1]
		}
		if gr.x > 0 && g.granules[gr.y][gr.x-1].behavior == nil {
			return &g.granules[gr.y][gr.x-1]
		}
	}

	return nil
}

var gWaterBehavior = WaterBehavior{}

type PlasticBehavior struct{}

func (b PlasticBehavior) Update(g *Game, gr *Granule) {
	if next := b.GetNextOpenGranule(g, gr); next != nil {
		gr.behavior = next.behavior
		next.behavior = b
		next.tick = gr.tick
	}
}

func (b PlasticBehavior) GetNextOpenGranule(g *Game, gr *Granule) *Granule {
	if gr.y >= len(g.granules)-1 || gr.y <= 0 {
		return nil
	}

	// Can we float up?
	if at := g.granules.At(gr.x, gr.y-1); at != nil && at.behavior == gWaterBehavior {
		return at
	}
	if at := g.granules.At(gr.x-1, gr.y-1); at != nil && at.behavior == gWaterBehavior {
		return at
	}
	if at := g.granules.At(gr.x+1, gr.y-1); at != nil && at.behavior == gWaterBehavior {
		return at
	}
	// Can we fall?
	if at := g.granules.At(gr.x, gr.y+1); at != nil && at.behavior == nil {
		return at
	}
	// Can we spread? Note that we alternate based on Y to remove left bias. This does cause jitter, ofc.
	dir := 1
	if gr.y%2 == 0 {
		dir = -1
	}

	if at := g.granules.At(gr.x+dir, gr.y); at != nil && (at.behavior == gWaterBehavior) {
		return at
	}
	if at := g.granules.At(gr.x-dir, gr.y); at != nil && (at.behavior == gWaterBehavior) {
		return at
	}

	// Can we pile?
	if at := g.granules.At(gr.x-1, gr.y+1); at != nil && (at.behavior == nil) {
		return at
	}
	if at := g.granules.At(gr.x+1, gr.y+1); at != nil && (at.behavior == nil) {
		return at
	}

	return nil
}

var gPlasticBehavior = PlasticBehavior{}

type AcidBehavior struct{}

func (b AcidBehavior) Update(g *Game, gr *Granule) {
	if next := g.granules.At(gr.x, gr.y+1); next != nil {
		gr.behavior = nil
		if next.behavior == nil {
			next.behavior = b
		} else {
			next.behavior = nil
		}
	}
}

var gAcidBehavior = AcidBehavior{}
