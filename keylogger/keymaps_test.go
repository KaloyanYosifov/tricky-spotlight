package keylogger

import "testing"

func TestItTriggersOnKeyPressEvent(t *testing.T) {
	var keyPressEventTriggered bool

	keyEventHandler := NewKeyEventHandler(func(eventHandler *KeyEventHandler) {
		keyPressEventTriggered = true
	}, nil)

	keyEventHandler.keyPressed(KEY_1)

	if !keyPressEventTriggered {
		t.Fatal("Key press event has not been triggered!")
	}
}

func TestItHasActiveKeyWhenPressed(t *testing.T) {
	keyEventHandler := NewKeyEventHandler(nil, nil)

	if len(keyEventHandler.currentActiveKeys) > 0 {
		t.Fatal("There are initial current active keys! They shouldn't be there!")
	}

	keyEventHandler.keyPressed(KEY_1)
	const amountOfActiveKeys = 1

	if len(keyEventHandler.currentActiveKeys) != amountOfActiveKeys {
		t.Fatal("Only one key should be pressed!")
	}

	if !keyEventHandler.IsKeyActive(KEY_1) {
		t.Fatal("Current active key is not equal to the key we pressed!")
	}
}

func TestItCanReleaseAKey(t *testing.T) {
	keyEventHandler := NewKeyEventHandler(nil, nil)
	keyEventHandler.keyPressed(KEY_1)
	keyEventHandler.keyReleased(KEY_1)

	if len(keyEventHandler.currentActiveKeys) > 0 {
		t.Fatal("We have released a key, there shouldnt be anything in current active keys!")
	}

	if keyEventHandler.IsKeyActive(KEY_1) {
		t.Fatal("The key is till pressed!")
	}
}

func TestItReleasesOnlyKeyThatIsActive(t *testing.T) {
	keyEventHandler := NewKeyEventHandler(nil, nil)
	keyEventHandler.keyPressed(KEY_1)
	keyEventHandler.keyReleased(KEY_2)

	if len(keyEventHandler.currentActiveKeys) == 0 {
		t.Fatal("The key was released when it shouldn't have!")
	}
}

func TestItCanReturnTrueIfKeyCombinationExists(t *testing.T) {
	keyEventHandler := NewKeyEventHandler(nil, nil)
	keyEventHandler.keyPressed(KEY_LEFT_SHIFT)
	keyEventHandler.keyPressed(KEY_a)

	if !keyEventHandler.IsKeyCombinationActive([]GlobalKey{KEY_LEFT_SHIFT, KEY_a}) {
		t.Fatal("The key combination is not active!")
	}
}

func TestItCanReturnTrueIfKeyCombinationExistsAndTheOrderDoesntMatter(t *testing.T) {
	keyEventHandler := NewKeyEventHandler(nil, nil)
	keyEventHandler.keyPressed(KEY_LEFT_SHIFT)
	keyEventHandler.keyPressed(KEY_a)

	if !keyEventHandler.IsKeyCombinationActive([]GlobalKey{KEY_a, KEY_LEFT_SHIFT}) {
		t.Fatal("The key combination is not active!")
	}
}
