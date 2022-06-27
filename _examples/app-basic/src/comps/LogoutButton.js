import Cookies from "universal-cookie";
import consts from "consts";

import { IconButton, useTheme } from "@chakra-ui/react";
import { FiLogOut } from "react-icons/fi";

function LogoutButton() {
	const theme = useTheme();

	return (
		<IconButton
			position="fixed"
			top={4}
			right={4}
			icon={<FiLogOut />}
			aria-label="Logout"
			colorScheme="blue"
			onClick={(e) => {
				e.preventDefault();
				const cookies = new Cookies();
				cookies.remove(consts.SiteReadCookie, { path: "/" });
				window.location.href = "/";
			}}
		/>
	);
}

export default LogoutButton;
