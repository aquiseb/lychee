import { createFragmentContainer, graphql } from "react-relay";

export default ComposedComponent => createFragmentContainer(ComposedComponent, {
	post: graphql`
		fragment withPostPreview_post on Post {
			id
			title
			content
		}
	`
});
