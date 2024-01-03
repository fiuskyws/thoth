package helper

// ToPtr converts any value of type T to
// it's pointer *T.
func ToPtr[T any](v T) *T {
	return &v
}
