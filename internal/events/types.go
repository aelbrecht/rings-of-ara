package events

type Event struct {
	Kind string
	Data interface{}
}

func MakeEventContainer() *EventQueue {
	return &EventQueue{
		Size:  0,
		Items: make([]Event, 1000),
	}
}

func (c *EventQueue) Add(kind string, data interface{}) {
	c.Items[c.Size] = Event{
		Kind: kind,
		Data: data,
	}
	c.Size += 1
}

type EventQueue struct {
	Size  int
	Items []Event
}
