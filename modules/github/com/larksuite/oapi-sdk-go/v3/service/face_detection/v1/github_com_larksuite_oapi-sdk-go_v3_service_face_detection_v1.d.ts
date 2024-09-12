// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/face_detection/v1'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {int,Ref,float64,Struct,error,bool} from 'go'
	export interface AttributeItem extends Struct<AttributeItem>{

			type:Ref<int>
			probability:Ref<float64>
	}
	export interface AttributeItemBuilder extends Struct<AttributeItemBuilder>{

			type(type_:int):Ref<AttributeItemBuilder>
			probability(probability:float64):Ref<AttributeItemBuilder>
			build():Ref<AttributeItem>
	}
	export interface DetectFaceAttributesImagePathReqBodyBuilder extends Struct<DetectFaceAttributesImagePathReqBodyBuilder>{

			image(image:string):Ref<DetectFaceAttributesImagePathReqBodyBuilder>
			build():Ref<DetectFaceAttributesImageReqBody>
	}
	export interface DetectFaceAttributesImageReq extends Struct<DetectFaceAttributesImageReq>{

			body:Ref<DetectFaceAttributesImageReqBody>
	}
	export interface DetectFaceAttributesImageReqBody extends Struct<DetectFaceAttributesImageReqBody>{

			image:Ref<string>
	}
	export interface DetectFaceAttributesImageReqBodyBuilder extends Struct<DetectFaceAttributesImageReqBodyBuilder>{

			image(image:string):Ref<DetectFaceAttributesImageReqBodyBuilder>
			build():Ref<DetectFaceAttributesImageReqBody>
	}
	export interface DetectFaceAttributesImageReqBuilder extends Struct<DetectFaceAttributesImageReqBuilder>{

			body(body:Ref<DetectFaceAttributesImageReqBody>):Ref<DetectFaceAttributesImageReqBuilder>
			build():Ref<DetectFaceAttributesImageReq>
	}
	export interface DetectFaceAttributesImageResp extends Struct<DetectFaceAttributesImageResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<DetectFaceAttributesImageRespData>
			success():bool
	}
	export interface DetectFaceAttributesImageRespData extends Struct<DetectFaceAttributesImageRespData>{

			imageInfo:Ref<Image>
			faceInfos:Ref<FaceInfo>[]
	}
	export interface FaceAttribute extends Struct<FaceAttribute>{

			gender:Ref<AttributeItem>
			age:Ref<int>
			emotion:Ref<AttributeItem>
			beauty:Ref<int>
			pose:Ref<FacePose>
			hat:Ref<AttributeItem>
			glass:Ref<AttributeItem>
			mask:Ref<AttributeItem>
	}
	export interface FaceAttributeBuilder extends Struct<FaceAttributeBuilder>{

			gender(gender:Ref<AttributeItem>):Ref<FaceAttributeBuilder>
			age(age:int):Ref<FaceAttributeBuilder>
			emotion(emotion:Ref<AttributeItem>):Ref<FaceAttributeBuilder>
			beauty(beauty:int):Ref<FaceAttributeBuilder>
			pose(pose:Ref<FacePose>):Ref<FaceAttributeBuilder>
			hat(hat:Ref<AttributeItem>):Ref<FaceAttributeBuilder>
			glass(glass:Ref<AttributeItem>):Ref<FaceAttributeBuilder>
			mask(mask:Ref<AttributeItem>):Ref<FaceAttributeBuilder>
			build():Ref<FaceAttribute>
	}
	export interface FaceDetectionService extends Struct<FaceDetectionService>{

			image:Ref<{
			
				detectFaceAttributes(ctx:context.Context,req:Ref<DetectFaceAttributesImageReq>,...options:larkcore.RequestOptionFunc[]):Ref<DetectFaceAttributesImageResp>
			}>
	}
	export interface FaceInfo extends Struct<FaceInfo>{

			position:Ref<FacePosition>
			attribute:Ref<FaceAttribute>
			quality:Ref<FaceQuality>
	}
	export interface FaceInfoBuilder extends Struct<FaceInfoBuilder>{

			position(position:Ref<FacePosition>):Ref<FaceInfoBuilder>
			attribute(attribute:Ref<FaceAttribute>):Ref<FaceInfoBuilder>
			quality(quality:Ref<FaceQuality>):Ref<FaceInfoBuilder>
			build():Ref<FaceInfo>
	}
	export interface FaceOcclude extends Struct<FaceOcclude>{

			eyebrow:Ref<float64>
			nose:Ref<float64>
			cheek:Ref<float64>
			mouth:Ref<float64>
			chin:Ref<float64>
			leftEye:Ref<float64>
			rightEye:Ref<float64>
	}
	export interface FaceOccludeBuilder extends Struct<FaceOccludeBuilder>{

			eyebrow(eyebrow:float64):Ref<FaceOccludeBuilder>
			nose(nose:float64):Ref<FaceOccludeBuilder>
			cheek(cheek:float64):Ref<FaceOccludeBuilder>
			mouth(mouth:float64):Ref<FaceOccludeBuilder>
			chin(chin:float64):Ref<FaceOccludeBuilder>
			leftEye(leftEye:float64):Ref<FaceOccludeBuilder>
			rightEye(rightEye:float64):Ref<FaceOccludeBuilder>
			build():Ref<FaceOcclude>
	}
	export interface FacePose extends Struct<FacePose>{

			pitch:Ref<int>
			yaw:Ref<int>
			roll:Ref<int>
	}
	export interface FacePoseBuilder extends Struct<FacePoseBuilder>{

			pitch(pitch:int):Ref<FacePoseBuilder>
			yaw(yaw:int):Ref<FacePoseBuilder>
			roll(roll:int):Ref<FacePoseBuilder>
			build():Ref<FacePose>
	}
	export interface FacePosition extends Struct<FacePosition>{

			upperLeft:Ref<Point>
			lowerRight:Ref<Point>
	}
	export interface FacePositionBuilder extends Struct<FacePositionBuilder>{

			upperLeft(upperLeft:Ref<Point>):Ref<FacePositionBuilder>
			lowerRight(lowerRight:Ref<Point>):Ref<FacePositionBuilder>
			build():Ref<FacePosition>
	}
	export interface FaceQuality extends Struct<FaceQuality>{

			sharpness:Ref<float64>
			brightness:Ref<float64>
			occlude:Ref<FaceOcclude>
	}
	export interface FaceQualityBuilder extends Struct<FaceQualityBuilder>{

			sharpness(sharpness:float64):Ref<FaceQualityBuilder>
			brightness(brightness:float64):Ref<FaceQualityBuilder>
			occlude(occlude:Ref<FaceOcclude>):Ref<FaceQualityBuilder>
			build():Ref<FaceQuality>
	}
	export interface Image extends Struct<Image>{

			width:Ref<int>
			height:Ref<int>
	}
	export interface ImageBuilder extends Struct<ImageBuilder>{

			width(width:int):Ref<ImageBuilder>
			height(height:int):Ref<ImageBuilder>
			build():Ref<Image>
	}
	export function newAttributeItemBuilder():Ref<AttributeItemBuilder>

	export function newDetectFaceAttributesImagePathReqBodyBuilder():Ref<DetectFaceAttributesImagePathReqBodyBuilder>

	export function newDetectFaceAttributesImageReqBodyBuilder():Ref<DetectFaceAttributesImageReqBodyBuilder>

	export function newDetectFaceAttributesImageReqBuilder():Ref<DetectFaceAttributesImageReqBuilder>

	export function newFaceAttributeBuilder():Ref<FaceAttributeBuilder>

	export function newFaceInfoBuilder():Ref<FaceInfoBuilder>

	export function newFaceOccludeBuilder():Ref<FaceOccludeBuilder>

	export function newFacePoseBuilder():Ref<FacePoseBuilder>

	export function newFacePositionBuilder():Ref<FacePositionBuilder>

	export function newFaceQualityBuilder():Ref<FaceQualityBuilder>

	export function newImageBuilder():Ref<ImageBuilder>

	export function newPointBuilder():Ref<PointBuilder>

	export function newService(config:Ref<larkcore.Config>):Ref<FaceDetectionService>

	export interface Point extends Struct<Point>{

			X:Ref<float64>
			Y:Ref<float64>
	}
	export interface PointBuilder extends Struct<PointBuilder>{

			X(x:float64):Ref<PointBuilder>
			Y(y:float64):Ref<PointBuilder>
			build():Ref<Point>
	}
	export function emptyDetectFaceAttributesImageRespData():DetectFaceAttributesImageRespData
	export function emptyRefDetectFaceAttributesImageRespData():Ref<DetectFaceAttributesImageRespData>
	export function refOfDetectFaceAttributesImageRespData(x:DetectFaceAttributesImageRespData,v:Ref<DetectFaceAttributesImageRespData>)
	export function unRefDetectFaceAttributesImageRespData(v:Ref<DetectFaceAttributesImageRespData>):DetectFaceAttributesImageRespData
	export function emptyImage():Image
	export function emptyRefImage():Ref<Image>
	export function refOfImage(x:Image,v:Ref<Image>)
	export function unRefImage(v:Ref<Image>):Image
	export function emptyFaceOcclude():FaceOcclude
	export function emptyRefFaceOcclude():Ref<FaceOcclude>
	export function refOfFaceOcclude(x:FaceOcclude,v:Ref<FaceOcclude>)
	export function unRefFaceOcclude(v:Ref<FaceOcclude>):FaceOcclude
	export function emptyDetectFaceAttributesImageReq():DetectFaceAttributesImageReq
	export function emptyRefDetectFaceAttributesImageReq():Ref<DetectFaceAttributesImageReq>
	export function refOfDetectFaceAttributesImageReq(x:DetectFaceAttributesImageReq,v:Ref<DetectFaceAttributesImageReq>)
	export function unRefDetectFaceAttributesImageReq(v:Ref<DetectFaceAttributesImageReq>):DetectFaceAttributesImageReq
	export function emptyFaceInfo():FaceInfo
	export function emptyRefFaceInfo():Ref<FaceInfo>
	export function refOfFaceInfo(x:FaceInfo,v:Ref<FaceInfo>)
	export function unRefFaceInfo(v:Ref<FaceInfo>):FaceInfo
	export function emptyAttributeItem():AttributeItem
	export function emptyRefAttributeItem():Ref<AttributeItem>
	export function refOfAttributeItem(x:AttributeItem,v:Ref<AttributeItem>)
	export function unRefAttributeItem(v:Ref<AttributeItem>):AttributeItem
	export function emptyFaceDetectionService():FaceDetectionService
	export function emptyRefFaceDetectionService():Ref<FaceDetectionService>
	export function refOfFaceDetectionService(x:FaceDetectionService,v:Ref<FaceDetectionService>)
	export function unRefFaceDetectionService(v:Ref<FaceDetectionService>):FaceDetectionService
	export function emptyFacePosition():FacePosition
	export function emptyRefFacePosition():Ref<FacePosition>
	export function refOfFacePosition(x:FacePosition,v:Ref<FacePosition>)
	export function unRefFacePosition(v:Ref<FacePosition>):FacePosition
	export function emptyDetectFaceAttributesImageResp():DetectFaceAttributesImageResp
	export function emptyRefDetectFaceAttributesImageResp():Ref<DetectFaceAttributesImageResp>
	export function refOfDetectFaceAttributesImageResp(x:DetectFaceAttributesImageResp,v:Ref<DetectFaceAttributesImageResp>)
	export function unRefDetectFaceAttributesImageResp(v:Ref<DetectFaceAttributesImageResp>):DetectFaceAttributesImageResp
	export function emptyFacePose():FacePose
	export function emptyRefFacePose():Ref<FacePose>
	export function refOfFacePose(x:FacePose,v:Ref<FacePose>)
	export function unRefFacePose(v:Ref<FacePose>):FacePose
	export function emptyPoint():Point
	export function emptyRefPoint():Ref<Point>
	export function refOfPoint(x:Point,v:Ref<Point>)
	export function unRefPoint(v:Ref<Point>):Point
	export function emptyDetectFaceAttributesImageReqBody():DetectFaceAttributesImageReqBody
	export function emptyRefDetectFaceAttributesImageReqBody():Ref<DetectFaceAttributesImageReqBody>
	export function refOfDetectFaceAttributesImageReqBody(x:DetectFaceAttributesImageReqBody,v:Ref<DetectFaceAttributesImageReqBody>)
	export function unRefDetectFaceAttributesImageReqBody(v:Ref<DetectFaceAttributesImageReqBody>):DetectFaceAttributesImageReqBody
	export function emptyFaceAttribute():FaceAttribute
	export function emptyRefFaceAttribute():Ref<FaceAttribute>
	export function refOfFaceAttribute(x:FaceAttribute,v:Ref<FaceAttribute>)
	export function unRefFaceAttribute(v:Ref<FaceAttribute>):FaceAttribute
	export function emptyFaceQuality():FaceQuality
	export function emptyRefFaceQuality():Ref<FaceQuality>
	export function refOfFaceQuality(x:FaceQuality,v:Ref<FaceQuality>)
	export function unRefFaceQuality(v:Ref<FaceQuality>):FaceQuality
}