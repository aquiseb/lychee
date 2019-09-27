import React from "react";

import withPostPreview from "../lib/withPostPreview"

const PostPreview = props => {
	return (
		<React.Fragment>
			<div>{props.post.title}</div>
			<div>{props.post.content}</div>
			<br />
		</React.Fragment>
	);
};

export default withPostPreview(PostPreview)