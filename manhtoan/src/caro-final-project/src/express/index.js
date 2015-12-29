let fs = require('fs');
let path = require('path');

let express = require('express');
let bodyParser = require('body-parser');

// Create an express app with router and parser
let app = express();

// app middleware
app.use(bodyParser.urlencoded({extended: false}));
app.use(bodyParser.json());

// Router
app.use('/', function(req, res){
    res.send('Hello world');
});

// Router middleware
app.use(function(req, res, next){
    let err = new Error('Content not found');
    err.code = 404;
    next(err);
});

// application level error handler
app.use(function(err, req, res, next){
    if(err.code !== 404){
        console.log(err);
    }
    // TODO - Complete error handler
    // let code = err.code || 500;
});

module.exports = app;