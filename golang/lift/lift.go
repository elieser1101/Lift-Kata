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
}

func (s System)AddRequest(liftIndex int, floor int) {
	s.lifts[liftIndex].Requests = append(s.lifts[liftIndex].Requests, floor)
}

func (s System)RemoveRequest(liftIndex int, index int) {
	s.lifts[liftIndex].Requests = append(s.lifts[liftIndex].Requests[:index], s.lifts[liftIndex].Requests[index+1:]...)
}

func (s System)SetLiftFloor(liftIndex int, floor int) {
	s.lifts[liftIndex].Floor = floor
}
