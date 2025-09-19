package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"
)

type (
	Factory   func() *ecs.ECS
	Factories map[string]Factory
)

type Manager struct {
	current   *ecs.ECS
	factories Factories
	scale     int
}

func NewManager(options ...ManagerOption) *Manager {
	m := &Manager{scale: 1, factories: make(Factories)}

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
		SwitchEvent.Unsubscribe(m.current.World, m.handleSwitchEvent)
	}

	m.current = m.factories[name]()
	SwitchEvent.Subscribe(m.current.World, m.handleSwitchEvent)
}

func (m *Manager) Register(name string, factory Factory) {
	m.factories[name] = factory
}

func (m *Manager) handleSwitchEvent(w donburi.World, event SwitchData) {
	m.Goto(event.Name)
	EnterEvent.Publish(m.current.World, event.Data)
}
