function fn(cb) {
	console.log('function fn(cb)')
	console.log('typeof cb:', typeof cb)
	typeof cb === 'function' && cb()
}

// fn()

// fn(42)

// fn(function () {
// 	console.log('anonymous function')
// })

function callback() {
	console.log('anonymous function')
}
fn(callback)

const obj = {
	callback
}
obj.callback

function fn2(n1) {
	return function (n2) {
		return n1 * n2
	}
}
const fn3 = fn2(10)
const result = fn3(2)
console.log(result)

function fn4() {
	fn4.count++
	return function () {
		console.log('returned function')
	}
}
fn4.count = 0
const fn5 = fn4()
fn5()
fn4()
fn4()
console.log(fn4.count)
