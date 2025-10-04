const arr1 = [1, 2, 3]
const arr2 = [4, 5, 6]

function fn() {
	console.log(arguments.length, arguments)
}

fn(1, 2, 3)
fn(arr1)
fn([1, 2, 3])
fn(...arr1)
fn(...[1, 2, 3])

arr1.push(...arr2)
console.log(arr1)