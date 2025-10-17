import somarImport from "./soma.js" // changing the name fo the default export of this module

// function media(...args) {
// 	const total = somar(...args)
// 	return total / args.length
// }

export default function media(...args) {
	const total = somarImport(...args)
	return total / args.length
}

// function somar(...args) {
// 	console.log('funcao interna somar nao exportada e que não sobreescreve')
// }
export function somar(...args) {
	console.log('funcao interna somar nao exportada e que não sobreescreve')
}

// export function validar(...args) {
// 	console.log('validando media')
// }
// export const XPTO = 'xpto'

function validar(...args) {
	console.log('validando media')
}
const XPTO = 'xpto'
export { validar, XPTO }