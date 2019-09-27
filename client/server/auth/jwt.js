import passport from "passport";
import passportJWT from "passport-jwt";
import db from "../db";

const JWTStrategy = passportJWT.Strategy;
// const ExtractJWT = passportJWT.ExtractJwt;

var cookieExtractor = function(req) {
	var token = null;
	if (req && req.cookies) token = req.cookies["lycheeToken"];
	return token;
};

export function useJwt(serverRuntimeConfig) {
	passport.use(
		new JWTStrategy(
			{
				jwtFromRequest: cookieExtractor,
				// jwtFromRequest: ExtractJWT.fromAuthHeaderAsBearerToken(),
				secretOrKey: serverRuntimeConfig.SERVER_SECRET
			},
			function(jwtPayload, cb) {
				//find the user in db if needed. This functionality may be omitted if you store everything you'll need in JWT payload.
				return db.user.findOne({ id: jwtPayload.id }, (err, user) => {
					cb(null, user);
				});
			}
		)
	);
}
