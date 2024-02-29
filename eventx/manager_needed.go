package eventx

type EventManagerNeeded interface {
	SetEventManager(em Manager) error
}
