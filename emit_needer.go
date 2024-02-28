package wailsx

type EmitNeeder interface {
	SetEmitter(emitter Emitter) error
}
