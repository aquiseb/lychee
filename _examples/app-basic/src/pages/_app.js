import App from "next/app";
import { ChakraProvider } from "@chakra-ui/react";

import Cookies from "universal-cookie";
import consts from "consts";

import theme from "../theme";

function MyApp({ Component, pageProps }) {
	return (
		<ChakraProvider resetCSS theme={theme}>
			<Component {...pageProps} />
		</ChakraProvider>
	);
}

MyApp.getInitialProps = async (appContext) => {
	const appProps = await App.getInitialProps(appContext);

	const cookies = new Cookies(appContext?.ctx?.req?.headers?.cookie);
	const password = cookies.get(consts.SiteReadCookie) ?? "";

	if (password === "lychee") {
		appProps.pageProps.hasReadPermission = true;
	}

	return { ...appProps };
};

export default MyApp;
