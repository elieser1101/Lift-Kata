fn main() {
    println!("Hello, world!");
    let direction = Direction::Up;
    let call = Call {floor: 1, direction: direction};
    let callb = Call {floor: 2, direction: direction};
    let callc = Call {floor: 3, direction: direction};
    println!("{:?}", direction);
    println!("{:?}", call);
    let lifta = Lift {id: "A".into(), floor: 1, requests: vec![], doors_open: true};
    let liftb = Lift {id: "B".into(), floor: 1, requests: vec![], doors_open: true};
    let liftc = Lift {id: "C".into(), floor: 1, requests: vec![], doors_open: true};
    println!("{:?}", lifta);
    let mut system = System::new(vec![], vec![lifta], vec![call]);
    println!("{:?}", system);
    system.add_floors(vec![1,2,3]);
    let mut lifts = vec![liftb, liftc];
    system.add_lifts(&mut lifts);
    let mut calls = vec![callb, callc];
    system.add_calls(&mut calls);
    println!("{:?}", system);
}
#[derive(Debug, Copy, Clone)]
pub enum Direction {
    Up,
    Down
} 

#[derive(Debug)]
pub struct Call {
    pub floor: i32,
    pub direction: Direction,
}

#[derive(Debug)]
pub struct Lift {
	pub id:        String,
	pub floor:     i32,
	pub requests:  Vec<i32>,
	pub doors_open: bool,
}

#[derive(Debug)]
pub struct System {
	pub floors: Vec<i32>,
	pub lifts: Vec<Lift>,
	pub calls: Vec<Call>,
}


impl System {
    fn new(floors: Vec<i32>, lifts: Vec<Lift>, calls: Vec<Call>) -> System {
        System {floors, lifts, calls}
    }

    fn add_floors(&mut self, floors: Vec<i32>){
        self.floors = floors
    }

    fn add_lifts(&mut self, lifts: &mut Vec<Lift>){
        self.lifts.append(lifts)
    }

    fn add_calls(&mut self, calls: &mut Vec<Call>){
        self.calls.append(calls)
    }
}

//// CallsFor ..
//func (s System) CallsFor(floor int) (calls []Call) {
//	calls = []Call{}
//	for _, c := range s.calls {
//		if c.Floor == floor {
//			calls = append(calls, c)
//		}
//	}
//	return calls
//}
//
//// Tick ..
//func (s System) Tick() {
//	panic("Implement this method")
//}
