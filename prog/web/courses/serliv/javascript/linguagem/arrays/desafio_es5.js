function sum() {
	const numbers = []

	// for (let index = 0; index < arguments.length; index++) {
	// 	numbers.push(arguments[index])
	// }

	Array.prototype.forEach.call(arguments, function (arg) {
		numbers.push(arg)
	})

	return numbers.reduce(function (sum, atual) {
		return sum + atual
	})
}

function avarage() {
	return sum.apply(null, arguments) / arguments.length
}

console.log(sum(1, 2, 3, 4, 5))
console.log(sum.call(null, 1, 2, 3, 4, 5))
console.log(sum.apply(null, [1, 2, 3, 4, 5]))

console.log(avarage(1, 2, 3, 4, 5))
console.log(avarage.call(null, 1, 2, 3, 4, 5))
console.log(avarage.apply(null, [1, 2, 3, 4, 5]))