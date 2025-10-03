const txtEmail = document.getElementById('txt-email')
const msgFeedback = document.getElementById('newsletter-feedback')

function cadastrarEmail() {
	let email = txtEmail.value
	msgFeedback.innerHTML = `O email ${email} foi cadastrado com sucesso.`
}
