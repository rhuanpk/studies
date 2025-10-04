let arr = [4, 5, 10, 20, 35, 4, 5]

// return the index of the finded value
console.log(arr.indexOf(5))
console.log(arr.indexOf(42))

// return the index of the last finded value
console.log(arr.lastIndexOf(5))
console.log(arr.lastIndexOf(42))

// return true if the finded value is in the array
console.log(arr.includes(5))
console.log(arr.includes(42))

// return the first element of condition
console.log(arr.find(function (el) {
	return el > 10
}))
console.log(arr.find(function (el) {
	return el > 42
}))

// return the index of first element of condition
console.log(arr.findIndex(function (el) {
	return el > 10
}))
console.log(arr.findIndex(function (el) {
	return el > 42
}))