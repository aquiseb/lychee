import express from "express";
import graphqlHTTP from "express-graphql";
import next from "next";
import passport from "passport";
import cookieParser from "cookie-parser";
import bodyParser from "body-parser";
import jwt from "jsonwebtoken";
import querystring from "querystring";

import schema from "./graphql/schema";
import db from "./db";

import { useFacebook, useJwt, useSerializer, getRedirectUrl, authorization } from "./auth";

import getConfig, { setConfig } from "next/config";

const port = parseInt(process.env.PORT, 10) || 3000;
const dev = process.env.NODE_ENV !== "production";
const app = next({ dev });
const handle = app.getRequestHandler();

(async function() {
	await app.prepare();
	let { serverRuntimeConfig } = getConfig();

	const server = express();

	server.use(bodyParser.json());
	server.use(bodyParser.urlencoded({ extended: false }));
	server.use(cookieParser());

	useSerializer();
	useJwt(serverRuntimeConfig);
	useFacebook(serverRuntimeConfig);

	server.use("/user", authorization, (req, res) => {
		res.send(req.user);
	});

	server.use(passport.initialize());

	server.get("/auth/facebook/callback", passport.authenticate("facebook", { session: true }), (req, res, next) => {
		const callbackState = querystring.parse(req.query.state);
		const user = req.user;

		req.login(user, { session: false }, err => {
			if (err) {
				res.send(err);
			}

			// generate a signed son web token with the contents of user object and return it in the response
			const lycheeToken = jwt.sign(user, serverRuntimeConfig.SERVER_SECRET);

			// Setting headers wouldn't work here because a redirect will execute a new http request,
			// use cookies or express-session to store the auth token and fetch it when you need it
			res.cookie("lycheeToken", lycheeToken, { httpOnly: false });
			req.lycheeToken = lycheeToken;
			res.redirect(callbackState.path);
		});
	});

	server.use(
		"/graphql",
		authorization,
		graphqlHTTP(({ user }, res, next) => {
			return {
				schema: schema,
				graphiql: true,
				pretty: true,
				context: {
					db,
					serverRuntimeConfig
				},
				rootValue: { user },
				credentials: "include"
			};
		})
	);

	server.get("*", (req, res) => {
		if (!req.path.startsWith("/_next")) {
			// Send the token via header for API calls
			if (req.lycheeToken) res.set("lycheeToken", req.lycheeToken);

			// Send the oauth redirection links to the client
			res.set(
				"oauth",
				JSON.stringify({
					facebookLink: getRedirectUrl(serverRuntimeConfig, req.path)
				})
			);
		}
		
		return handle(req, res);
	});

	server.listen(port, err => {
		if (err) throw err;
		console.log(`> Ready on http://localhost:${port}`);
	});
})();
