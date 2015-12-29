module.exports = {
    context: __dirname +"/src",
    entry: "./draw.js",
    output: {
        path: __dirname+'/static',
        filename: "bundle.js"
    },
    
    module:{
        loaders:[
            {
                test:/\.css$/,
                loader:'style!css'
            },
            {
                test:/\.js$/,
                exclude: /node_modules/,
                loader:'babel'
            }
        ]
    }
};