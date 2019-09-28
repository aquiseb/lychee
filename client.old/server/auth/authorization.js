import passport from "passport";

const authMiddleware = (req, res, next) => {
	passport.authenticate("jwt", { session: false }, (err, user, info) => {
		if (user) req.user = user;
		// do some authorization stuff here
		next();
	})(req, res, next);
};

export default authMiddleware;
