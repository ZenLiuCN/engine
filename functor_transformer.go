package engine

func TransformErr01[B0, C any](f func() (B0, error), t func(B0, error) (C, error)) func() (C, error) {
	return func() (C, error) {
		c, err := t(f())
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform01[B0, C any](f func() B0, t func(B0) C) func() C {
	return func() C {
		return t(f())
	}
}

func TransformErr02[B0, B1, C any](f func() (B0, B1, error), t func(B0, B1, error) (C, error)) func() (C, error) {
	return func() (C, error) {
		c, err := t(f())
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform02[B0, B1, C any](f func() (B0, B1), t func(B0, B1) C) func() C {
	return func() C {
		return t(f())
	}
}

func TransformErr03[B0, B1, B2, C any](f func() (B0, B1, B2, error), t func(B0, B1, B2, error) (C, error)) func() (C, error) {
	return func() (C, error) {
		c, err := t(f())
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform03[B0, B1, B2, C any](f func() (B0, B1, B2), t func(B0, B1, B2) C) func() C {
	return func() C {
		return t(f())
	}
}

func TransformErr04[B0, B1, B2, B3, C any](f func() (B0, B1, B2, B3, error), t func(B0, B1, B2, B3, error) (C, error)) func() (C, error) {
	return func() (C, error) {
		c, err := t(f())
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform04[B0, B1, B2, B3, C any](f func() (B0, B1, B2, B3), t func(B0, B1, B2, B3) C) func() C {
	return func() C {
		return t(f())
	}
}

func TransformErr11[A0, B0, C any](f func(A0) (B0, error), t func(B0, error) (C, error)) func(A0) (C, error) {
	return func(a0 A0) (C, error) {
		c, err := t(f(a0))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform11[A0, B0, C any](f func(A0) B0, t func(B0) C) func(A0) C {
	return func(a0 A0) C {
		return t(f(a0))
	}
}

func TransformErr12[A0, B0, B1, C any](f func(A0) (B0, B1, error), t func(B0, B1, error) (C, error)) func(A0) (C, error) {
	return func(a0 A0) (C, error) {
		c, err := t(f(a0))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform12[A0, B0, B1, C any](f func(A0) (B0, B1), t func(B0, B1) C) func(A0) C {
	return func(a0 A0) C {
		return t(f(a0))
	}
}

func TransformErr13[A0, B0, B1, B2, C any](f func(A0) (B0, B1, B2, error), t func(B0, B1, B2, error) (C, error)) func(A0) (C, error) {
	return func(a0 A0) (C, error) {
		c, err := t(f(a0))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform13[A0, B0, B1, B2, C any](f func(A0) (B0, B1, B2), t func(B0, B1, B2) C) func(A0) C {
	return func(a0 A0) C {
		return t(f(a0))
	}
}

func TransformErr14[A0, B0, B1, B2, B3, C any](f func(A0) (B0, B1, B2, B3, error), t func(B0, B1, B2, B3, error) (C, error)) func(A0) (C, error) {
	return func(a0 A0) (C, error) {
		c, err := t(f(a0))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform14[A0, B0, B1, B2, B3, C any](f func(A0) (B0, B1, B2, B3), t func(B0, B1, B2, B3) C) func(A0) C {
	return func(a0 A0) C {
		return t(f(a0))
	}
}

func TransformErr21[A0, A1, B0, C any](f func(A0, A1) (B0, error), t func(B0, error) (C, error)) func(A0, A1) (C, error) {
	return func(a0 A0, a1 A1) (C, error) {
		c, err := t(f(a0, a1))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform21[A0, A1, B0, C any](f func(A0, A1) B0, t func(B0) C) func(A0, A1) C {
	return func(a0 A0, a1 A1) C {
		return t(f(a0, a1))
	}
}

func TransformErr22[A0, A1, B0, B1, C any](f func(A0, A1) (B0, B1, error), t func(B0, B1, error) (C, error)) func(A0, A1) (C, error) {
	return func(a0 A0, a1 A1) (C, error) {
		c, err := t(f(a0, a1))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform22[A0, A1, B0, B1, C any](f func(A0, A1) (B0, B1), t func(B0, B1) C) func(A0, A1) C {
	return func(a0 A0, a1 A1) C {
		return t(f(a0, a1))
	}
}

func TransformErr23[A0, A1, B0, B1, B2, C any](f func(A0, A1) (B0, B1, B2, error), t func(B0, B1, B2, error) (C, error)) func(A0, A1) (C, error) {
	return func(a0 A0, a1 A1) (C, error) {
		c, err := t(f(a0, a1))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform23[A0, A1, B0, B1, B2, C any](f func(A0, A1) (B0, B1, B2), t func(B0, B1, B2) C) func(A0, A1) C {
	return func(a0 A0, a1 A1) C {
		return t(f(a0, a1))
	}
}

func TransformErr24[A0, A1, B0, B1, B2, B3, C any](f func(A0, A1) (B0, B1, B2, B3, error), t func(B0, B1, B2, B3, error) (C, error)) func(A0, A1) (C, error) {

	return func(a0 A0, a1 A1) (C, error) {
		c, err := t(f(a0, a1))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform24[A0, A1, B0, B1, B2, B3, C any](f func(A0, A1) (B0, B1, B2, B3), t func(B0, B1, B2, B3) C) func(A0, A1) C {
	return func(a0 A0, a1 A1) C {
		return t(f(a0, a1))
	}
}

func TransformErr31[A0, A1, A2, B0, C any](f func(A0, A1, A2) (B0, error), t func(B0, error) (C, error)) func(A0, A1, A2) (C, error) {
	return func(a0 A0, a1 A1, a2 A2) (C, error) {
		c, err := t(f(a0, a1, a2))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform31[A0, A1, A2, B0, C any](f func(A0, A1, A2) B0, t func(B0) C) func(A0, A1, A2) C {
	return func(a0 A0, a1 A1, a2 A2) C {
		return t(f(a0, a1, a2))
	}
}

func TransformErr32[A0, A1, A2, B0, B1, C any](f func(A0, A1, A2) (B0, B1, error), t func(B0, B1, error) (C, error)) func(A0, A1, A2) (C, error) {
	return func(a0 A0, a1 A1, a2 A2) (C, error) {
		c, err := t(f(a0, a1, a2))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform32[A0, A1, A2, B0, B1, C any](f func(A0, A1, A2) (B0, B1), t func(B0, B1) C) func(A0, A1, A2) C {
	return func(a0 A0, a1 A1, a2 A2) C {
		return t(f(a0, a1, a2))
	}
}

func TransformErr33[A0, A1, A2, B0, B1, B2, C any](f func(A0, A1, A2) (B0, B1, B2, error), t func(B0, B1, B2, error) (C, error)) func(A0, A1, A2) (C, error) {

	return func(a0 A0, a1 A1, a2 A2) (C, error) {
		c, err := t(f(a0, a1, a2))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform33[A0, A1, A2, B0, B1, B2, C any](f func(A0, A1, A2) (B0, B1, B2), t func(B0, B1, B2) C) func(A0, A1, A2) C {
	return func(a0 A0, a1 A1, a2 A2) C {
		return t(f(a0, a1, a2))
	}
}

func TransformErr34[A0, A1, A2, B0, B1, B2, B3, C any](f func(A0, A1, A2) (B0, B1, B2, B3, error), t func(B0, B1, B2, B3, error) (C, error)) func(A0, A1, A2) (C, error) {
	return func(a0 A0, a1 A1, a2 A2) (C, error) {
		c, err := t(f(a0, a1, a2))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform34[A0, A1, A2, B0, B1, B2, B3, C any](f func(A0, A1, A2) (B0, B1, B2, B3), t func(B0, B1, B2, B3) C) func(A0, A1, A2) C {
	return func(a0 A0, a1 A1, a2 A2) C {
		return t(f(a0, a1, a2))
	}
}

func TransformErr41[A0, A1, A2, A3, B0, C any](f func(A0, A1, A2, A3) (B0, error), t func(B0, error) (C, error)) func(A0, A1, A2, A3) (C, error) {
	return func(a0 A0, a1 A1, a2 A2, a3 A3) (C, error) {
		c, err := t(f(a0, a1, a2, a3))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform41[A0, A1, A2, A3, B0, C any](f func(A0, A1, A2, A3) B0, t func(B0) C) func(A0, A1, A2, A3) C {
	return func(a0 A0, a1 A1, a2 A2, a3 A3) C {
		return t(f(a0, a1, a2, a3))
	}
}

func TransformErr42[A0, A1, A2, A3, B0, B1, C any](f func(A0, A1, A2, A3) (B0, B1, error), t func(B0, B1, error) (C, error)) func(A0, A1, A2, A3) (C, error) {

	return func(a0 A0, a1 A1, a2 A2, a3 A3) (C, error) {
		c, err := t(f(a0, a1, a2, a3))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform42[A0, A1, A2, A3, B0, B1, C any](f func(A0, A1, A2, A3) (B0, B1), t func(B0, B1) C) func(A0, A1, A2, A3) C {
	return func(a0 A0, a1 A1, a2 A2, a3 A3) C {
		return t(f(a0, a1, a2, a3))
	}
}

func TransformErr43[A0, A1, A2, A3, B0, B1, B2, C any](f func(A0, A1, A2, A3) (B0, B1, B2, error), t func(B0, B1, B2, error) (C, error)) func(A0, A1, A2, A3) (C, error) {
	return func(a0 A0, a1 A1, a2 A2, a3 A3) (C, error) {
		c, err := t(f(a0, a1, a2, a3))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform43[A0, A1, A2, A3, B0, B1, B2, C any](f func(A0, A1, A2, A3) (B0, B1, B2), t func(B0, B1, B2) C) func(A0, A1, A2, A3) C {
	return func(a0 A0, a1 A1, a2 A2, a3 A3) C {
		return t(f(a0, a1, a2, a3))
	}
}

func TransformErr44[A0, A1, A2, A3, B0, B1, B2, B3, C any](f func(A0, A1, A2, A3) (B0, B1, B2, B3, error), t func(B0, B1, B2, B3, error) (C, error)) func(A0, A1, A2, A3) (C, error) {
	return func(a0 A0, a1 A1, a2 A2, a3 A3) (C, error) {
		c, err := t(f(a0, a1, a2, a3))
		if err != nil {
			return c, err
		}
		return c, nil
	}
}

func Transform44[A0, A1, A2, A3, B0, B1, B2, B3, C any](f func(A0, A1, A2, A3) (B0, B1, B2, B3), t func(B0, B1, B2, B3) C) func(A0, A1, A2, A3) C {
	return func(a0 A0, a1 A1, a2 A2, a3 A3) C {
		return t(f(a0, a1, a2, a3))
	}
}
