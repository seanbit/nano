package session

type (
	// LifetimeHandler represents a callback
	// that will be called when a session close or
	// session low-level connection broken.
	LifetimeHandler func(*Session)

	lifetime struct {
		// callbacks that emitted on session create
		onCreate []LifetimeHandler
		// callbacks that emitted on session closed
		onClosed []LifetimeHandler
	}
)

var Lifetime = &lifetime{}

// OnCreate set the Callback which will be called
func (lt *lifetime) OnCreate(h LifetimeHandler) {
	lt.onCreate = append(lt.onCreate, h)
}

func (lt *lifetime) Create(s *Session) {
	if len(lt.onCreate) < 1 {
		return
	}

	for _, h := range lt.onCreate {
		h(s)
	}
}

// OnClosed set the Callback which will be called
// when session is closed Waring: session has closed.
func (lt *lifetime) OnClosed(h LifetimeHandler) {
	lt.onClosed = append(lt.onClosed, h)
}

func (lt *lifetime) Close(s *Session) {
	if len(lt.onClosed) < 1 {
		return
	}

	for _, h := range lt.onClosed {
		h(s)
	}
}
