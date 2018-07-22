package bootstap

import "sync"


// default filter execution points
const (
	BeforeStatic = iota
	BeforeRouter
	BeforeExec
	AfterExec
	FinishRouter
)

// ControllerRegister containers registered router rules, controller handlers and filters.
type ControllerRegister struct {
	routers      map[string]*Tree
	enablePolicy bool
	policies     map[string]*Tree
	enableFilter bool
	filters      [FinishRouter + 1][]*FilterRouter
	pool         sync.Pool
}
