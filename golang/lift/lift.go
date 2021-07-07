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
func (s System) Tick() {//RECEIVES call and request object/mock to help simulate what is coming
	for i, l := range s.lifts {
        s.checkFloorOperation(i)//open/closedor
		if len(l.Requests) > 0 {
			s.MoveLift(i)
			s.CheckFullfilledRequests(i)// should remove request fullfilled
		}
	}
}

func (s System)AddRequest(liftIndex int, floor int) {
	s.lifts[liftIndex].Requests = append(s.lifts[liftIndex].Requests, floor)
}

func (s System)RemoveRequest(liftIndex int, index int) {
    if index < len(s.lifts[liftIndex].Requests){
        s.lifts[liftIndex].Requests = append(s.lifts[liftIndex].Requests[:index], s.lifts[liftIndex].Requests[index+1:]...)
    }
}

//should happen before removing request
func (s System)FullfilledRequests(liftIndex int) bool {
    //flor in requstes
    //door is open
    //then remove?

    floorsMatch := contains(s.lifts[liftIndex].Requests, s.lifts[liftIndex].Floor)
    return floorsMatch && s.lifts[liftIndex].DoorsOpen
}

func (s System)CheckFullfilledRequests(liftIndex int) {
    if s.FullfilledRequests(liftIndex){
        s.RemoveRequest(liftIndex, 0)//Assumes there is always somethin in the request
    }
}

func (s System)SetLiftFloor(liftIndex int, floor int) {
	if !s.lifts[liftIndex].DoorsOpen {
		s.lifts[liftIndex].Floor = floor
	}
}

//am i using it?
func (s System)MoveLift(liftIndex int) {
    //get direction
    //move one step in that direction
    lift := s.lifts[liftIndex]
    if lift.Floor < lift.Requests[0] {//asuming there is always a request is not compatible with calls
        s.SetLiftFloor(liftIndex, lift.Floor + 1)
    } else if lift.Floor > lift.Requests[0] {
        s.SetLiftFloor(liftIndex, lift.Floor - 1)
    }
}

func (s System)OpenDoors(liftIndex int) {
	s.lifts[liftIndex].DoorsOpen = true
}

func (s System)CloseDoors(liftIndex int) {
	s.lifts[liftIndex].DoorsOpen = false
}

func (s System)checkFloorOperation(liftIndex int) {
    floorsMatch := contains(s.lifts[liftIndex].Requests, s.lifts[liftIndex].Floor)
	if floorsMatch {
        s.OpenDoors(liftIndex)
    } else {
        s.CloseDoors(liftIndex)
    }
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
