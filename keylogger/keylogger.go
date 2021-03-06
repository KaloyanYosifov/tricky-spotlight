package keylogger

import (
	"github.com/KaloyanYosifov/keylogger"
)

func WaitForKeyEvents(eventHandler *KeyEventHandler) {
	keyboard := keylogger.FindKeyboardDevice()

	if len(keyboard) <= 0 {
		panic("No keyboard is found!")
		return
	}

	k, err := keylogger.New(keyboard)
	if err != nil {
		panic(err)
		return
	}
	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		if e.Type == keylogger.EvKey {
			if e.KeyPress() {
				eventHandler.keyPressed(GlobalKey(e.Code))
			}

			// if the state of key is released
			if e.KeyRelease() {
				eventHandler.keyReleased(GlobalKey(e.Code))
			}
		}
	}
}
