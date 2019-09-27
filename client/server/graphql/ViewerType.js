import { GraphQLObjectType, GraphQLBoolean } from "graphql";

import UserType from "./UserType";

const ViewerType = new GraphQLObjectType({
	name: "Viewer",
	fields: () => ({
		isLoggedIn: {
			type: GraphQLBoolean,
			resolve: (obj, args, { db, res }, { rootValue: { user } }) => {
				if (user) return true;
				return false;
			}
		},
		user: {
			type: UserType,
			resolve: (obj, args, { db }, { rootValue: { user = {} } }) => {
				return user;
			}
		}
	})
});

export default ViewerType;
