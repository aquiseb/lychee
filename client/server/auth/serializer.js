import passport from "passport";

function serializer() {
	passport.serializeUser(function(user, done) {
		done(null, user.id);
	});

	passport.deserializeUser(function(id, done) {
		db.user.findOne({ id }, function(err, user) {
			done(err, user);
		});
	});
}

export default serializer;
