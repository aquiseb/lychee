import { useRouter } from "next/router";

import { Flex, Box } from "@chakra-ui/react";

import { Hero } from "../comps/Hero";
import { Container } from "../comps/Container";
import { DarkModeSwitch } from "../comps/DarkModeSwitch";
import LogoutButton from "../comps/LogoutButton";
import Login from "../comps/Login";

const Index = (props) => {
	const router = useRouter();

	const { data } = props;
	const { post = {} } = data ?? {};

	if (!props.hasReadPermission) {
		return <Login redirectPath={router.asPath} />;
	}

	return (
		<Container id="lychee-container" height="100vh">
			<Hero />
			<Flex h="100vh" alignItems="center">
				<Box mt={40}>
					Fetched post id {post.id} - {post.title}
				</Box>
			</Flex>
			<DarkModeSwitch />
			<LogoutButton />
		</Container>
	);
};

const query = `
{
    post(id: "1") {
        id
		title
    }
}
`;

const variables = {
	id: 1,
};

export async function getServerSideProps(ctx) {
	let data;
	try {
		data = await fetch("http://localhost:4002/graphql", {
			method: "post",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({ query, variables }),
		});

		({ data } = await data.json());
	} catch (err) {
		console.log("Fetch error ::", err);
	}

	return { props: { data } };
}

export default Index;
