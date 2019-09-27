import { GraphQLSchema, GraphQLObjectType } from "graphql";
import { connectionArgs, connectionFromArray } from "graphql-relay";

import data from "../db";
import PostType, { PostConnection } from "./PostType";
import UserType from "./UserType";
import ViewerType from "./ViewerType";
import LoginType from "./LoginType";

const QueryType = new GraphQLObjectType({
	name: "Root",
	fields: () => ({
		viewer: {
			type: ViewerType,
			resolve: () => ({})
		},
		allPosts: {
			type: PostConnection.connectionType,
			args: connectionArgs,
			resolve: (obj, args, { db }, { rootValue: { user } }) => {
				const result = data.post.find();
				return connectionFromArray(result, args);
			}
		}
	})
});

const MutationType = new GraphQLObjectType({
	name: "Mutation",
	fields: {
		login: LoginType
	}
});

export default new GraphQLSchema({
	query: QueryType,
	mutation: MutationType
});
