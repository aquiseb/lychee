import { GraphQLObjectType, GraphQLString } from "graphql";
import { globalIdField, connectionArgs, connectionFromArray } from "graphql-relay";
import { PostConnection } from "./PostType";

const UserType = new GraphQLObjectType({
	name: "User",
	fields: () => ({
		id: globalIdField("User"),
		firstname: {
			type: GraphQLString
		},
		lastname: {
			type: GraphQLString
		},
		posts: {
			type: PostConnection.connectionType,
			args: connectionArgs,
			resolve: (obj, args, { db }, { rootValue: {} }) => {
				const postsOfUser = db.post.find({ creatorId: "u2" });
				return connectionFromArray(postsOfUser, args);
			}
		}
	})
});

export default UserType;
