import styled, { css } from "styled-components";
import {
	compose,
	layout,
	width,
	maxWidth,
	space,
	flex,
	color,
	flexbox,
	typography,
	border,
	display,
	justifyContent,
	alignItems,
	order,
	justifySelf,
	alignSelf,
	flexDirection,
	position,
	bottom
} from "styled-system";

// Add styled-system functions to your component
export const Row = styled.div`
	display: flex;
	flex-wrap: wrap;
	width: 100%;
	${props => {
		if(!props.position == "absolute") {
			return css`
				margin-right: -15px;
				margin-left: -15px;
			`
		}
	}}

	${compose(
		color,
		layout,
		space,
		flexbox,
		typography,
		position,
		bottom
	)}
`;

export const Col = styled.div`
	padding-right: 15px;
	padding-left: 15px;
	position: relative;

	${compose(
		color,
		layout,
		space,
		typography,
		border,
		display,
		justifyContent,
		alignItems,
		order,
		justifySelf,
		alignSelf,
		flexDirection,
		position,
		bottom
	)}
`;
