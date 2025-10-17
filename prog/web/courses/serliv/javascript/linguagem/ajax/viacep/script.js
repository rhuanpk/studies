; (function () {
	'use strict'

	const cep = document.getElementById('cep')
	const place = document.getElementById('place')
	const hood = document.getElementById('hood')
	const city = document.getElementById('city')
	const state = document.getElementById('state')
	const modal = document.querySelector('.error-modal')
	const msg = document.querySelector('.error-msg')
	const btn = document.querySelector('.error-close')

	function getAddress(cep) {
		// return fetch(`https://viacep.com.br/ws/${cep}/json`)
		return fetch(`https://viacep.com.br/ws/${cep}/json`, { mode: 'cors' })
			.then(response => {
				if (!response.ok) {
					throw new Error(`failure when request viacep "${cep}"`)
				}
				return response.json()
			})
	}

	function writeData(address) {
		place.value = address.logradouro || ''
		hood.value = address.bairro || ''
		city.value = address.localidade || ''
		state.value = address.uf || ''
	}

	function wipeData(address) {
		place.value = ''
		hood.value = ''
		city.value = ''
		state.value = ''
	}

	function showError(error) {
		msg.textContent = error
		modal.showModal()
	}

	btn.addEventListener('click', () => {
		modal.close()
	})

	cep.addEventListener('input', () => {
		let str = cep.value // if in a function it could "this.value"
		str = str.replace('-', '').trim()

		if (str.length === 8) {
			// const address = getAddress(cep)
			getAddress(str)
				.then(address => {
					if (address.error) { // "error" json property
						throw new Error("invalid cpf")
					}
					writeData(address)
				})
				.catch(error => {
					wipeData()
					showError(error)
				})
		}
	})
})()