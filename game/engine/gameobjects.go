package engine

type GameObjects struct {
	objects []any
}

type Beginner interface{ BeginUpdate(dt float64) }
type Interactor interface {
	InteractWith(any)
}
type Ender interface {
	EndUpdate(dt float64, objects *GameObjects)
}

func (g *GameObjects) Insert(o ...any) {
	g.objects = append(g.objects, o...)
}

func (g *GameObjects) Remove(o any) {
	for i, x := range g.objects {
		if x == o {
			g.objects = append(g.objects[:i], g.objects[i+1:]...)
			return
		}
	}
}

func (g *GameObjects) Pairwise() [][2]any {
	results := make([][2]any, 0, len(g.objects)*(len(g.objects)-1))

	for i, o1 := range g.objects {
		for j, o2 := range g.objects {
			if i != j {
				results = append(results, [2]any{o1, o2})
			}
		}
	}

	return results
}

func (g *GameObjects) Update(dt float64) {
	for _, o := range g.objects {
		if b, ok := o.(Beginner); ok {
			b.BeginUpdate(dt)
		}

		for _, pair := range g.Pairwise() {
			if interactor, ok := pair[0].(Interactor); ok {
				interactor.InteractWith(pair[1])
			}
		}

		if e, ok := o.(Ender); ok {
			e.EndUpdate(dt, g)
		}
	}
}
