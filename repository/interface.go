package repository

type Base interface {
	Read() ([]byte, error)
	Write([]byte, bool) error
	Close()
}
