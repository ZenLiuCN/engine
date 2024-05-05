declare module "golang/math/cmplx" {
	// @ts-ignore
	import type {complex128,bool,float64} from 'go'


	export function tanh(x:complex128):complex128
	export function atanh(x:complex128):complex128
	export function isNaN(x:complex128):bool
	export function log(x:complex128):complex128
	export function sinh(x:complex128):complex128
	export function sqrt(x:complex128):complex128
	export function pow(x,y:complex128):complex128
	export function cos(x:complex128):complex128
	export function abs(x:complex128):float64
	export function asinh(x:complex128):complex128
	export function log10(x:complex128):complex128
	export function polar(x:complex128):[float64,float64]
	export function tan(x:complex128):complex128
	export function rect(r,θ:float64):complex128
	export function cot(x:complex128):complex128
	export function isInf(x:complex128):bool
	export function sin(x:complex128):complex128
	export function asin(x:complex128):complex128
	export function exp(x:complex128):complex128
	export function phase(x:complex128):float64
	export function acos(x:complex128):complex128
	export function acosh(x:complex128):complex128
	export function conj(x:complex128):complex128
	export function inf():complex128
	export function atan(x:complex128):complex128
	export function naN():complex128
	export function cosh(x:complex128):complex128

}