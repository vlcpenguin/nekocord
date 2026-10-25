package util

func ValueOrNil[T any](pointer *T) any {
	if pointer == nil {
		return "<nil>"
	}
	return *pointer
}
