package engine

func Empty[T any]() (v T) {
	return
}
func EmptyRefer[T any]() *T {
	var v T
	return &v
}
func ReferOf[T any](v T, x *T) {
	*x = v
}
func UnRefer[T any](v *T) T {
	return *v
}
