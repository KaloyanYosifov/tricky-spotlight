package keylogger

type GlobalKey uint16

type KeyEventHandler struct {
	onKeyPress        func(eventHandler *KeyEventHandler)
	onKeyRelease      func(eventHandler *KeyEventHandler)
	currentActiveKeys []GlobalKey
}

const (
	// Trivial
	KEY_SPACE       GlobalKey = 57
	KEY_CTRL                  = 56
	KEY_META                  = 29
	KEY_ALT                   = 125
	KEY_LEFT_SHIFT            = 42
	KEY_RIGHT_SHIFT           = 54

	// Numbers
	KEY_0 = 11
	KEY_1 = 2
	KEY_2 = 3
	KEY_3 = 4
	KEY_4 = 5
	KEY_5 = 6
	KEY_6 = 7
	KEY_7 = 8
	KEY_8 = 9
	KEY_9 = 10

	// Alphabet Lowercase
	KEY_a = 30
	KEY_b = 48
	KEY_c = 46
	KEY_d = 32
	KEY_e = 18
	KEY_f = 33
	KEY_g = 34
	KEY_h = 35
	KEY_i = 23
	KEY_j = 36
	KEY_k = 37
	KEY_l = 38
	KEY_m = 50
	KEY_n = 49
	KEY_o = 24
	KEY_p = 25
	KEY_q = 16
	KEY_r = 19
	KEY_s = 31
	KEY_t = 20
	KEY_u = 22
	KEY_v = 47
	KEY_w = 17
	KEY_x = 45
	KEY_y = 21
	KEY_z = 44

	// Alphabet Uppercase
	KEY_A = KEY_a << KEY_LEFT_SHIFT
	KEY_B = KEY_b << KEY_LEFT_SHIFT
	KEY_C = KEY_c << KEY_LEFT_SHIFT
	KEY_D = KEY_d << KEY_LEFT_SHIFT
	KEY_E = KEY_e << KEY_LEFT_SHIFT
	KEY_F = KEY_f << KEY_LEFT_SHIFT
	KEY_G = KEY_g << KEY_LEFT_SHIFT
	KEY_H = KEY_h << KEY_LEFT_SHIFT
	KEY_I = KEY_i << KEY_LEFT_SHIFT
	KEY_J = KEY_j << KEY_LEFT_SHIFT
	KEY_K = KEY_k << KEY_LEFT_SHIFT
	KEY_L = KEY_l << KEY_LEFT_SHIFT
	KEY_M = KEY_m << KEY_LEFT_SHIFT
	KEY_N = KEY_n << KEY_LEFT_SHIFT
	KEY_O = KEY_o << KEY_LEFT_SHIFT
	KEY_P = KEY_p << KEY_LEFT_SHIFT
	KEY_Q = KEY_q << KEY_LEFT_SHIFT
	KEY_R = KEY_r << KEY_LEFT_SHIFT
	KEY_S = KEY_s << KEY_LEFT_SHIFT
	KEY_T = KEY_t << KEY_LEFT_SHIFT
	KEY_U = KEY_u << KEY_LEFT_SHIFT
	KEY_V = KEY_v << KEY_LEFT_SHIFT
	KEY_W = KEY_w << KEY_LEFT_SHIFT
	KEY_X = KEY_x << KEY_LEFT_SHIFT
	KEY_Y = KEY_y << KEY_LEFT_SHIFT
	KEY_Z = KEY_z << KEY_LEFT_SHIFT
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

func (kEH *KeyEventHandler) IsKeyActive(key GlobalKey) bool {
	for _, currentKey := range kEH.currentActiveKeys {
		if currentKey == key {
			return true
		}
	}

	return false
}

func NewKeyEventHandler(onKeyPress func(eventHandler *KeyEventHandler), onKeyRelease func(eventHandler *KeyEventHandler)) *KeyEventHandler {
	return &KeyEventHandler{
		onKeyPress:        onKeyPress,
		onKeyRelease:      onKeyRelease,
		currentActiveKeys: []GlobalKey{},
	}
}
