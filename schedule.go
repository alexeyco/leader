package leader

import "time"

// Task command to be executed.
type Task struct {
	Name string
	Func func() error
}

// Schedule
func Schedule(leader Leader, interval *time.Duration) {

}
