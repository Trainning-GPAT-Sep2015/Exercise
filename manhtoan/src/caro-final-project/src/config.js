var fs = require('fs');
var path = require('path');

// Parse Json object from ./config.json file in directory
var config = JSON.parse(fs.readFileSync(path.join(__dirname, '.././config.json')));

// auto initialize config for server
var port = config.server.port;
config.server.port = process.env.port || port;

// Initialize for others
// TODO


module.exports = config;