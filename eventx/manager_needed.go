package eventx

type EventManagerNeeded interface {
	SetEventManager(em EventManager) error
}
