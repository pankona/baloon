package scene

import (
	"context"

	"github.com/pankona/gomo-simra/simra/fps"
)

// spawner emits spawn event in random
type spawner struct {
	eq chan *command
}

func (s *spawner) run(context context.Context) {
	for {
		select {
		case <-context.Done():
			return
		case <-fps.After(60):
			// TODO: in random
			s.eq <- newCommand(cmdBalloonSpawn, nil)
		}
	}
}

func (s *spawner) OnEvent(e interface{}) {
	_ = e.(*command)
	// TODO: implement
}
