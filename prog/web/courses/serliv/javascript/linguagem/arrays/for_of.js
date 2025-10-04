const arr = [1, 2, 3]

const obj = {
	nome: 'Maria',
	idade: 28,
	email: 'maria@server.com'
}

// is for iterates over objects
for (const prop in obj) {
	console.log(prop, obj[prop])
}

// is for iterates over arrays
for (const el of arr) {
	console.log(el)
}