package service

type Storage interface {
	Store([]byte)
	Get([]byte)
}
