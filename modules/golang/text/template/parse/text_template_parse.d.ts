// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'golang/text/template/parse'{

	// @ts-ignore
	import * as fmt from 'golang/fmt'
	// @ts-ignore
	import type {int,Struct,bool,uint,uint64,complex128,Ref,map,int64,float64,error} from 'go'
	export interface ActionNode extends Struct<ActionNode>,Node,fmt.Stringer{

			nodeType:NodeType
			pos:Pos
			line:int
			pipe:Ref<PipeNode>
			string():string
			copy():Node
	}
	export interface BoolNode extends fmt.Stringer,Node,Struct<BoolNode>{

			nodeType:NodeType
			pos:Pos
			true:bool
			string():string
			copy():Node
	}
	export interface BranchNode extends Struct<BranchNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			line:int
			pipe:Ref<PipeNode>
			list:Ref<ListNode>
			elseList:Ref<ListNode>
			string():string
			copy():Node
	}
	export interface BreakNode extends Struct<BreakNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			line:int
			copy():Node
			string():string
	}
	export interface ChainNode extends Struct<ChainNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			node:Node
			field:string[]
			add(field:string):void
			string():string
			copy():Node
	}
	export interface CommandNode extends Node,Struct<CommandNode>,fmt.Stringer{

			nodeType:NodeType
			pos:Pos
			args:Node[]
			string():string
			copy():Node
	}
	export interface CommentNode extends Struct<CommentNode>,Node,fmt.Stringer{

			nodeType:NodeType
			pos:Pos
			text:string
			string():string
			copy():Node
	}
	export interface ContinueNode extends Node,fmt.Stringer,Struct<ContinueNode>{

			nodeType:NodeType
			pos:Pos
			line:int
			copy():Node
			string():string
	}
	export interface DotNode extends Struct<DotNode>,Node,fmt.Stringer{

			nodeType:NodeType
			pos:Pos
			type():NodeType
			string():string
			copy():Node
	}
	export interface FieldNode extends Struct<FieldNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			ident:string[]
			string():string
			copy():Node
	}
	export interface IdentifierNode extends Struct<IdentifierNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			ident:string
			setPos(pos:Pos):Ref<IdentifierNode>
			setTree(t:Ref<Tree>):Ref<IdentifierNode>
			string():string
			copy():Node
	}
	export interface IfNode extends Node,Struct<IfNode>,fmt.Stringer{

			branchNode:BranchNode
			copy():Node
	}
	export function isEmptyTree(n:Node):bool

	export interface ListNode extends Struct<ListNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			nodes:Node[]
			string():string
			copyList():Ref<ListNode>
			copy():Node
	}
	export interface Mode extends uint{

	}
	export function New(name:string,...funcs:map<string[],any>[]):Ref<Tree>

	export function newIdentifier(ident:string):Ref<IdentifierNode>

	export interface NilNode extends Struct<NilNode>,Node,fmt.Stringer{

			nodeType:NodeType
			pos:Pos
			type():NodeType
			string():string
			copy():Node
	}
	export interface Node{

			copy():Node
			position():Pos
			string():string
			type():NodeType
	}
	export const NodeAction:NodeType
	export const NodeBool:NodeType
	export const NodeBreak:NodeType
	export const NodeChain:NodeType
	export const NodeCommand:NodeType
	export const NodeComment:NodeType
	export const NodeContinue:NodeType
	export const NodeDot:NodeType
	export const NodeField:NodeType
	export const NodeIdentifier:NodeType
	export const NodeIf:NodeType
	export const NodeList:NodeType
	export const NodeNil:NodeType
	export const NodeNumber:NodeType
	export const NodePipe:NodeType
	export const NodeRange:NodeType
	export const NodeString:NodeType
	export const NodeTemplate:NodeType
	export const NodeText:NodeType
	export interface NodeType extends int{

	type():NodeType
	}
	export const NodeVariable:NodeType
	export const NodeWith:NodeType
	export interface NumberNode extends fmt.Stringer,Struct<NumberNode>,Node{

			nodeType:NodeType
			pos:Pos
			isInt:bool
			isUint:bool
			isFloat:bool
			isComplex:bool
			int64:int64
			uint64:uint64
			float64:float64
			complex128:complex128
			text:string
			string():string
			copy():Node
	}
	export function parse(name:string,text:string,leftDelim:string,rightDelim:string,...funcs:map<string[],any>[]):[map<string,Ref<Tree>>]

	export const ParseComments:Mode
	export interface PipeNode extends Struct<PipeNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			line:int
			isAssign:bool
			decl:Ref<VariableNode>[]
			cmds:Ref<CommandNode>[]
			string():string
			copyPipe():Ref<PipeNode>
			copy():Node
	}
	export interface Pos extends int{

	position():Pos
	}
	export interface RangeNode extends Struct<RangeNode>,fmt.Stringer,Node{

			branchNode:BranchNode
			copy():Node
	}
	export const SkipFuncCheck:Mode
	export interface StringNode extends Struct<StringNode>,Node,fmt.Stringer{

			nodeType:NodeType
			pos:Pos
			quoted:string
			text:string
			string():string
			copy():Node
	}
	export interface TemplateNode extends Struct<TemplateNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			line:int
			name:string
			pipe:Ref<PipeNode>
			string():string
			copy():Node
	}
	export interface TextNode extends Struct<TextNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			text:Uint8Array
			string():string
			copy():Node
	}
	export interface Tree extends Struct<Tree>{

			name:string
			parseName:string
			root:Ref<ListNode>
			mode:Mode
			copy():Ref<Tree>
			errorContext(n:Node):[string,string]
			parse(text:string,leftDelim:string,rightDelim:string,treeSet:map<string,Ref<Tree>>,...funcs:map<string[],any>[]):Ref<Tree>
	}
	export interface VariableNode extends Struct<VariableNode>,fmt.Stringer,Node{

			nodeType:NodeType
			pos:Pos
			ident:string[]
			string():string
			copy():Node
	}
	export interface WithNode extends Struct<WithNode>,fmt.Stringer,Node{

			branchNode:BranchNode
			copy():Node
	}
	export function emptyChainNode():ChainNode
	export function emptyRefChainNode():Ref<ChainNode>
	export function refOfChainNode(x:ChainNode,v:Ref<ChainNode>)
	export function unRefChainNode(v:Ref<ChainNode>):ChainNode
	export function emptyDotNode():DotNode
	export function emptyRefDotNode():Ref<DotNode>
	export function refOfDotNode(x:DotNode,v:Ref<DotNode>)
	export function unRefDotNode(v:Ref<DotNode>):DotNode
	export function emptyBreakNode():BreakNode
	export function emptyRefBreakNode():Ref<BreakNode>
	export function refOfBreakNode(x:BreakNode,v:Ref<BreakNode>)
	export function unRefBreakNode(v:Ref<BreakNode>):BreakNode
	export function emptyCommandNode():CommandNode
	export function emptyRefCommandNode():Ref<CommandNode>
	export function refOfCommandNode(x:CommandNode,v:Ref<CommandNode>)
	export function unRefCommandNode(v:Ref<CommandNode>):CommandNode
	export function emptyContinueNode():ContinueNode
	export function emptyRefContinueNode():Ref<ContinueNode>
	export function refOfContinueNode(x:ContinueNode,v:Ref<ContinueNode>)
	export function unRefContinueNode(v:Ref<ContinueNode>):ContinueNode
	export function emptyStringNode():StringNode
	export function emptyRefStringNode():Ref<StringNode>
	export function refOfStringNode(x:StringNode,v:Ref<StringNode>)
	export function unRefStringNode(v:Ref<StringNode>):StringNode
	export function emptyTextNode():TextNode
	export function emptyRefTextNode():Ref<TextNode>
	export function refOfTextNode(x:TextNode,v:Ref<TextNode>)
	export function unRefTextNode(v:Ref<TextNode>):TextNode
	export function emptyVariableNode():VariableNode
	export function emptyRefVariableNode():Ref<VariableNode>
	export function refOfVariableNode(x:VariableNode,v:Ref<VariableNode>)
	export function unRefVariableNode(v:Ref<VariableNode>):VariableNode
	export function emptyWithNode():WithNode
	export function emptyRefWithNode():Ref<WithNode>
	export function refOfWithNode(x:WithNode,v:Ref<WithNode>)
	export function unRefWithNode(v:Ref<WithNode>):WithNode
	export function emptyBranchNode():BranchNode
	export function emptyRefBranchNode():Ref<BranchNode>
	export function refOfBranchNode(x:BranchNode,v:Ref<BranchNode>)
	export function unRefBranchNode(v:Ref<BranchNode>):BranchNode
	export function emptyIdentifierNode():IdentifierNode
	export function emptyRefIdentifierNode():Ref<IdentifierNode>
	export function refOfIdentifierNode(x:IdentifierNode,v:Ref<IdentifierNode>)
	export function unRefIdentifierNode(v:Ref<IdentifierNode>):IdentifierNode
	export function emptyListNode():ListNode
	export function emptyRefListNode():Ref<ListNode>
	export function refOfListNode(x:ListNode,v:Ref<ListNode>)
	export function unRefListNode(v:Ref<ListNode>):ListNode
	export function emptyNumberNode():NumberNode
	export function emptyRefNumberNode():Ref<NumberNode>
	export function refOfNumberNode(x:NumberNode,v:Ref<NumberNode>)
	export function unRefNumberNode(v:Ref<NumberNode>):NumberNode
	export function emptyPipeNode():PipeNode
	export function emptyRefPipeNode():Ref<PipeNode>
	export function refOfPipeNode(x:PipeNode,v:Ref<PipeNode>)
	export function unRefPipeNode(v:Ref<PipeNode>):PipeNode
	export function emptyBoolNode():BoolNode
	export function emptyRefBoolNode():Ref<BoolNode>
	export function refOfBoolNode(x:BoolNode,v:Ref<BoolNode>)
	export function unRefBoolNode(v:Ref<BoolNode>):BoolNode
	export function emptyCommentNode():CommentNode
	export function emptyRefCommentNode():Ref<CommentNode>
	export function refOfCommentNode(x:CommentNode,v:Ref<CommentNode>)
	export function unRefCommentNode(v:Ref<CommentNode>):CommentNode
	export function emptyFieldNode():FieldNode
	export function emptyRefFieldNode():Ref<FieldNode>
	export function refOfFieldNode(x:FieldNode,v:Ref<FieldNode>)
	export function unRefFieldNode(v:Ref<FieldNode>):FieldNode
	export function emptyIfNode():IfNode
	export function emptyRefIfNode():Ref<IfNode>
	export function refOfIfNode(x:IfNode,v:Ref<IfNode>)
	export function unRefIfNode(v:Ref<IfNode>):IfNode
	export function emptyNilNode():NilNode
	export function emptyRefNilNode():Ref<NilNode>
	export function refOfNilNode(x:NilNode,v:Ref<NilNode>)
	export function unRefNilNode(v:Ref<NilNode>):NilNode
	export function emptyRangeNode():RangeNode
	export function emptyRefRangeNode():Ref<RangeNode>
	export function refOfRangeNode(x:RangeNode,v:Ref<RangeNode>)
	export function unRefRangeNode(v:Ref<RangeNode>):RangeNode
	export function emptyTemplateNode():TemplateNode
	export function emptyRefTemplateNode():Ref<TemplateNode>
	export function refOfTemplateNode(x:TemplateNode,v:Ref<TemplateNode>)
	export function unRefTemplateNode(v:Ref<TemplateNode>):TemplateNode
	export function emptyTree():Tree
	export function emptyRefTree():Ref<Tree>
	export function refOfTree(x:Tree,v:Ref<Tree>)
	export function unRefTree(v:Ref<Tree>):Tree
	export function emptyActionNode():ActionNode
	export function emptyRefActionNode():Ref<ActionNode>
	export function refOfActionNode(x:ActionNode,v:Ref<ActionNode>)
	export function unRefActionNode(v:Ref<ActionNode>):ActionNode
}