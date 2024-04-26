package observer

import "fmt"

type Listener interface {
	Fight(Beaten Listener, Notifier Notifier)
	HandlerEvent()
	GetName() string
	GetParty() string
	Title() string
	SetSignal(e *Event)
}

const (
	PGaiBang  string = "丐帮"
	PMingJiao string = "明教"
)

type Hero struct {
	Name   string
	Party  string
	Signal chan *Event
}

func NewHero(Name string, Party string) Listener {
	return &Hero{
		Name:   Name,
		Party:  Party,
		Signal: make(chan *Event),
	}
}

func (h *Hero) Fight(Beaten Listener, bxs Notifier) {
	msg := fmt.Sprintf("%s 揍了 %s", h.Title(), Beaten.Title())
	event := Event{
		Notifier: bxs,
		Beater:   h,
		Beaten:   Beaten,
		Msg:      msg,
	}
	bxs.Notify(&event)
}

func (h *Hero) HandlerEvent() {
	fmt.Println(h)
	var signal *Event
	for {
		signal = <-h.Signal
		fmt.Printf(signal.Msg)
		// 是本人 忽略
		if signal.Beater == h || signal.Beaten == h {
			fmt.Printf("|%s 忽略\n", h.GetName())
		} else if signal.Beater.GetParty() == h.Party {
			fmt.Printf("|%s 拍手叫好\n", h.Title())
		} else {
			fmt.Printf("|%s 要反击\n", h.Title())
			//h.Fight(signal.Beater, signal.Notifier)
		}
	}
}
func (h *Hero) GetName() string {
	return h.Name
}
func (h *Hero) GetParty() string {
	return h.Party
}

func (h *Hero) Title() string {
	return fmt.Sprintf("[%s: %s]", h.GetParty(), h.GetName())
}
func (h *Hero) SetSignal(e *Event) {
	h.Signal <- e
}
