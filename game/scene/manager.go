package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"
)

type Scenes map[string]*ecs.ECS

type Manager struct {
	current *ecs.ECS
	scenes  Scenes
	scale   int
}

func NewManager(options ...ManagerOption) *Manager {
	m := &Manager{scale: 1, scenes: make(Scenes)}

	for _, opt := range options {
		opt(m)
	}

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

func (m *Manager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth * m.scale, outsideHeight * m.scale
}

func (m *Manager) Goto(name string) {
	if m.current != nil {
		m.current.Pause()
	}

	m.current = m.scenes[name]
	m.current.Resume()
}

func (m *Manager) Register(name string, scene *ecs.ECS) {
	m.scenes[name] = scene
	scene.Pause()
	SwitchEvent.Subscribe(scene.World, m.handleSwitchEvent)
}

func (m *Manager) handleSwitchEvent(w donburi.World, event SwitchData) {
	m.Goto(event.Name)
	EnterEvent.Publish(m.current.World, event.Data)
}
