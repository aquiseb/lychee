import React from "react";
import styled, { ThemeContext } from "styled-components";
import { Row, Col } from "@ui/Grid";
import Form from "@comps/Form";
import Layout from "@comps/Layout";

import { color, system } from "styled-system";

const Title = styled.h1`
	color: ${({ theme }) => theme.colors.primary};
`;

const Shadow = styled(Row)`
	height: 1px;
	margin-top: -1px;
	position: relative !important;
	background-color: #ccc;

	&::before {
		content: "" !important;
		height: 18px !important;
		box-shadow: 0 -9px 15px -5px rgba(0, 0, 0, 0.09) !important;
		border-radius: 75% !important;
		display: block !important;
		position: absolute !important;
		width: 100% !important;
		z-index: -1 !important;
	}
`;

const BottomNavigation = styled(Row)`
	background-color: ${props => props.theme.colors.fg};
`;

const Home = props => {
	const [values, setValues] = React.useState({ values: {} });
	const theme = React.useContext(ThemeContext);
	console.log("values --", values);
	return (
		<Layout bgSplit={[1, 2 / 3]}>
			<Row>
				<Col width={[1, 2 / 3]} padding={30}>
					<Title>What kind of data are you listing ?</Title>
				</Col>
			</Row>
			<Row>
				<Col width={[1, 2 / 3]} padding={30} paddingBottom={60}>
					<Form
						handleChange={setValues}
						render={rp => {
							return (
								<>
									<BottomNavigation
										position="fixed"
										bottom={0}
										left={0}
										height={60}
										zIndex={111}
										width={[1, 2 / 3]}
										alignItems="center"
										justifyContent="flex-end"
									>
										<Shadow width={[1, 3.9 / 5]} />
										<Row width={[1, 4 / 5]} bg={theme.colors.fg} paddingLeft={[0, 60]} height="100%" alignItems="center">
											{rp.children}
										</Row>
									</BottomNavigation>
								</>
							);
						}}
					/>
				</Col>
				<Col width={[1, 1 / 3]} position={["block", "fixed"]} right={["auto", 0]} padding={30}>{values.dataType && ( <p>You are listing the {values.dataType[2]} of your {values.dataType[1]}.</p> )}</Col>
			</Row>
		</Layout>
	);
};

export default Home;
