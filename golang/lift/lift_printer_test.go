package lift_test

import (
	"testing"

	approvaltests "github.com/approvals/go-approval-tests"
	"github.com/lift-kata/lift"
)

func TestPrintNoLifts(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddFloors(0, 1, 2, 3)
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}

func TestNoDoors(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(lift.Lift{"A", 3, []int{0}, false})
	liftSystem.AddCalls(lift.Call{1, lift.Down})
	liftSystem.AddFloors(0, 1, 2, 3)
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewSimplePrinter()))
}

func TestSimpleLiftSystem(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(
		lift.Lift{"A", 3, []int{0}, false},
		lift.Lift{"B", 2, []int{}, false},
		lift.Lift{"C", 2, []int{}, true},
		lift.Lift{"D", 0, []int{0}, false})
	liftSystem.AddCalls(lift.Call{1, lift.Down})
	liftSystem.AddFloors(0, 1, 2, 3)
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}

func TestIllegalState(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(lift.Lift{"A", 0, []int{0}, true})
	liftSystem.AddFloors(0, 1)
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}

//lift requests slice is in ordered based on the direction up/down
func TestLargeLiftSystem(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(		
		lift.Lift{"A", 0, []int{3, 5, 7}, false},
		lift.Lift{"B", 2, []int{}, true},
		lift.Lift{"C", -2, []int{-2, 0}, false},
		lift.Lift{"D", 8, []int{0, -1, -2}, true},
		lift.Lift{"SVC", 10, []int{0, -1}, false},
		lift.Lift{"F", 8, []int{}, false})
	liftSystem.AddFloors(-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	liftSystem.AddCalls(		
		lift.Call{1, lift.Down},
		lift.Call{6, lift.Down},
		lift.Call{5, lift.Up},
		lift.Call{5, lift.Down},
		lift.Call{-1, lift.Up})
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}

func TestSimpleLiftMove(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(
		lift.Lift{"A", 3, []int{0}, false},
		lift.Lift{"B", 2, []int{}, false},
		lift.Lift{"C", 2, []int{}, true},
		lift.Lift{"D", 0, []int{0}, false})
	liftSystem.AddCalls(lift.Call{1, lift.Down})
	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.SetLiftFloor(0,0)
	liftSystem.RemoveRequest(0, 0)
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}

func TestFloorRequests(t *testing.T) {
	liftSystem := lift.NewSystem()
	liftSystem.AddLifts(
		lift.Lift{"A", 3, []int{0}, false},
		lift.Lift{"B", 2, []int{}, false},
		lift.Lift{"C", 2, []int{}, true},
		lift.Lift{"D", 0, []int{0}, false})
	liftSystem.AddCalls(lift.Call{1, lift.Down})
	liftSystem.AddFloors(0, 1, 2, 3)
	liftSystem.AddRequest(0,2)
	liftSystem.AddRequest(0,1)//what would happen if this happens first, request will be unordered
	approvaltests.VerifyString(t, lift.PrintLifts(liftSystem, lift.NewPrinter()))
}
