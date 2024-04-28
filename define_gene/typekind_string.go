// Code generated by "stringer -type=TypeKind,FieldKind"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TypeKindStruct-1]
	_ = x[TypeKindArray-2]
	_ = x[TypeKindInterface-3]
	_ = x[TypeKindMap-4]
	_ = x[TypeKindFunc-5]
	_ = x[TypeKindIdent-6]
	_ = x[TypeKindSelector-7]
}

const _TypeKind_name = "TypeKindStructTypeKindArrayTypeKindInterfaceTypeKindMapTypeKindFuncTypeKindIdentTypeKindSelector"

var _TypeKind_index = [...]uint8{0, 14, 27, 44, 55, 67, 80, 96}

func (i TypeKind) String() string {
	i -= 1
	if i < 0 || i >= TypeKind(len(_TypeKind_index)-1) {
		return "TypeKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TypeKind_name[_TypeKind_index[i]:_TypeKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FieldKindIdent-1]
	_ = x[FieldKindStar-2]
	_ = x[FieldKindMap-3]
}

const _FieldKind_name = "FieldKindIdentFieldKindStarFieldKindMap"

var _FieldKind_index = [...]uint8{0, 14, 27, 39}

func (i FieldKind) String() string {
	i -= 1
	if i < 0 || i >= FieldKind(len(_FieldKind_index)-1) {
		return "FieldKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _FieldKind_name[_FieldKind_index[i]:_FieldKind_index[i+1]]
}
