import { useColorMode, IconButton, useTheme } from "@chakra-ui/react";
import { SunIcon, MoonIcon } from "@chakra-ui/icons";

export const DarkModeSwitch = () => {
	const theme = useTheme();
	const { colorMode, toggleColorMode } = useColorMode(theme.config.initialColorMode);

	const isDark = colorMode === "dark";
	return <IconButton position="fixed" top={4} right={16} icon={isDark ? <SunIcon /> : <MoonIcon />} aria-label="Toggle Theme" colorScheme="blue" onClick={toggleColorMode} />;
};
