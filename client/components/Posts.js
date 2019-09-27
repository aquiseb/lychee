import React from "react";
import Router from "next/router";
import _ from "lodash";

import PostPreview from "./PostPreview";
import withPosts from "../lib/withPosts";

class Posts extends React.Component {
	render() {
		const { props } = this;

		let afterParam = _.get(props, "data.allPosts.pageInfo.endCursor");
		afterParam = afterParam ? `&after=${afterParam}` : "";

		let hasNextPage = _.get(props, "data.allPosts.pageInfo.hasNextPage");
		hasNextPage = hasNextPage || props.relayVariables.before;

		let hasPrevPage = _.get(props, "data.allPosts.pageInfo.hasPreviousPage");
		hasPrevPage = hasPrevPage || props.relayVariables.after;

		let beforeParam = _.get(props, "data.allPosts.pageInfo.startCursor");
		beforeParam = beforeParam ? `&before=${beforeParam}` : "";

		const nextOnClick = () => Router.push(`/?first=2${afterParam}`);
		const prevOnClick = () => Router.push(`/?last=2${beforeParam}`);

		return (
			<div>
				<h1>Posts</h1>
				{props.data.allPosts.edges.map(({ node }) => (
					<PostPreview key={node.id} post={node} />
				))}
				<br />
				<button disabled={!hasPrevPage} onClick={prevOnClick}>
					Previous Page
				</button>
				&nbsp;
				<button disabled={!hasNextPage} onClick={nextOnClick}>
					Next Page
				</button>
			</div>
		);
	}
}

export default withPosts(Posts);
