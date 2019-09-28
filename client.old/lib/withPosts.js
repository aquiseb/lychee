import { createFragmentContainer, graphql } from "react-relay";

export default ComposedComponent => createFragmentContainer(ComposedComponent, {
	data: graphql`
		fragment withPosts_data on Root
			@argumentDefinitions(after: { type: "String" }, before: { type: "String" }, first: { type: "Int", defaultValue: 2 }, last: { type: "Int" }) {
			allPosts(after: $after, before: $before, first: $first, last: $last) {
				pageInfo {
					hasNextPage
					hasPreviousPage
					startCursor
					endCursor
				}
				edges {
					node {
						...withPostPreview_post
						title
						id
					}
				}
			}
		}
	`
});