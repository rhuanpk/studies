// array destructuring
const arr = [1, 2, 3]

//const n1 = arr[0]
//const n2 = arr[1]
const [n1, n2] = arr

// array destructuring
const obj = { name: 'Daniel', surname: 'Morales', company: 'Serliv' }
const person = { name: 'João', surname: 'Duarte', age: 'Serliv' }
const family = { name: 'Souza', surname: 'Ferrer', count: '2' }

//const name = obj.name
//const company = obj.company
const { name, company } = obj
console.log(name, company)

//const { name: nomeAlternativo, company: daPropriedade } = obj
//console.log(nomeAlternativo, daPropriedade)

// const { name, company, email = null } = obj
// console.log(name, company, email)

// const { name, ...rest } = person
// console.log(name, rest)

// function xpto({ name, surname }) {
// 	console.log(`Hello family ${name} ${surname}!`)
// }
// xpto(family)
// xpto({ name: 'Contrast', count: '2' })