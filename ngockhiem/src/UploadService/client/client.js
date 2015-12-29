import express from 'express';

const app = express();


app.set('views', './views');

app.set('view engine', 'jade');
app.use(express.static(__dirname + '/views'));
app.get('/', function(req, res) {
	res.render('index', { title: 'Upload Service', message: 'Choose file'});
	res.status(202).end();
});

const client = app.listen(3000, function() {
	const host = client.address().address;
	const port = client.address().port;

	console.log('Example app listening at http://%s:%s', host, port);
});