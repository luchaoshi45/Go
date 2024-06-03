package observer

type Event struct {
	Notifier Notifier
	Beater   Listener
	Beaten   Listener
	Msg      string
}
