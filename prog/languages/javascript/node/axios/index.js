const axios = require('axios');

class Fetch {
	async run() {
		try {
			const response = await axios.get('https://httpbin.org/get')
			return response.data;
		} catch (error) {
			console.log('axios error:', error);
		}
	}

	async call() {
		var response = await this.run();
		console.log(response);
	}
}

async function run() {
	try {
		const response = await axios.get('https://httpbin.org/get')
		return response.data;
	} catch (error) {
		console.log('axios error:', error);
	}
}

async function index() {
	const fetch = new Fetch();
	var response = await fetch.run();
	console.log(response);

	await fetch.call();

	response = await run();
	console.log(response);
}

index();
