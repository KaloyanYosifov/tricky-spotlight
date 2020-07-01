package keylogger

type Keys uint16

type KeyEventHandler struct {
	onKeyPress        func(eventHandler *KeyEventHandler)
	onKeyRelease      func(eventHandler *KeyEventHandler)
	currentActiveKeys []Keys
}

const (
	KEY_1 Keys = 48
	KEY_2 Keys = 48
)

func (kEH *KeyEventHandler) keyPressed(key Keys) {
	kEH.currentActiveKeys = append(kEH.currentActiveKeys, key)
	kEH.onKeyPress(kEH)
}

func (kEH *KeyEventHandler) keyReleased(key Keys) {
	activeKeys := make([]Keys, 0, len(kEH.currentActiveKeys))

	for _, currentKey := range kEH.currentActiveKeys {
		if currentKey == key {
			continue
		}

		activeKeys = append(activeKeys, currentKey)
	}

	kEH.currentActiveKeys = activeKeys
	kEH.onKeyRelease(kEH)
}
