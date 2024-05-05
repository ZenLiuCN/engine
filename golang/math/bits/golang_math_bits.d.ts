declare module "golang/math/bits" {
	// @ts-ignore
	import type {int,uint64,uint32,uint16,uint8,uint} from 'go'


	export function trailingZeros(x:uint):int
	export function len64(x:uint64):int
	export function div(hi,lo,y:uint):[uint,uint]
	export function leadingZeros(x:uint):int
	export function onesCount(x:uint):int
	export function onesCount64(x:uint64):int
	export function reverse(x:uint):uint
	export function reverseBytes64(x:uint64):uint64
	export function rem32(hi,lo,y:uint32):uint32
	export function rem64(hi,lo,y:uint64):uint64
	export function leadingZeros64(x:uint64):int
	export function onesCount32(x:uint32):int
	export const UintSize=0
	export function len16(x:uint16):int
	export function mul64(x,y:uint64):[uint64,uint64]
	export function trailingZeros32(x:uint32):int
	export function rotateLeft64(x:uint64,k:int):uint64
	export function rem(hi,lo,y:uint):uint
	export function trailingZeros8(x:uint8):int
	export function rotateLeft8(x:uint8,k:int):uint8
	export function sub64(x,y,borrow:uint64):[uint64,uint64]
	export function reverse64(x:uint64):uint64
	export function reverseBytes16(x:uint16):uint16
	export function sub32(x,y,borrow:uint32):[uint32,uint32]
	export function leadingZeros8(x:uint8):int
	export function trailingZeros64(x:uint64):int
	export function reverse8(x:uint8):uint8
	export function reverseBytes32(x:uint32):uint32
	export function sub(x,y,borrow:uint):[uint,uint]
	export function leadingZeros16(x:uint16):int
	export function trailingZeros16(x:uint16):int
	export function rotateLeft16(x:uint16,k:int):uint16
	export function reverse16(x:uint16):uint16
	export function reverse32(x:uint32):uint32
	export function len32(x:uint32):int
	export function div64(hi,lo,y:uint64):[uint64,uint64]
	export function add(x,y,carry:uint):[uint,uint]
	export function onesCount16(x:uint16):int
	export function len(x:uint):int
	export function add32(x,y,carry:uint32):[uint32,uint32]
	export function mul32(x,y:uint32):[uint32,uint32]
	export function onesCount8(x:uint8):int
	export function rotateLeft(x:uint,k:int):uint
	export function add64(x,y,carry:uint64):[uint64,uint64]
	export function div32(hi,lo,y:uint32):[uint32,uint32]
	export function leadingZeros32(x:uint32):int
	export function reverseBytes(x:uint):uint
	export function len8(x:uint8):int
	export function mul(x,y:uint):[uint,uint]
	export function rotateLeft32(x:uint32,k:int):uint32

}