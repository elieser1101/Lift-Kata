package lift

// Direction ..
type Direction int

// Directions ..
const (
	Up Direction = iota
	Down
)

// Call ..
type Call struct {
	Floor     int
	Direction Direction
}

// Lift ..
type Lift struct {
	ID        string
	Floor     int
	Requests  []int
	DoorsOpen bool
}

// System ..
type System struct {
	floors []int
	lifts  []Lift
	calls  []Call
}

// NewSystem ..
func NewSystem() *System {
	return &System{floors: []int{}, lifts: []Lift{}, calls: []Call{}}
}

// AddFloors ..
func (s *System) AddFloors(floors ...int) {
	s.floors = append(s.floors, floors...)
}

// AddLifts ..
func (s *System) AddLifts(lifts ...Lift) {
	s.lifts = append(s.lifts, lifts...)
}

// AddCalls ..
func (s *System) AddCalls(calls ...Call) {
	s.calls = append(s.calls, calls...)
}

func (s *System) RemoveCall(callIndex int) {
	s.calls = append(s.calls[:callIndex], s.calls[callIndex+1:]...)
}

//need to be called by wrapper
func (s System)FullfilledCall(callIndex int) bool {
	for _, l := range s.lifts {
		if l.Floor == s.calls[callIndex].Floor {//this asumes len(s.calll) > 0
			if l.DoorsOpen {
				if len(l.Requests) > 0 {
					if s.calls[callIndex].Direction == Up && l.Floor < l.Requests[0] {//if they are in the same floor and if the requests are above 
						return true
					}
					if s.calls[callIndex].Direction == Down && l.Floor > l.Requests[0] {
						return true
					}
				} else {
					return true
				}
			}
		}
	}
	return false
}

// CallsFor ..
func (s System) CallsFor(floor int) (calls []Call) {
	calls = []Call{}
	for _, c := range s.calls {
		if c.Floor == floor {
			calls = append(calls, c)
		}
	}
	return calls
}

// Tick ..
func (s System) Tick() {
	panic("Implement this method")
    //check all lifts
    //if lift has rquest movelift(update flor and remove request)
    //
}

func (s System)AddRequest(liftIndex int, floor int) {
	s.lifts[liftIndex].Requests = append(s.lifts[liftIndex].Requests, floor)
}

func (s System)RemoveRequest(liftIndex int, index int) {
	s.lifts[liftIndex].Requests = append(s.lifts[liftIndex].Requests[:index], s.lifts[liftIndex].Requests[index+1:]...)
}

//should happen before removing request
func (s System)FullfilledRequests(liftIndex int) bool {
    //flor in requstes
    //door is open
    //then remove?

    floorsMatch := contains(s.lifts[liftIndex].Requests, s.lifts[liftIndex].Floor)
    return floorsMatch && s.lifts[liftIndex].DoorsOpen
}

func (s System)SetLiftFloor(liftIndex int, floor int) {
	s.lifts[liftIndex].Floor = floor
}

//am i using it?
func (s System)MoveLift(liftIndex int, floor int) {
	s.lifts[liftIndex].Floor = floor
}

func (s System)OpenDoors(liftIndex int) {
	s.lifts[liftIndex].DoorsOpen = true
}

func (s System)CloseDoors(liftIndex int) {
	s.lifts[liftIndex].DoorsOpen = false
}

//utils
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
