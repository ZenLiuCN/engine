// Code generated by "stringer -type=Kind"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StructType-0]
	_ = x[MapType-1]
	_ = x[PrimitiveType-2]
	_ = x[InterfaceType-3]
	_ = x[ArrayType-4]
	_ = x[AliasType-5]
	_ = x[ChanType-6]
	_ = x[FuncType-7]
	_ = x[MethodDecl-8]
	_ = x[FuncDecl-9]
	_ = x[IdentTypeFunc-10]
	_ = x[IdentType-11]
}

const _Kind_name = "StructTypeMapTypePrimitiveTypeInterfaceTypeArrayTypeAliasTypeChanTypeFuncTypeMethodDeclFuncDeclIdentTypeFuncIdentType"

var _Kind_index = [...]uint8{0, 10, 17, 30, 43, 52, 61, 69, 77, 87, 95, 108, 117}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
