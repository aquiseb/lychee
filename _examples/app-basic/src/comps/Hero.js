import { Flex, Heading } from "@chakra-ui/react";

export const Hero = ({ title }) => (
	<Flex position="absolute" justifyContent="center" alignItems="center" height="100vh" bgGradient="linear(to-l, heroGradientStart, heroGradientEnd)" bgClip="text" pointerEvents="none" h="full">
		<Heading fontSize="6vw" pointerEvents="none">
			{title}
		</Heading>
	</Flex>
);

Hero.defaultProps = {
	title: "lychee",
};
