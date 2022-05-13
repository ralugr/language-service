package repository

// Base interface for all repositories
type Base interface {
	// Read retrieves data from the source
	Read() ([]byte, error)

	// Write data o the source
	Write([]byte, bool) error

	// Close connection to the source
	Close()
}
