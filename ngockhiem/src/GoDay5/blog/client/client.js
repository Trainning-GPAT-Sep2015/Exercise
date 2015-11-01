var express = require('express');
var request = require('request');
var app = express();

var token = "";
var options = {
	url: 'http://localhost:8001/'
};
app.get("/", function(req, res) {
	request(options, function(error, response, body) {
		if (!error && response.statusCode == 200) {
			res.send(body); // Show the HTML for the Google homepage. 
			res.status(response.statusCode);
		} else {
			res.send(body);
			res.status(404);
		}
	});
});
app.get("/khiem", function(req, res) {
	request("http://unkk50d62c4b.ngockhiem27.koding.io/", function(error, response, body) {
		if (!error && response.statusCode == 200) {
			res.send(body); // Show the HTML for the Google homepage. 
			res.status(response.statusCode);
		} else {
			res.send(body);
			res.status(404);
		}
	});
});
app.get("/article/:name", function(req, res) {
	name = req.params.name;
	var options = {
		url: 'http://localhost:8001/article/' + name,
		headers: {
			'Authentication': token
		}
	};
	request(options, function(error, response, body) {
		if (!error && response.statusCode == 200) {
			res.send(body); // Show the HTML for the Google homepage. 
			res.status(response.statusCode);
		} else {
			res.send(body);
			res.status(404);
		}
	});
});

app.get("/login", function(req, res) {
	request('http://localhost:8001/login', function(error, response, body) {
		if (!error && response.statusCode == 200) {
			token = body;
			res.send(body);
			res.status(404)
		}
	});
});

app.get("/login/reset", function(req, res) {
	request('http://localhost:8001/login/reset', function(error, response, body) {
		if (!error && response.statusCode == 200) {
			token = body;
			res.send(body);
			res.status(404)
		}
	});
});

var client = app.listen("3000", function() {
	var host = client.address().address;
	var port = client.address().port;
	console.log('Example app listening at http://%s:%s', host, port);
});