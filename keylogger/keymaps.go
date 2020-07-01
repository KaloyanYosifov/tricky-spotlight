package keylogger

type GlobalKey uint16

type KeyEventHandler struct {
	onKeyPress        func(eventHandler *KeyEventHandler)
	onKeyRelease      func(eventHandler *KeyEventHandler)
	currentActiveKeys []GlobalKey
}

const (
	KEY_1 GlobalKey = 48
	KEY_2           = 49
)

func (kEH *KeyEventHandler) keyPressed(key GlobalKey) {
	kEH.currentActiveKeys = append(kEH.currentActiveKeys, key)

	if kEH.onKeyPress != nil {
		kEH.onKeyPress(kEH)
	}
}

func (kEH *KeyEventHandler) keyReleased(key GlobalKey) {
	activeKeys := make([]GlobalKey, 0, len(kEH.currentActiveKeys))

	for _, currentKey := range kEH.currentActiveKeys {
		if currentKey == key {
			continue
		}

		activeKeys = append(activeKeys, currentKey)
	}

	kEH.currentActiveKeys = activeKeys

	if kEH.onKeyRelease != nil {
		kEH.onKeyRelease(kEH)
	}
}

func NewKeyEventHandler(onKeyPress func(eventHandler *KeyEventHandler), onKeyRelease func(eventHandler *KeyEventHandler)) *KeyEventHandler {
	return &KeyEventHandler{
		onKeyPress:        onKeyPress,
		onKeyRelease:      onKeyRelease,
		currentActiveKeys: []GlobalKey{},
	}
}
