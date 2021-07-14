package behavioral

const (
	StateOnWay = iota
	StateInQueue
	StateArrived
	StateDeparted
)

type train interface {
	arrive()
	depart()
}

type PassengerTrain struct {
	mediator
	state int
}

func NewPassengerTrain(m mediator) *PassengerTrain {
	return &PassengerTrain{m, StateOnWay}
}

func (pt *PassengerTrain) arrive() {
	if !pt.mediator.RequestArrival(pt) {
		pt.state = StateInQueue
		return
	}
	pt.state = StateArrived
}

func (pt *PassengerTrain) depart() {
	pt.state = StateDeparted
	pt.mediator.NotifyAboutDepartion()
}

type mediator interface {
	RequestArrival(train) bool
	NotifyAboutDepartion()
}

type StationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func (sm *StationManager) RequestArrival(t train) bool {
	if sm.isPlatformFree {
		sm.isPlatformFree = false
		return true
	}
	sm.trainQueue = append(sm.trainQueue, t)
	return false
}

func (sm *StationManager) NotifyAboutDepartion() {
	sm.isPlatformFree = true
	if len(sm.trainQueue) > 0 {
		firstTrain := sm.trainQueue[0]
		firstTrain.arrive()
		sm.trainQueue = sm.trainQueue[1:]
	}
}
