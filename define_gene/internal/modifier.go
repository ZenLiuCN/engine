package internal

//go:generate stringer -type=Mod
type (
	Mod  uint
	Mods = Stack[Mod]
)

//go:generate stringer -type=Dir
type Dir int

const (
	ONC Dir = iota
	ENT
	EXT
)

func (m Mod) IsPointerElt() bool {
	return m == ModPointerElt
}
func (m Mod) IsChanElt() bool {
	return m == ModChanElt
}
func (m Mod) IseNamedElt() bool {
	return m == ModeNamedElt
}
func (m Mod) IsMapKey() bool {
	return m == ModMapKey
}
func (m Mod) IsMapValue() bool {
	return m == ModMapValue
}
func (m Mod) IsArrayElt() bool {
	return m == ModArrayElt
}
func (m Mod) IsSliceElt() bool {
	return m == ModSliceElt
}
func (m Mod) IsParam() bool {
	return m == ModParam
}
func (m Mod) IsResult() bool {
	return m == ModResult
}
func (m Mod) IsEmbedded() bool {
	return m == ModEmbedded
}
func (m Mod) IsMethod() bool {
	return m == ModMethod
}
func (m Mod) IsFunction() bool {
	return m == ModFunction
}
func (m Mod) IsField() bool {
	return m == ModField
}
func (m Mod) IsTuple() bool {
	return m == ModTuple
}
func (m Mod) IsEllipsis() bool {
	return m == ModEllipsis
}
func (m Mod) IsTypeParam() bool {
	return m == ModTypeParam
}
func (m Mod) IsSelect() bool {
	return m == ModSelect
}
func (m Mod) IsOperate() bool {
	return m == ModOperate
}
func (m Mod) IsOperateLeft() bool {
	return m == ModOpLeft
}
func (m Mod) IsOperateRight() bool {
	return m == ModOpRight
}

const (
	// ModPointerElt visit a star | pointer element
	ModPointerElt Mod = iota
	// ModChanElt visit a chan element
	ModChanElt
	// ModeNamedElt visit a named element
	ModeNamedElt
	// ModMapKey visit a map key
	ModMapKey
	// ModMapValue visit a map value
	ModMapValue
	// ModArrayElt visit an array element
	ModArrayElt
	// ModSliceElt visit a slice element
	ModSliceElt
	// ModParam visit a parameter
	ModParam
	// ModResult visit a result
	ModResult
	ModEmbedded
	// ModMethod mark visit a method element
	ModMethod
	// ModFunction visit a function element
	ModFunction
	// ModField visit a field
	ModField
	// ModTuple mark for visit [ast.FieldList] or [types.Tuple]
	ModTuple
	ModEllipsis
	ModTypeParam
	ModSelect
	ModOperate
	ModOpLeft
	ModOpRight
)
