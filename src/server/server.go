package server

type (
	API interface {
		Start() error
		Close() error
	}
)
