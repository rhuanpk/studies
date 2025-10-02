// function teste(cb) {
const teste = function (cb) {
	console.log('funcao teste')

	// if (typeof cb === 'function') {
	// 	cb('callback param')
	// }
	typeof cb === 'function' && cb('callback param')
}

// function fn() {
const fn = function (param) {
	console.log('funcao callback')
	console.log(param)
}

teste(function () {
	console.log('funcao callback')
})

teste(fn)
