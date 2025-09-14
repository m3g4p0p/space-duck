package scene

import "github.com/hajimehoshi/ebiten/v2"

type Manager struct {
	current ebiten.Game
	scenes  map[string]ebiten.Game
}

func NewManager() *Manager {
	m := &Manager{scenes: make(map[string]ebiten.Game)}

	return m
}

func (m *Manager) Update() error {
	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {}
