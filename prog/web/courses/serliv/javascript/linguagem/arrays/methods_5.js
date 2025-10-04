let arr = [5, 4, 3, 2, 1]
console.log(arr)

console.log(arr.reverse())
console.log(arr)

// let result = arr.reduce(function (accumulator, atual, index, arr) {
// 	console.log('\nindex:', index)
// 	console.log('accumulator:', accumulator)
// 	console.log('atual:', atual)
// })
// console.log('\nresult:', result)

// let result = arr.reduce(function (accumulator, atual, index, arr) {
// 	console.log('\nindex:', index)
// 	console.log('accumulator:', accumulator)
// 	console.log('atual:', atual)
// 	return accumulator + atual
// })
// console.log('\nresult:', result)

// let result = arr.reduce(function (accumulator, atual, index, arr) {
// 	console.log('\nindex:', index)
// 	console.log('accumulator:', accumulator)
// 	console.log('atual:', atual)
// 	return accumulator + atual
// }, -1)
// console.log('\nresult:', result)

// let result = arr.reduce(function (accumulator, atual, index, arr) {
// 	console.log('\nindex:', index)
// 	console.log('accumulator:', accumulator)
// 	console.log('atual:', atual)
// 	return accumulator + atual
// }, 'a')
// console.log('\nresult:', result)


const names = ['Daniel', 'Maria', 'Joana', 'João']
let countedNames = names.reduce(function (namesObj, atualName) {
	let firstLetter = atualName[0]
	if (namesObj[firstLetter]) {
		namesObj[firstLetter]++
	} else {
		namesObj[firstLetter] = 1
	}
	return namesObj
}, {})
console.log()
console.log(countedNames)