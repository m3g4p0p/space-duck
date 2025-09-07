package util

import "github.com/yohamta/donburi"

func GetValue[T any](
	entry *donburi.Entry,
	comp *donburi.ComponentType[T],
	defaultValue T,
) T {
	if entry.HasComponent(comp) {
		return comp.GetValue(entry)
	}

	return defaultValue
}
