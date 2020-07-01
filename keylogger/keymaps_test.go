package keylogger

import "testing"

func TestItTriggersOnKeyPressEvent(t *testing.T) {
	var keyPressEventTriggered bool

	keyEventHandler := KeyEventHandler{
		onKeyPress: func(eventHandler *KeyEventHandler) {
			keyPressEventTriggered = true
		},
		onKeyRelease: func(eventHandler *KeyEventHandler) {

		},
		currentActiveKeys: make([]Keys, 0, 0),
	}

	keyEventHandler.keyPressed(KEY_1)

	if !keyPressEventTriggered {
		t.Fatal("Key press event has not been triggered!")
	}
}
