import React, { Component } from "react";

import Posts from "../components/Posts";
import withIndex from "../lib/withIndex";
import loginMutation from "../lib/loginMutation";
import { Modal, ModalContent, CloseButton } from "../components/Modal";

class Index extends Component {
	static displayName = `Index`;

	constructor(props) {
		super(props);
		this.state = {};
	}

	componentDidMount() {
		console.log("Home did mount");
	}

	static async getInitialProps(context) {
		let { after, before, first, last } = context.query;
		if (last === undefined) {
			first = first || 2;
		}

		return {
			relayVariables: {
				after,
				before,
				first: first ? parseInt(first, 10) : first,
				last: last ? parseInt(last, 10) : last
			}
		};
	}

	login = () => {
		const { environment } = this.props.relay;
		loginMutation.commit({
			environment,
			input: { id: "u1" },
			onCompleted: () => alert("Successfully logged in. Check your cookies."),
			onError: () => alert("An error occured.")
		});
	};

	toggleModal = redirectUrl => this.setState({ isModalOpen: true, redirectUrl });

	render() {
		const { viewer, __fragments, __id, __fragmentOwner, relayVariables, ...nonQueryProps } = this.props;
		return (
			<div>
				<Modal open={this.state.isModalOpen}>
					<ModalContent>
						<CloseButton>&times;</CloseButton>
						<p>Some text in the Modal..</p>
						<iframe src={this.state.redirectUrl}></iframe>
					</ModalContent>
				</Modal>
				<h1>User</h1>
				<p>Is user logged in? {`${viewer.isLoggedIn}`}</p>
				{viewer.isLoggedIn ? <p>Welcome: {viewer.user.firstname}</p> : <button onClick={this.login}>Log in</button>}
				{!viewer.isLoggedIn ? <a href={this.props.oauth.facebookLink}>Facebook log in</a> : null}
				<Posts {...nonQueryProps} data={{ __fragments, __id, __fragmentOwner }} relayVariables={relayVariables} />
			</div>
		);
	}
}

// If you define the query here, nextjs throws an error:
// found page without a React Component as default export in pages/
export default withIndex(Index);
