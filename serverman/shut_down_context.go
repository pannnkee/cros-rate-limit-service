package serverman

import "time"

type ShutDownContext struct {
	Chan         chan struct{}
	DeadLineTime time.Time
}

func (m *ShutDownContext) DeadLine() (deadline time.Time, ok bool) {
	deadline = m.DeadLineTime
	ok = true
	return
}

func (m *ShutDownContext) Done() <-chan struct{} {
	return m.Chan
}

func (m *ShutDownContext) Err() <-chan struct{} {
	return nil
}

func (m *ShutDownContext) Value() <-chan struct{} {
	return nil
}
