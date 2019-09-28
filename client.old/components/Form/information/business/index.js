import { Form, Input, Icon } from "antd";

const Business = props => {
	const { getFieldDecorator, getFieldsError, getFieldError, isFieldTouched } = props.form;
	const usernameError = isFieldTouched("business.name") && getFieldError("business.name");

	return (
		<Form.Item validateStatus={usernameError ? "error" : ""} help={usernameError || ""}>
			{getFieldDecorator("business.name", {
				rules: [{ required: true, message: "Please input your username!" }]
			})(<Input size="large" prefix={<Icon type="shop" style={{ color: "rgba(0,0,0,.25)" }} />} placeholder="Enter the name of your business" />)}
		</Form.Item>
	);
};

export default Business;
