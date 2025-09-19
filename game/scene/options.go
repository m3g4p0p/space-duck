package scene

type ManagerOption func(m *Manager)

func WithScene(name string, factory Factory) ManagerOption {
	return func(m *Manager) {
		m.Register(name, factory)
	}
}

func WithInitial(name string) ManagerOption {
	return func(m *Manager) {
		m.Goto(name)
	}
}
