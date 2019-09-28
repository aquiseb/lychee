import { GraphQLNonNull, GraphQLString } from "graphql";
import { mutationWithClientMutationId } from "graphql-relay";

import UserType from "./UserType";

const LoginType = mutationWithClientMutationId({
	name: "Login",
	inputFields: {
		id: {
			type: new GraphQLNonNull(GraphQLString)
		}
	},
	outputFields: {
		user: {
			type: UserType,
			resolve: () => ({})
		}
	},
	mutateAndGetPayload: (args, { db, res }, { rootValue: { user } }) => {
		console.log("Setting the cookie!", rootValue);
		const userId = args.id;
		res.setHeader("Set-Cookie", [`nextjs_authentication_token=crypted_${userId}_token`]);
		return {};
	}
});

export default LoginType;
