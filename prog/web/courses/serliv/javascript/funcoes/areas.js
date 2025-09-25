function calcularAreaRetangulo(l, a) {
	if (typeof l !== 'number' || typeof a !== 'number') {
		throw Error("calcularAreaRetangulo(): only accepts numbers")
	}
	return l * a
}

function calcularAreaTriangulo(b, a) {
	if (typeof b !== 'number' || typeof a !== 'number') {
		throw Error("calcularAreaTriangulo(): only accepts numbers")
	}
	return (b * a) / 2
}

function calcularAreaCirculo(r) {
	if (typeof r !== 'number') {
		throw Error("calcularAreaCirculo(): only accepts numbers")
	}
	return Math.PI * r ** 2
}