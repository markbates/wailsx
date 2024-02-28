package wailsx

type StateDataProvider interface {
	StateData() (StateData, error)
}
