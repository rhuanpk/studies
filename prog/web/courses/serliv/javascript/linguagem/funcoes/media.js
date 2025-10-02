(function () {
	function media() {
		let total = 0
		let qtd = arguments.length

		if (qtd === 0) {
			return 0
		}

		for (let index = 0; index < qtd; index++) {
			const arg = parseFloat(arguments[index]);

			// if (typeof arg !== 'number') {
			if (isNaN(arg)) {
				throw Error('expect only numbered arguments')
			}

			total += arg;
		}

		// return total / qtd
		return (total / qtd) || -1
	}

	console.log(media())
	console.log(media(2, 4))
	console.log(media(2, 4, 6))
	console.log(media('2', 4, 6))
	// console.log(media('a', 4, 6))
})()