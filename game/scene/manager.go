package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"
)

type Scenes map[string]*ecs.ECS

type Manager struct {
	current *ecs.ECS
	scenes  Scenes
}

func NewManager(initial string, scenes Scenes) *Manager {
	if scenes == nil {
		scenes = make(Scenes)
	}

	m := &Manager{scenes: scenes}
	return m
}

func (m *Manager) Update() error {
	if m.current != nil {
		m.current.Update()
		events.ProcessAllEvents(m.current.World)
	}

	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	if m.current != nil {
		m.current.Draw(screen)
	}
}
