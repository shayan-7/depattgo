package behavioral

type Observer interface {
	Update()
	GetID() int
}

type EventManager struct {
	observers []Observer
}

func NewEventManager() *EventManager {
	return &EventManager{make([]Observer, 0)}
}

func (ep *EventManager) Register(o Observer) {
	ep.observers = append(ep.observers, o)
}

func (ep *EventManager) UnRegister(o Observer) {
	l := len(ep.observers)
	for i, v := range ep.observers {
		if v.GetID() == o.GetID() {
			ep.observers[l-1], ep.observers[i] = v, ep.observers[l-1]
			ep.observers = ep.observers[:l-1]
			return
		}
	}
}

func (ep *EventManager) Notify() {
	for _, o := range ep.observers {
		o.Update()
	}
}

type Publisher struct {
	eventManager *EventManager
}

func (p *Publisher) Do() {
	p.eventManager.Notify()
}

func NewPublisher(em *EventManager) *Publisher {
	return &Publisher{em}
}

type Listener struct {
	ID                int
	NumOfNotification int
}

func NewListener(id int) *Listener {
	return &Listener{id, 0}
}

func (l *Listener) Update() {
	l.NumOfNotification++
}

func (l *Listener) GetID() int {
	return l.ID
}
