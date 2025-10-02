/*
	muito abaixo do peso 16 a 16,9 kg/m2
	abaixo do peso 17 a 18,4 kg/m2
	normal 18,5 a 24,9 kg/m2
	acima do peso 25 a 29,9 kg/m2
	obesidade grau i 30 a 34,9 kg/m2
	obesidade grau ii 35 a 40 kg/m2
	obesidade grau iii maior que 40 kg/m2
*/

function calculaIMC(peso, altura, callback) {
	if (peso === undefined || altura === undefined) {
		throw Error('need two parameters: weight and height')
	}

	let imc = peso / (altura ** 2)

	if (typeof callback === 'function') {
		return callback(imc)
	}

	return imc
}

function classificaIMC(imc) {
	if (imc <= 16.9) {
		return 'muito a baixo do peso'
	} else if (imc <= 18.4) {
		return 'abaixo do peso'
	} else if (imc <= 24.9) {
		return 'normal'
	} else if (imc <= 29.9) {
		return 'acima do peso'
	} else if (imc <= 34.9) {
		return 'obesidade 1'
	} else if (imc <= 40) {
		return 'obesidade 2'
	} else {
		return 'obesidade 3'
	}
}

console.log(calculaIMC(60, 1.65))
console.log(calculaIMC(60, 1.65, classificaIMC))
