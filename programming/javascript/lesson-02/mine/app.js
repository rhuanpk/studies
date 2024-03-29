alert('hello world');
let secretNumber = 1;
console.log(secretNumber)
let input = prompt('enter the secret number:')
if (input == secretNumber) {
	alert(`you GET secret number (${secretNumber}) :D`)
} else {
	alert('you MISSED secret number :(')
}
