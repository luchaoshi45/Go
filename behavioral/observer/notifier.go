package observer

type Notifier interface {
	AddListener(l Listener)
	RemoveListener(l Listener)
	Notify(event *Event)
}

type BaiXiaoSheng struct {
	ListenerList []Listener
}

func NewBaiXiaoSheng() Notifier {
	return new(BaiXiaoSheng)
}

func (bxs *BaiXiaoSheng) AddListener(l Listener) {
	bxs.ListenerList = append(bxs.ListenerList, l)
}

func (bxs *BaiXiaoSheng) RemoveListener(l Listener) {
	for i, lis := range bxs.ListenerList {
		if l == lis {
			bxs.ListenerList = append(bxs.ListenerList[:i], bxs.ListenerList[i+1:]...)
			break
		}
	}
}

func (bxs *BaiXiaoSheng) Notify(event *Event) {
	for _, lis := range bxs.ListenerList {
		lis.SetSignal(event)
	}
}
