package engine

type GoError struct {
	Error error
}

func Safe00(f func() error) func() *GoError {
	return func() *GoError {
		err := f()
		if err != nil {
			return &GoError{Error: err}
		}
		return nil
	}
}

func Safe01[B0 any](f func() (B0, error)) func() (B0, *GoError) {
	return func() (B0, *GoError) {
		b0, err := f()
		if err != nil {
			return b0, &GoError{Error: err}
		}
		return b0, nil
	}
}

type Safe[R any] struct {
	Result R
	Err    *GoError
}

func Safe21[A0, A1, B0 any](f func(A0, A1) (B0, error)) func(A0, A1) *Safe[B0] {
	return func(a0 A0, a1 A1) *Safe[B0] {
		b0, err := f(a0, a1)
		if err != nil {
			return &Safe[B0]{Result: b0, Err: &GoError{Error: err}}
		}
		return &Safe[B0]{Result: b0}
	}
}
func Safe31[A0, A1, A2, B0 any](f func(A0, A1, A2) (B0, error)) func(A0, A1, A2) *Safe[B0] {
	return func(a0 A0, a1 A1, a2 A2) *Safe[B0] {
		b0, err := f(a0, a1, a2)
		if err != nil {
			return &Safe[B0]{Result: b0, Err: &GoError{Error: err}}
		}
		return &Safe[B0]{Result: b0}
	}
}
