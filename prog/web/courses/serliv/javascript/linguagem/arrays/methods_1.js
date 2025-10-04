let arr = [1, 5, 10, 'ola', true]

// returns true if all iterations are true
const onlyNumbers = arr.every(function (el) {
	console.log(el)
	return typeof el === 'number'
})
console.log(onlyNumbers)

// returns true if at least one iterations are true
const someNumber = arr.some(function (el) {
	console.log(el)
	return typeof el === 'number'
})
console.log(someNumber)

// returns a new array with the filtered elements
console.log(arr)
//arr = arr.filter(function (el, i, arr) {
const numbers = arr.filter(function (el, i, arr) {
	return typeof el === 'number'
})
console.log(numbers)

// iterates overall elements
arr.forEach(function (el, i, arr) {
	console.log(i, el)
})

// transform the same array
arr = arr.map(function (el, i, arr) {
	return el + el
})
console.log(arr)