(function () {
	const username = 'Daniel'
	// const username = null

	if (username) {
		const topBarElement = document.createElement('div')
		topBarElement.className = 'top-bar'
		topBarElement.innerHTML = `<p>Bem-vindo(a), <b>${username}</b></p>`

		const parentElement = document.querySelector('.hero')
		parentElement.insertBefore(topBarElement, parentElement.firstElementChild)
	}
})()