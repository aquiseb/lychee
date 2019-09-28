import passport from "passport";
import FacebookStrategy from "passport-facebook";
import getConfig from "next/config";
import db from "../db";

export const getRedirectUrl = ({ auth: authConfig, app: appConfig }, pathWhereUserLeft) => {
	return [
		`https://www.facebook.com/v4.0/dialog/oauth`,
		`?response_type=code`,
		`&redirect_uri=${appConfig.URL}/auth/facebook/callback`,
		`&client_id=${authConfig.FACEBOOK_CLIENTID}`,
		`&state=path=${pathWhereUserLeft}`
	].join("");
};

export function useFacebook(serverRuntimeConfig) {
	const { auth: authConfig, app: appConfig } = serverRuntimeConfig;

	passport.use(
		new FacebookStrategy(
			{
				clientID: authConfig.FACEBOOK_CLIENTID,
				clientSecret: authConfig.FACEBOOK_CLIENTSECRET,
				callbackURL: appConfig.URL + "/auth/facebook/callback"
			},
			function(accessToken, refreshToken, profile, cb) {
				db.user.updateOne({ id: "u2" }, { $set: { facebookId: profile.id } }, (err, user) => {
					cb(null, user);
				});
			}
		)
	);
}
