// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/board/v1'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {error,Ref,Struct,bool,float64,int} from 'go'
	export interface CompositeShape extends Struct<CompositeShape>{

			type:Ref<string>
	}
	export interface CompositeShapeBuilder extends Struct<CompositeShapeBuilder>{

			type(type_:string):Ref<CompositeShapeBuilder>
			build():Ref<CompositeShape>
	}
	export interface Connector extends Struct<Connector>{

			startObject:Ref<ConnectorAttachedObject>
			endObject:Ref<ConnectorAttachedObject>
			captions:Ref<ConnectorCaption>
	}
	export interface ConnectorAttachedObject extends Struct<ConnectorAttachedObject>{

			id:Ref<string>
	}
	export interface ConnectorAttachedObjectBuilder extends Struct<ConnectorAttachedObjectBuilder>{

			id(id:string):Ref<ConnectorAttachedObjectBuilder>
			build():Ref<ConnectorAttachedObject>
	}
	export interface ConnectorBuilder extends Struct<ConnectorBuilder>{

			startObject(startObject:Ref<ConnectorAttachedObject>):Ref<ConnectorBuilder>
			endObject(endObject:Ref<ConnectorAttachedObject>):Ref<ConnectorBuilder>
			captions(captions:Ref<ConnectorCaption>):Ref<ConnectorBuilder>
			build():Ref<Connector>
	}
	export interface ConnectorCaption extends Struct<ConnectorCaption>{

			data:Ref<Text>[]
	}
	export interface ConnectorCaptionBuilder extends Struct<ConnectorCaptionBuilder>{

			data(data:Ref<Text>[]):Ref<ConnectorCaptionBuilder>
			build():Ref<ConnectorCaption>
	}
	export interface DepartmentId extends Struct<DepartmentId>{

			departmentId:Ref<string>
			openDepartmentId:Ref<string>
	}
	export interface DepartmentIdBuilder extends Struct<DepartmentIdBuilder>{

			departmentId(departmentId:string):Ref<DepartmentIdBuilder>
			openDepartmentId(openDepartmentId:string):Ref<DepartmentIdBuilder>
			build():Ref<DepartmentId>
	}
	export interface Image extends Struct<Image>{

			token:Ref<string>
	}
	export interface ImageBuilder extends Struct<ImageBuilder>{

			token(token:string):Ref<ImageBuilder>
			build():Ref<Image>
	}
	export interface ListWhiteboardNodeReq extends Struct<ListWhiteboardNodeReq>{

	}
	export interface ListWhiteboardNodeReqBuilder extends Struct<ListWhiteboardNodeReqBuilder>{

			whiteboardId(whiteboardId:string):Ref<ListWhiteboardNodeReqBuilder>
			build():Ref<ListWhiteboardNodeReq>
	}
	export interface ListWhiteboardNodeResp extends Struct<ListWhiteboardNodeResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<ListWhiteboardNodeRespData>
			success():bool
	}
	export interface ListWhiteboardNodeRespData extends Struct<ListWhiteboardNodeRespData>{

			nodes:Ref<WhiteboardNode>[]
	}
	export interface MindMap extends Struct<MindMap>{

			parentId:Ref<string>
	}
	export interface MindMapBuilder extends Struct<MindMapBuilder>{

			parentId(parentId:string):Ref<MindMapBuilder>
			build():Ref<MindMap>
	}
	export function New(config:Ref<larkcore.Config>):Ref<V1>

	export function newCompositeShapeBuilder():Ref<CompositeShapeBuilder>

	export function newConnectorAttachedObjectBuilder():Ref<ConnectorAttachedObjectBuilder>

	export function newConnectorBuilder():Ref<ConnectorBuilder>

	export function newConnectorCaptionBuilder():Ref<ConnectorCaptionBuilder>

	export function newDepartmentIdBuilder():Ref<DepartmentIdBuilder>

	export function newImageBuilder():Ref<ImageBuilder>

	export function newListWhiteboardNodeReqBuilder():Ref<ListWhiteboardNodeReqBuilder>

	export function newMindMapBuilder():Ref<MindMapBuilder>

	export function newSectionBuilder():Ref<SectionBuilder>

	export function newStyleBuilder():Ref<StyleBuilder>

	export function newTableBuilder():Ref<TableBuilder>

	export function newTableCellBuilder():Ref<TableCellBuilder>

	export function newTableCellMergeInfoBuilder():Ref<TableCellMergeInfoBuilder>

	export function newTableMetaBuilder():Ref<TableMetaBuilder>

	export function newTextBuilder():Ref<TextBuilder>

	export function newWhiteboardNodeBuilder():Ref<WhiteboardNodeBuilder>

	export interface Section extends Struct<Section>{

			title:Ref<string>
	}
	export interface SectionBuilder extends Struct<SectionBuilder>{

			title(title:string):Ref<SectionBuilder>
			build():Ref<Section>
	}
	export interface Style extends Struct<Style>{

			fillOpacity:Ref<float64>
			borderStyle:Ref<string>
			borderWidth:Ref<string>
			borderOpacity:Ref<float64>
			hFlip:Ref<bool>
			vFlip:Ref<bool>
	}
	export interface StyleBuilder extends Struct<StyleBuilder>{

			fillOpacity(fillOpacity:float64):Ref<StyleBuilder>
			borderStyle(borderStyle:string):Ref<StyleBuilder>
			borderWidth(borderWidth:string):Ref<StyleBuilder>
			borderOpacity(borderOpacity:float64):Ref<StyleBuilder>
			hFlip(hFlip:bool):Ref<StyleBuilder>
			vFlip(vFlip:bool):Ref<StyleBuilder>
			build():Ref<Style>
	}
	export interface Table extends Struct<Table>{

			meta:Ref<TableMeta>
			title:Ref<string>
			cells:Ref<TableCell>[]
	}
	export interface TableBuilder extends Struct<TableBuilder>{

			meta(meta:Ref<TableMeta>):Ref<TableBuilder>
			title(title:string):Ref<TableBuilder>
			cells(cells:Ref<TableCell>[]):Ref<TableBuilder>
			build():Ref<Table>
	}
	export interface TableCell extends Struct<TableCell>{

			rowIndex:Ref<int>
			colIndex:Ref<int>
			mergeInfo:Ref<TableCellMergeInfo>
			children:string[]
			text:Ref<Text>
	}
	export interface TableCellBuilder extends Struct<TableCellBuilder>{

			rowIndex(rowIndex:int):Ref<TableCellBuilder>
			colIndex(colIndex:int):Ref<TableCellBuilder>
			mergeInfo(mergeInfo:Ref<TableCellMergeInfo>):Ref<TableCellBuilder>
			children(children:string[]):Ref<TableCellBuilder>
			text(text:Ref<Text>):Ref<TableCellBuilder>
			build():Ref<TableCell>
	}
	export interface TableCellMergeInfo extends Struct<TableCellMergeInfo>{

			rowSpan:Ref<int>
			colSpan:Ref<int>
	}
	export interface TableCellMergeInfoBuilder extends Struct<TableCellMergeInfoBuilder>{

			rowSpan(rowSpan:int):Ref<TableCellMergeInfoBuilder>
			colSpan(colSpan:int):Ref<TableCellMergeInfoBuilder>
			build():Ref<TableCellMergeInfo>
	}
	export interface TableMeta extends Struct<TableMeta>{

			rowNum:Ref<int>
			colNum:Ref<int>
	}
	export interface TableMetaBuilder extends Struct<TableMetaBuilder>{

			rowNum(rowNum:int):Ref<TableMetaBuilder>
			colNum(colNum:int):Ref<TableMetaBuilder>
			build():Ref<TableMeta>
	}
	export interface Text extends Struct<Text>{

			text:Ref<string>
			fontWeight:Ref<string>
			fontSize:Ref<int>
			horizontalAlign:Ref<string>
			verticalAlign:Ref<string>
	}
	export interface TextBuilder extends Struct<TextBuilder>{

			text(text:string):Ref<TextBuilder>
			fontWeight(fontWeight:string):Ref<TextBuilder>
			fontSize(fontSize:int):Ref<TextBuilder>
			horizontalAlign(horizontalAlign:string):Ref<TextBuilder>
			verticalAlign(verticalAlign:string):Ref<TextBuilder>
			build():Ref<Text>
	}
	export interface V1 extends Struct<V1>{

			whiteboardNode:Ref<{
			
				list(ctx:context.Context,req:Ref<ListWhiteboardNodeReq>,...options:larkcore.RequestOptionFunc[]):Ref<ListWhiteboardNodeResp>
			}>
	}
	export interface WhiteboardNode extends Struct<WhiteboardNode>{

			id:Ref<string>
			type:Ref<string>
			parentId:Ref<string>
			children:string[]
			X:Ref<float64>
			Y:Ref<float64>
			angle:Ref<float64>
			width:Ref<float64>
			height:Ref<float64>
			text:Ref<Text>
			style:Ref<Style>
			image:Ref<Image>
			compositeShape:Ref<CompositeShape>
			connector:Ref<Connector>
			section:Ref<Section>
			table:Ref<Table>
			mindMap:Ref<MindMap>
	}
	export interface WhiteboardNodeBuilder extends Struct<WhiteboardNodeBuilder>{

			id(id:string):Ref<WhiteboardNodeBuilder>
			type(type_:string):Ref<WhiteboardNodeBuilder>
			parentId(parentId:string):Ref<WhiteboardNodeBuilder>
			children(children:string[]):Ref<WhiteboardNodeBuilder>
			X(x:float64):Ref<WhiteboardNodeBuilder>
			Y(y:float64):Ref<WhiteboardNodeBuilder>
			angle(angle:float64):Ref<WhiteboardNodeBuilder>
			width(width:float64):Ref<WhiteboardNodeBuilder>
			height(height:float64):Ref<WhiteboardNodeBuilder>
			text(text:Ref<Text>):Ref<WhiteboardNodeBuilder>
			style(style:Ref<Style>):Ref<WhiteboardNodeBuilder>
			image(image:Ref<Image>):Ref<WhiteboardNodeBuilder>
			compositeShape(compositeShape:Ref<CompositeShape>):Ref<WhiteboardNodeBuilder>
			connector(connector:Ref<Connector>):Ref<WhiteboardNodeBuilder>
			section(section:Ref<Section>):Ref<WhiteboardNodeBuilder>
			table(table:Ref<Table>):Ref<WhiteboardNodeBuilder>
			mindMap(mindMap:Ref<MindMap>):Ref<WhiteboardNodeBuilder>
			build():Ref<WhiteboardNode>
	}
	export function emptyTable():Table
	export function emptyRefTable():Ref<Table>
	export function refOfTable(x:Table,v:Ref<Table>)
	export function unRefTable(v:Ref<Table>):Table
	export function emptyConnector():Connector
	export function emptyRefConnector():Ref<Connector>
	export function refOfConnector(x:Connector,v:Ref<Connector>)
	export function unRefConnector(v:Ref<Connector>):Connector
	export function emptyListWhiteboardNodeRespData():ListWhiteboardNodeRespData
	export function emptyRefListWhiteboardNodeRespData():Ref<ListWhiteboardNodeRespData>
	export function refOfListWhiteboardNodeRespData(x:ListWhiteboardNodeRespData,v:Ref<ListWhiteboardNodeRespData>)
	export function unRefListWhiteboardNodeRespData(v:Ref<ListWhiteboardNodeRespData>):ListWhiteboardNodeRespData
	export function emptySection():Section
	export function emptyRefSection():Ref<Section>
	export function refOfSection(x:Section,v:Ref<Section>)
	export function unRefSection(v:Ref<Section>):Section
	export function emptyTableCell():TableCell
	export function emptyRefTableCell():Ref<TableCell>
	export function refOfTableCell(x:TableCell,v:Ref<TableCell>)
	export function unRefTableCell(v:Ref<TableCell>):TableCell
	export function emptyCompositeShape():CompositeShape
	export function emptyRefCompositeShape():Ref<CompositeShape>
	export function refOfCompositeShape(x:CompositeShape,v:Ref<CompositeShape>)
	export function unRefCompositeShape(v:Ref<CompositeShape>):CompositeShape
	export function emptyListWhiteboardNodeReq():ListWhiteboardNodeReq
	export function emptyRefListWhiteboardNodeReq():Ref<ListWhiteboardNodeReq>
	export function refOfListWhiteboardNodeReq(x:ListWhiteboardNodeReq,v:Ref<ListWhiteboardNodeReq>)
	export function unRefListWhiteboardNodeReq(v:Ref<ListWhiteboardNodeReq>):ListWhiteboardNodeReq
	export function emptyDepartmentId():DepartmentId
	export function emptyRefDepartmentId():Ref<DepartmentId>
	export function refOfDepartmentId(x:DepartmentId,v:Ref<DepartmentId>)
	export function unRefDepartmentId(v:Ref<DepartmentId>):DepartmentId
	export function emptyTableMeta():TableMeta
	export function emptyRefTableMeta():Ref<TableMeta>
	export function refOfTableMeta(x:TableMeta,v:Ref<TableMeta>)
	export function unRefTableMeta(v:Ref<TableMeta>):TableMeta
	export function emptyStyle():Style
	export function emptyRefStyle():Ref<Style>
	export function refOfStyle(x:Style,v:Ref<Style>)
	export function unRefStyle(v:Ref<Style>):Style
	export function emptyV1():V1
	export function emptyRefV1():Ref<V1>
	export function refOfV1(x:V1,v:Ref<V1>)
	export function unRefV1(v:Ref<V1>):V1
	export function emptyWhiteboardNode():WhiteboardNode
	export function emptyRefWhiteboardNode():Ref<WhiteboardNode>
	export function refOfWhiteboardNode(x:WhiteboardNode,v:Ref<WhiteboardNode>)
	export function unRefWhiteboardNode(v:Ref<WhiteboardNode>):WhiteboardNode
	export function emptyConnectorCaption():ConnectorCaption
	export function emptyRefConnectorCaption():Ref<ConnectorCaption>
	export function refOfConnectorCaption(x:ConnectorCaption,v:Ref<ConnectorCaption>)
	export function unRefConnectorCaption(v:Ref<ConnectorCaption>):ConnectorCaption
	export function emptyImage():Image
	export function emptyRefImage():Ref<Image>
	export function refOfImage(x:Image,v:Ref<Image>)
	export function unRefImage(v:Ref<Image>):Image
	export function emptyMindMap():MindMap
	export function emptyRefMindMap():Ref<MindMap>
	export function refOfMindMap(x:MindMap,v:Ref<MindMap>)
	export function unRefMindMap(v:Ref<MindMap>):MindMap
	export function emptyTableCellMergeInfo():TableCellMergeInfo
	export function emptyRefTableCellMergeInfo():Ref<TableCellMergeInfo>
	export function refOfTableCellMergeInfo(x:TableCellMergeInfo,v:Ref<TableCellMergeInfo>)
	export function unRefTableCellMergeInfo(v:Ref<TableCellMergeInfo>):TableCellMergeInfo
	export function emptyConnectorAttachedObject():ConnectorAttachedObject
	export function emptyRefConnectorAttachedObject():Ref<ConnectorAttachedObject>
	export function refOfConnectorAttachedObject(x:ConnectorAttachedObject,v:Ref<ConnectorAttachedObject>)
	export function unRefConnectorAttachedObject(v:Ref<ConnectorAttachedObject>):ConnectorAttachedObject
	export function emptyListWhiteboardNodeResp():ListWhiteboardNodeResp
	export function emptyRefListWhiteboardNodeResp():Ref<ListWhiteboardNodeResp>
	export function refOfListWhiteboardNodeResp(x:ListWhiteboardNodeResp,v:Ref<ListWhiteboardNodeResp>)
	export function unRefListWhiteboardNodeResp(v:Ref<ListWhiteboardNodeResp>):ListWhiteboardNodeResp
	export function emptyText():Text
	export function emptyRefText():Ref<Text>
	export function refOfText(x:Text,v:Ref<Text>)
	export function unRefText(v:Ref<Text>):Text
}