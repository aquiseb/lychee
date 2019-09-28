import { Form, Input, Tooltip, Icon, Cascader, Select, Checkbox, Button, AutoComplete, InputNumber } from "antd";

import { Row, Col } from "@ui/Grid";
import DataType from "./DataType";

import dynamic from "next/dynamic";

const sleep = ms => {
	return new Promise(resolve => setTimeout(resolve, ms));
};

class RegistrationForm extends React.Component {
	state = {
		confirmDirty: false,
		allComps: { Module_0: () => null, Module_1: () => null },
		allComponents: [() => null]
	};

	handleSubmit = e => {
		e.preventDefault();
		this.props.form.validateFieldsAndScroll((err, values) => {
			if (!err) {
				console.log("Received values of form: ", values);
			}
		});
	};

	handleConfirmBlur = e => {
		const { value } = e.target;
		this.setState({ confirmDirty: this.state.confirmDirty || !!value });
	};

	compareToFirstPassword = (rule, value, callback) => {
		const { form } = this.props;
		if (value && value !== form.getFieldValue("password")) {
			callback("Two passwords that you enter is inconsistent!");
		} else {
			callback();
		}
	};

	validateToNextPassword = (rule, value, callback) => {
		const { form } = this.props;
		if (value && this.state.confirmDirty) {
			form.validateFields(["confirm"], { force: true });
		}
		callback();
	};

	handleChange = pathArray => {
		const props = { pathArray };
		if (!props.pathArray || !(props.pathArray || []).length) {
			return this.setState({ allComps: [] }, () => this.props.handleChange({}));
		}

		let allPaths = [];
		let allComps = {};

		for (let idx = 0; idx < props.pathArray.length; idx++) {
			if (allPaths.length > 0) allPaths[idx] = allPaths[idx - 1] + "/" + props.pathArray[idx];
			else allPaths[0] = props.pathArray[0];

			allComps["Module_" + idx] = dynamic(() => import("./" + allPaths[idx]));
		}

		this.setState({ allComps }, () => {
			this.props.handleChange(this.props.form.getFieldsValue());
		});
	};

	getComps = () => {
		return Object.keys(this.state.allComps).map((key, i) => {
			const Comp = this.state.allComps[key];
			return <Comp key={`comp_${i}`} form={this.props.form} />;
		});
	};

	render() {
		return (
			<Form layout="vertical" onSubmit={this.handleSubmit}>
				<DataType form={this.props.form} handleChange={this.handleChange} />
				{/* {values["dataType"] && this.getDynamicComponent({ pathArray: values["dataType"], form: this.props.form })} */}
				{this.getComps()}
				{this.props.render({
					children: (
						<Row justifyContent="space-between">
							<Col>
								<Button type="link">&lt; Back</Button>
							</Col>
							<Col>
								<Button type="primary" htmlType="submit">
									Register
								</Button>
							</Col>
						</Row>
					)
				})}
			</Form>
		);
	}
}

const WrappedRegistrationForm = Form.create({ name: "register" })(RegistrationForm);
export default WrappedRegistrationForm;
