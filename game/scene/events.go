package scene

import "github.com/yohamta/donburi/features/events"

type SceneData struct {
	Score int
}

type SwitchData struct {
	Name string
	Data SceneData
}

var (
	SwitchEvent = events.NewEventType[SwitchData]()
	EnterEvent  = events.NewEventType[SceneData]()
)
