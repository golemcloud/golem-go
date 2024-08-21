package ptr

// New is a helper for creating pointers inline
func New[T any](t T) *T { return &t }
