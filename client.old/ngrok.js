const concurrently = require('concurrently');
const ngrok = require("ngrok");

(async function() {
	let url;
	try {
		console.log("sleep ---");
		url = await ngrok.connect({
			addr: 3000
		});
		console.log("URL ---", url);
		process.env.NGROK = url;
	} catch (err) {
		console.error(err);
		throw new Error("Could not create tunnel");
    }
    
    await concurrently([
        { command: 'nodemon --exec babel-node server', name: 'server' }
    ], {
        raw: true,
        prefix: 'name',
        killOthers: ['failure', 'success'],
        restartTries: 3,
    })
    
})();

module.exports = ngrok;
