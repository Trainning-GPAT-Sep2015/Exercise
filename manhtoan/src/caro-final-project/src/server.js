let http = require('http');
let express = require('./express');
let {port} = require('./config').server;

(async function(){
    try{
        // Initialize rethinkdb

        // Create http server with express app and start listening
        let server = http.createServer(express);
        server.listen(port, () => console.log('Server listing on port ' + port));
    } catch(err){
        console.log(err);
        process.exit(1);
    }
})();