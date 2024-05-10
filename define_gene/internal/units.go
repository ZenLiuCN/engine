package internal

type Stack[T comparable] []T

func (s Stack[T]) First() (v T) {
	if s.Empty() {
		return
	}
	return s[0]
}
func (s Stack[T]) Nth(n int) T {
	if len(s) > n {
		return s[n]
	}
	var empty T
	return empty
}
func (s Stack[T]) LNth(n int) T {
	if len(s) >= n {
		return s[len(s)-n]
	}
	var empty T
	return empty
}
func (s Stack[T]) Last() (v T) {
	if s.Empty() {
		return
	}
	return s[s.Len()-1]
}
func (s Stack[T]) LastIndex(v T) int {
	if s.Empty() {
		return -1
	}
	n := s.Len()
	for i := n - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}
func (s Stack[T]) Len() int {
	return len(s)
}
func (s Stack[T]) Empty() bool {
	return len(s) == 0
}
func (s Stack[T]) MaxIdx() int {
	return len(s) - 1
}

func (s Stack[T]) NotEmpty() bool {
	return len(s) > 0
}
func (s Stack[T]) Pop() (x Stack[T], v T) {
	if len(s) == 0 {
		return
	}
	v = s[len(s)-1]
	x = s[:len(s)-1]
	return
}
func (s Stack[T]) Push(v T) (x Stack[T]) {
	x = append(s, v)
	return
}
func (s Stack[T]) Clean() (x Stack[T]) {
	x = s[0:0]
	return
}

func (s *Stack[T]) PopSitu() (v T) {
	if len(*s) == 0 {
		return
	}
	v = s.Last()
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack[T]) PushSitu(v T) {
	*s = append(*s, v)
	return
}
func (s *Stack[T]) CleanSitu() {
	*s = (*s)[0:0]
}
