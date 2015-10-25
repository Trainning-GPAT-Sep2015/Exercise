'use strict';
var express = require('express');
var controller = express();
var bodyParser = require('body-parser');
var fs = require('fs');
var marked = require('marked');
var path = require('path');
var jade = require('jade');

/**********************READ CONFIG FILE*************************/
var config;

function readConfig(path) {
    var configPath = path + 'config.json';
    var config_json = fs.readFileSync(configPath, 'utf8');
    config = JSON.parse(config_json);
};

readConfig('./');

/***************************************************************/

/************************MIDDLEWARES****************************/

function authMiddleware(req, res, next) {
    if (req.params.name === 'secret.md') {
        res.status(401).send('Denied access');
    } else {
        next();
    }
}

/***************************************************************/

/************************CHECK LOGIN****************************/
var isLogin = false;



/***************************************************************/

/***********************CONTROLLER******************************/
controller.get('/', function(req, res) {
    var files = fs.readdir(config.server.path.database, function(err, files) {
        if (err) {
            console.log(err);
            res.status(404).send('Not found directory');
        }

        // Sort file md only
        var files_md = new Array();
        for (let file of files) {
            if (file.indexOf('.md') !== -1) {
                files_md.push(file);
            }
        }

        var indexPage = 'index.jade';
        var homeData = {
            title: 'Welcome to my Simple Blog',
            isLogin: isLogin,
            user: {
                name: 'User1'
            },
            file_title: 'List all articles',
            files: files_md,
            author: {
                facebook: 'nmt0504',
                name: 'ManhToan'
            }
        };
        var absTempPath = path.join(config.server.path.views, indexPage);
        fs.readFile(absTempPath, 'utf8', function(err, content) {
            if (err) {
                console.log(err);
                res.status(404).send('404 Not Found');
            }
            var html = jade.render(content, homeData);
            res.status(200).send(html);
        });
    });
});

// define authMiddleware
controller.use('/articles/:name', authMiddleware);

controller.get('/articles/:name', function(req, res) {
    var absPath = path.join(config.server.path.database, req.params.name);
    var file_content = fs.readFile(absPath, 'utf8', function(err, content) {
        if (err) {
            console.log(err);
            res.status(404).send('Can not find article');
        } else {
            res.send(marked(content));
        }
    });
});

controller.get('/login',function(req, res) {
    var loginPage = 'loginform.jade';
    var loginData = {
        title: 'Log In',
        author: {
            facebook: 'nmt0504',
            name: 'ManhToan'
        }
    };
    var absTempPath = path.join(config.server.path.views, loginPage);
    fs.readFile(absTempPath, 'utf8', function(err, content) {
        if (err) {
            console.log(err);
            res.status(404).send('404 Not Found');
        }
        var html = jade.render(content, loginData);
        res.status(200).send(html);
    });
});

controller.use('/login', bodyParser.urlencoded({
    extended: true
}));

controller.post('/login', function(req, res) {
    console.log(req.body);
    res.send('success');
});
/**************************************************************/

var server = controller.listen(config.server.port, function() {
    var host = config.server.addr;
    var port = config.server.port;

    console.log('Example app listening at http://%s:%s', host, port);
});