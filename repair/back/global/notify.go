package global

type NotifyType int

const (
	_ NotifyType = iota

	NotifyBlueprint
	NotifyImage
)

type Notify struct {
	Type      NotifyType
	Periodic  int64
	Blueprint int64
	Filename  string
}

var _ch chan Notify

func init() {
	_ch = make(chan Notify, 1000)
}

func SendNotify(item Notify) {
	_ch <- item
}

func GetChannel() chan Notify {
	return _ch
}
