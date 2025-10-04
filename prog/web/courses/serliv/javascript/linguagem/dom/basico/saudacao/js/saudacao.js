(function () {
	// const username = 'Daniel'
	const username = null
	const element = document.querySelector('.top-bar>p')

	if (username) {
		// element.textContent += username
		element.innerHTML += `<b>${username}</br>`
	} else {
		// const element2Remove = element.parentElement
		// element2Remove.parentElement.removeChild(element2Remove)

		// element.parentElement.style.display = 'none'

		element.parentElement.remove()
	}
})()