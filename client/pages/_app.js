import App from "next/app";
import React from "react";
import { ThemeProvider } from "styled-components";
import nextCookies from "next-cookies";
import Head from "next/head";
import { createGlobalStyle } from "styled-components";
import { Normalize } from "styled-normalize";

const theme = {
	colors: {
		primary: "#484848",
		secondary: "#767676",
		accent: "#0070f3",
		error: "#ff5a5f",
		fg: "#ffffff",
		bg: "#f6f6f6"
	}
};

const GlobalStyle = createGlobalStyle`
	h1, h2, h3, h4, h5, h6 {
		font-family: "Libre Baskerville", serif;
		color: ${props => props.theme.colors.primary}
	}

	#__next  {
		font-family: 'Lato', sans-serif;

		.has-error .ant-form-explain, .has-error .ant-form-split {
			color: ${props => props.theme.colors.error};
		}

		.has-error .ant-calendar-picker-icon::after, .has-error .ant-time-picker-icon::after, .has-error .ant-picker-icon::after, .has-error .ant-select-arrow, .has-error .ant-cascader-picker-arrow {
			color: ${props => props.theme.colors.error};
		}

		.has-error .ant-input, .has-error .ant-input:hover {
			border-color: ${props => props.theme.colors.error};
		}

		p, span, label {
			color: ${props => props.theme.colors.secondary};
		}

		button {
			span {
				color: unset;
				font-weight: bold;
				letter-spacing: .1em;
			}
		}

		.ant-form-item-required::before {
			display: none;
		}
	}
`;

export default class MyApp extends App {
	static async getInitialProps({ Component, ctx }, ...rest) {
		let pageProps = {};

		if (Component.getInitialProps) {
			pageProps = await Component.getInitialProps(ctx);
		}

		// const lycheeToken = ctx.res.get("lycheeToken"); // if sending from the headers
		const { lycheeToken } = nextCookies(ctx);
		let oauth = {};

		if (ctx.res) {
			oauth = ctx.res.get("oauth");
			if (oauth) oauth = JSON.parse(oauth);
		}

		return { pageProps, lycheeToken, oauth };
	}

	constructor(props) {
		super(props);
		this.state = {};
	}

	componentDidMount() {
		console.log("APP DID MOUNT");
		// Should replace with isLoggedIn
		// Replace the fragment that remains from facebook login
		if (this.props.lycheeToken && window.location.hash == "#_=_") history.replaceState(null, null, " ");
	}

	render() {
		const { Component, pageProps } = this.props;
		return (
			<>
				<Head>
					<link href="https://fonts.googleapis.com/css?family=Lato|Libre+Baskerville|Montserrat|Raleway&display=swap" rel="stylesheet"></link>
				</Head>
				<Normalize />
				<ThemeProvider theme={theme}>
					<GlobalStyle />
					<Component {...pageProps} oauth={this.props.oauth} />
				</ThemeProvider>
			</>
		);
	}
}
