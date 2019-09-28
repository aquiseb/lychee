import React from "react";
import { StyleSheet, Text, View } from "react-native";
// import { App as Appli } from "@lychee/core"

const coucou = "yooooooooo";

export default class App extends React.Component {
	render() {
		return (
			<View style={styles.container}>
				<Text>Coucou Rose!!!!!</Text>
			</View>
		);
	}
}

const styles = StyleSheet.create({
	container: {
		flex: 1,
		backgroundColor: "#fff",
		alignItems: "center",
		justifyContent: "center"
	}
});
