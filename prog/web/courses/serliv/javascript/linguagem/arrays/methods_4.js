console.log('----- array -----')
let arr = [1, 3, 5, 7, 9]
console.log(arr)

console.log('----- push -----')
console.log(arr.push(11, 13, true, 'ola mundo'))
console.log(arr)

console.log('----- pop -----')
console.log(arr.pop())
console.log(arr)

console.log('----- shift -----')
console.log(arr.shift())
console.log(arr)

console.log('----- unshift -----')
console.log(arr.unshift('a', 'b', 'c'))
console.log(arr)

console.log('----- slice -----')
// from index start at index finish-1
// or from start (without finish) to end
console.log(arr.slice(2, 4))
console.log(arr)

console.log('----- splice -----')
// from index start more count items
// or from start (without finish) to end
// or and add on the place all arguments
console.log(arr.splice(2, 4, 'hello', 'world'))
console.log(arr)