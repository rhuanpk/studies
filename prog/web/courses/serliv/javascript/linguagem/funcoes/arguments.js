function sum() {
// const sum = function () {
// const sum = () => {
	console.log(arguments)
	let total = 0
	for (let index = 0; index < arguments.length; index++) {
		total += arguments[index]
	}
	return total
}

console.log(sum(1, 2, 3))
console.log(sum(1, 2, 3, 4, 5))
console.log(sum(1, 2, 3, 4, 5, 12, 14, 50))

console.log(sum.name)
