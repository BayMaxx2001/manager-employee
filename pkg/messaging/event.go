package messaging

type Event interface {
	Name() string
	JSON() []byte
}
