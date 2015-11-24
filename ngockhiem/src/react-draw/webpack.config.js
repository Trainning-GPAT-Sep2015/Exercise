module.exports = {
	context: __dirname + "/src",
	entry: "./app.js",
	output: {
		path: __dirname + "/dist",
		filename: "bundle.js"
	},
	module:
	{
		loaders: [
			{test: /\.(js|jsx)$/, exclude: /(node_modules|bower_components)/, loader: 'babel'},
			{test: /\.(jpg|png|jpeg)$/, loader: "file"}
		]
	}

}