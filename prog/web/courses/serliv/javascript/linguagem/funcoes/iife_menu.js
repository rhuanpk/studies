// (function () {
// 	let isValid = true
// 	console.log('menu', isValid)

// 	function init() {
// 		console.log('menu init')
// 	}
// 	init()
// })()

// (function (param) {
// 	console.log('menu', param)

// 	let isValid = true
// 	console.log('menu', isValid)

// 	function init() {
// 		console.log('menu init')
// 	}
// 	init()
// })('param')

(function (win, doc) {
	"use strict"
	win.alert('menu alert')

	let isValid = true
	console.log('menu', isValid)

	function init() {
		console.log('menu init')
	}
	init()
})(window, document)
