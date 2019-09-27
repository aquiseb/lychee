import styled from "styled-components";
import { background } from "styled-system";

const Screen = styled.div`
	min-height: 100vh;

    /* Get a split background based on responsive breaking points */
	${({ bgSplit, ...rest }) => {
		if (!bgSplit) return

		let modifiedProps = { ...rest, background: [] };

		if (typeof bgSplit == "boolean") {
			bgSplit = [1];
		}

		bgSplit.forEach((splitSize, idx) => {
			if (splitSize < 1)
				modifiedProps.background[idx] = `linear-gradient(90deg, ${rest.theme.colors.fg} ${splitSize * 100}%, ${rest.theme.colors.bg} ${(1 - splitSize) * 100}%)`;
			else modifiedProps.background[idx] = rest.theme.colors.primary;
		});

		return background(modifiedProps);
	}}
`;

const Wrapper = styled.div`
	margin: 0 auto;
	max-width: 912px;
	padding-right: 10px;
	padding-left: 10px;
`;

export default props => (
	<>
		<Screen {...props}>
			<Wrapper>{props.children}</Wrapper>
		</Screen>
	</>
);
