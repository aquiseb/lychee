import { Flex, FormControl, FormLabel, FormErrorMessage, Button, Input, Text } from "@chakra-ui/react";
import { useForm } from "react-hook-form";

import Cookies from "universal-cookie";
import consts from "consts";

function Login(props) {
	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm();

	const onSubmit = (data) => {
		console.log(data);

		const date = new Date();
		const days = 1;
		const oneDay = date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);

		const cookies = new Cookies();
		cookies.set(consts.SiteReadCookie, data.password, {
			path: "/",
			maxAge: oneDay,
		});

		window.location.href = props.redirectPath ?? "/";
	};

	return (
		<Flex justifyContent="center" textAlign="center" alignItems="center" h="100vh">
			<form onSubmit={handleSubmit(onSubmit)}>
				<FormControl isInvalid={errors.password}>
					<FormLabel htmlFor="password">
						<Text align="center">The password is: lychee</Text>
					</FormLabel>
					<Input
						id="password"
						type="password"
						placeholder="password"
						textAlign="center"
						{...register("password", {
							required: "Please enter the password",
							minLength: 3,
							maxLength: 100,
						})}
					/>
					<Flex justifyContent="center">{errors.password && <FormErrorMessage>{errors.password.message}</FormErrorMessage>}</Flex>
				</FormControl>

				<Button mt={4} borderRadius="md" bg="gray.400" _hover={{ bg: "gray.300" }} variant="ghost" type="submit" color="white">
					GO
				</Button>
			</form>
		</Flex>
	);
}

export default Login;
