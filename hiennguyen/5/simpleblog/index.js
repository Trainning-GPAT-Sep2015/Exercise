const express=require('express');
const fs=require('fs');
const app=express();
const config=JSON.parse(fs.readFileSync('config.json','utf-8'));

app.set('view engine','jade');
app.set('views','./views');

const listFiles=fs.readdirSync(config.path);
var start;

app.use(function(req,res,next){
	console.log('Request Type:',req.method);
	start=Date.now();
	next();
	console.log(Date.now()-start);
});

app.get('/',function(req,res){
	res.render('index',{listfiles:listFiles});
});

app.use('/articles/:name',function(req,res,next){
	if (req.params.name!=='secret.jade'){
		next();
	}
	else {
		res.send('<title>Access denied</title>'
			+'<body>You are not allowed to see my secret page</body>');
	}
});

app.get('/articles/:name',function(req,res){
	res.render('articles/'+req.params.name);
});

const server=app.listen(config.port,function(){
	const host=server.address().address;
	const port=server.address().port;
	console.log('Listening at http://%s:%s',host,port);
});