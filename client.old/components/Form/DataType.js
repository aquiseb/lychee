import { Form, Input, Tooltip, Icon, Cascader, Select, Checkbox, Button, AutoComplete } from "antd";

const postTypes = [
	{
		value: "product",
		label: "Product",
		children: [
			{
				value: "vehicle",
				label: "Vehicle",
				children: [
					{
						value: "car",
						label: "Car"
					}
				]
			},
			{
				value: "cloth",
				label: "Cloth",
				children: [
					{
						value: "dress",
						label: "Dress"
					}
				]
			}
		]
	},
	{
		value: "information",
		label: "Information",
		children: [
			{
				value: "business",
				label: "Business",
				children: [
					{
						value: "openingHours",
						label: "Opening Hours"
					}
				]
			}
		]
	}
];

function filter(inputValue, path) {
	return path.some(option => option.label.toLowerCase().indexOf(inputValue.toLowerCase()) > -1);
}

const DataType = props => {
	return (
		<Form.Item
			label={
				<span>
					First, let's narrow things down&nbsp;&nbsp;
					<Tooltip title="Select the correct category for the data type that you are submitting">
						<Icon type="question-circle-o" />
					</Tooltip>
				</span>
			}
		>
			{props.form.getFieldDecorator("dataType", {
				// initialValue: ["product", "vehicle", "car"],
				rules: [{ type: "array", required: true, message: "Please select an option" }]
			})(<Cascader options={postTypes} size="large" showSearch={{ filter }} onChange={props.handleChange} />)}
		</Form.Item>
	);
};

export default DataType;
