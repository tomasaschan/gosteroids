package engine

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
)

type GameObjects struct {
	objects []any
}

type Interactor interface{ InteractWith(any) }

type Result struct {
	RemoveSelf bool
	ClearAll   bool
	NewObjects []any
}

type Actor interface{ Act(dt time.Duration) Result }
type Controlled interface {
	Control(pressed []Key, justPressed []Key)
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

func (g *GameObjects) Update(dt time.Duration, pressedKeys []Key, justPressedKeys []Key) {
	g.interact()
	g.control(pressedKeys, justPressedKeys)
	g.act(dt)
}

func (g *GameObjects) interact() {
	for m, o := range g.objects {
		if i, ok := o.(Interactor); ok {
			for n, other := range g.objects {
				if n != m {
					i.InteractWith(other)
				}
			}
		}
	}
}

func (g *GameObjects) control(pressedKeys []Key, justPressedKeys []Key) {
	for _, o := range g.objects {
		if e, ok := o.(Controlled); ok {
			e.Control(pressedKeys, justPressedKeys)
		}
	}
}

func (g *GameObjects) act(dt time.Duration) {
	for _, o := range g.objects {
		if e, ok := o.(Actor); ok {
			result := e.Act(dt)
			if result.RemoveSelf {
				fmt.Println("removing", o)
				g.Remove(o)
			}
			if result.ClearAll {
				g.Clear()
			}
			g.Insert(result.NewObjects...)
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
