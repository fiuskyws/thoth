package storage

type (
	// API represents all the methods to implement Storage
	API interface {
		// Set inserts data
		Set(key, value string) error
		// Get retrieves data
		Get(key string) (string, error)
	}
)
