package main

type Granules [][]Granule

func MakeGranules(width, height int) Granules {
	granules := make(Granules, height)
	for i := range granules {
		granules[i] = make([]Granule, width)
		for j := range granules[i] {
			granules[i][j].x = j
			granules[i][j].y = i
		}
	}
	return granules
}

func (g Granules) At(x, y int) *Granule {
	if y < 0 || y >= len(g) {
		return nil
	}
	if x < 0 || x >= len(g[y]) {
		return nil
	}
	return &g[y][x]
}

func (g Granules) FillCircle(x, y, r int, b Behavior) {
	for i := -r; i <= r; i++ {
		for j := -r; j <= r; j++ {
			if i*i+j*j > r*r {
				continue
			}
			if at := g.At(x+j, y+i); at != nil {
				at.behavior = b
			}
		}
	}
}
