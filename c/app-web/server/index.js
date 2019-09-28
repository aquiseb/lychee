import express from "express";
import next from "next";

import getConfig, { setConfig } from "next/config";

const port = parseInt(process.env.PORT, 10) || 3000;
const dev = process.env.NODE_ENV !== "production";
const app = next({ dev });
const handle = app.getRequestHandler();

(async function() {
	await app.prepare();
	let { serverRuntimeConfig } = getConfig();

	const server = express();

	server.get("*", (req, res) => {
		return handle(req, res);
	});

	server.listen(port, err => {
		if (err) throw err;
		console.log(`> Ready on http://localhost:${port}`);
	});
})();
