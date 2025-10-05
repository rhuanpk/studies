; (function () {
	'use strict'

	const txtTitulo = document.getElementById('txtTitulo')
	const btn = document.getElementById('btn')
	const formCadastro = document.querySelector('.formCadastro')
	const feedbackMessage = document.getElementById('feedbackMessage')
	const feedbackMessageCloseBtn = feedbackMessage.getElementsByTagName('button')[0]
	function showErrorMessage(msg, cb) {
		function hideErrorMessage() {
			feedbackMessageCloseBtn.removeEventListener('click', hideErrorMessage)
			feedbackMessageCloseBtn.removeEventListener('keyup', pressedKeyboardOnBtn)
			feedbackMessage.classList.remove('show')
			if (typeof cb === 'function') {
				cb()
			}
		}
		function pressedKeyboardOnBtn(e) {
			if (e.keyCode === 27) {
				hideErrorMessage()
			}
		}
		feedbackMessageCloseBtn.addEventListener('click', hideErrorMessage)
		feedbackMessageCloseBtn.addEventListener('keyup', pressedKeyboardOnBtn)
		feedbackMessage.getElementsByTagName('p')[0].textContent = msg
		feedbackMessage.classList.add('show')
		feedbackMessageCloseBtn.focus()
	}

	// btn.addEventListener('click', e => {
	// 	if (!txtTitulo.value) { // if title input is empty
	// 		alert('Preencha todos os campos!')
	// 		e.preventDefault() // "e" is the own button
	// 		txtTitulo.focus()
	// 	}
	// })
	formCadastro.addEventListener('submit', e => {
		if (!txtTitulo.value) { // if title input is empty
			showErrorMessage('Preencha todos os campos!', function () {
				txtTitulo.focus()
			})
			e.preventDefault() // "e" is the own button
		}
	})

	const txtDescricao = document.getElementById('txtDescricao')
	const contadorContainer = document.getElementById('contador')
	const resta = contadorContainer.getElementsByTagName('span')[0]
	const maxima = txtDescricao.maxLength
	contadorContainer.style.display = 'unset'
	function setNumero(numero) {
		resta.textContent = numero
	}
	function checkLength() {
		// let numeroLetrasDigitadas = txtDescricao.value.length
		const numeroLetrasDigitadas = this.value.length
		// const caracteresRestantes = maxima - numeroLetrasDigitadas
		const caracteresRestantes = parseInt(maxima) - parseInt(numeroLetrasDigitadas)
		setNumero(caracteresRestantes)
	}
	setNumero(maxima)
	// txtDescricao.addEventListener('keyup', () => { console.log('keyup') }) // triggered when the key is released
	// txtDescricao.addEventListener('keydown', () => { console.log('keydown') }) // triggered when the key is pressed
	// txtDescricao.addEventListener('keypress', () => { console.log('keypress') }) // triggered when a "character" (letter, number, ponctuation or symbol) is typed (deprecated)
	// txtDescricao.addEventListener('input', () => { console.log('input') }) // triggered on input changes
	txtDescricao.addEventListener('input', checkLength)

	btn.disabled = true
	const chkAceito = document.getElementById('chkAceito')
	chkAceito.addEventListener('change', function () {
		btn.disabled = !this.checked
	})
})()