package asteroids_test

import (
	"testing"
	"time"

	"github.com/tomasaschan/gosteroids/game/asteroids"
	"github.com/tomasaschan/gosteroids/game/engine"
)

func TestShipmakerCreatesShipAfterDelay(t *testing.T) {
	swarm := &engine.GameObjects{}
	swarm.Insert(asteroids.NewShipMaker())

	swarm.Update(asteroids.ShipEmergenceDelay-1*time.Millisecond, engine.NoInput, engine.NoInput)

	if len(swarm.ObjectsMatching(func(o any) bool { _, ok := o.(*asteroids.Ship); return ok })) != 0 {
		t.Fatal("ship created too early")
	}

	swarm.Update(2*time.Millisecond, engine.NoInput, engine.NoInput)

	if len(swarm.ObjectsMatching(func(o any) bool { _, ok := o.(*asteroids.Ship); return ok })) != 1 {
		t.Fatal("ship not created")
	}
}
