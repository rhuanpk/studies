// function somar(...args) {
// 	if (args.length === 0) {
// 		throw new Error("A função requer pelo menos um parâmetro.")
// 	}
// 	if (args.some(arg => typeof arg !== 'number' || isNaN(arg))) {
// 		throw new Error("Parâmetros inválidos: Todos os parâmetros devem ser números.")
// 	}
// 	return args.reduce((total, atual) => total + atual, 0)
// }

export default function somar(...args) { // the default keyword says that this entity is the default when this module is imported
	if (args.length === 0) {
		throw new Error("A função requer pelo menos um parâmetro.")
	}
	if (args.some(arg => typeof arg !== 'number' || isNaN(arg))) {
		throw new Error("Parâmetros inválidos: Todos os parâmetros devem ser números.")
	}
	return args.reduce((total, atual) => total + atual, 0)
}