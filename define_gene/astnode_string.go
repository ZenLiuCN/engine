// Code generated by "stringer -type=AstNode"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AstNothing-0]
	_ = x[AstIdent-1]
	_ = x[AstSelectorExpr-2]
	_ = x[AstStarExpr-3]
	_ = x[AstFuncType-4]
	_ = x[AstArrayType-5]
	_ = x[AstMapType-6]
	_ = x[AstStructType-7]
	_ = x[AstChanType-8]
	_ = x[AstEllipsis-9]
}

const _AstNode_name = "AstNothingAstIdentAstSelectorExprAstStarExprAstFuncTypeAstArrayTypeAstMapTypeAstStructTypeAstChanTypeAstEllipsis"

var _AstNode_index = [...]uint8{0, 10, 18, 33, 44, 55, 67, 77, 90, 101, 112}

func (i AstNode) String() string {
	if i < 0 || i >= AstNode(len(_AstNode_index)-1) {
		return "AstNode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AstNode_name[_AstNode_index[i]:_AstNode_index[i+1]]
}
