function sum() {
	// let total = 0
	// Array.from(arguments).forEach(function (value) {
	// 	total += value
	// })
	// return total
	return Array.from(arguments).reduce((sum, atual) => sum += atual)
}

function avarage() {
	return sum(...arguments) / arguments.length
}

console.log(sum(1, 2, 3))
console.log(avarage(1, 2, 3))