package scene

import "github.com/yohamta/donburi/ecs"

type ManagerOption func(m *Manager)

func WithScene(name string, scene *ecs.ECS) ManagerOption {
	return func(m *Manager) {
		if m.scenes == nil {
			m.scenes = make(Scenes)
		}

		m.scenes[name] = scene
	}
}

func WithScenes(scenes Scenes) ManagerOption {
	return func(m *Manager) {
		m.scenes = scenes
	}
}

func WithInitial(name string) ManagerOption {
	return func(m *Manager) {
		m.current = m.scenes[name]
	}
}
