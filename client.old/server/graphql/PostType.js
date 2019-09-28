import { GraphQLObjectType, GraphQLString } from "graphql";
import { globalIdField, connectionDefinitions } from "graphql-relay";

const PostType = new GraphQLObjectType({
	name: "Post",
	description: "A post",
	fields: () => ({
		id: globalIdField("PostType"),
		content: {
			type: GraphQLString
		},
		title: {
			type: GraphQLString
		}
	})
});

export const PostConnection = connectionDefinitions({
	name: "Post",
	nodeType: PostType
});

export default PostConnection;
