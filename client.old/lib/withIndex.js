import { graphql } from "react-relay";

import withData from "./withData";

export default ComposedComponent =>
	withData(ComposedComponent, {
		query: graphql`
			query withIndex_Query($after: String, $before: String, $first: Int, $last: Int) {
				...withPosts_data @arguments(after: $after, before: $before, first: $first, last: $last)
				viewer {
					isLoggedIn
					user {
						firstname
						lastname
					}
				}
			}
		`
	});
