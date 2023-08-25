package engine

import (
	"time"

	"github.com/faiface/pixel"
)

type GameObjects struct {
	objects []any
}

type Beginner interface{ BeginUpdate() }
type Interactor interface{ InteractWith(any) }
type Ender interface {
	EndUpdate(dt time.Duration, objects *GameObjects)
}
type Controlled interface {
	Control(pressedKeys []Key, justPressedKeys []Key)
}

func (g *GameObjects) Clear() {
	g.objects = make([]any, 0)
}

func (g *GameObjects) Insert(o ...any) {
	g.objects = append(g.objects, o...)
}

func (g *GameObjects) Remove(o any) {
	for i, x := range g.objects {
		if x == o {
			g.objects[i] = g.objects[len(g.objects)-1]
			g.objects[len(g.objects)-1] = nil
			g.objects = g.objects[:len(g.objects)-1]
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

func (g *GameObjects) Update(dt time.Duration, pressedKeys []Key, justPressedKeys []Key) {
	for _, o := range g.objects {
		if b, ok := o.(Beginner); ok {
			b.BeginUpdate()
		}

		for _, pair := range g.Pairwise() {
			if interactor, ok := pair[0].(Interactor); ok {
				interactor.InteractWith(pair[1])
			}
		}

		if e, ok := o.(Controlled); ok {
			e.Control(pressedKeys, justPressedKeys)
		}

		if e, ok := o.(Ender); ok {
			e.EndUpdate(dt, g)
		}
	}
}

func (g *GameObjects) Draw(target pixel.Target) {

	for _, o := range g.objects {
		if d, ok := o.(Drawable); ok {
			d.Draw(target)
		}
	}
}

func (g *GameObjects) ObjectsMatching(predicate func(any) bool) (results []any) {
	for _, o := range g.objects {
		if predicate(o) {
			results = append(results, o)
		}
	}

	return
}
