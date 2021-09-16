mod lift;

use lift::{Call, Lift, System, Direction};

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
