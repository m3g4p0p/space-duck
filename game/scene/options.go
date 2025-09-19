package scene

import "github.com/yohamta/donburi/ecs"

type ManagerOption func(m *Manager)

func WithScene(name string, scene *ecs.ECS) ManagerOption {
	return func(m *Manager) {
		m.Register(name, scene)
	}
}

func WithInitial(name string) ManagerOption {
	return func(m *Manager) {
		m.Goto(name)
	}
}
