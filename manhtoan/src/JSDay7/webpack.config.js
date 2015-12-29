module.exports={
    context: __dirname+'/src',
    entry:'./main.js',
    output:{
        path:__dirname+'/static',
        filename:'bundle.js'
    },
    module:{
        loaders:[
        {
            test:/\.jpg$/,
            loader:'url?limit=10000'
        },
        {
            test:/\.js$/,
            loader: "babel"
        }
        ]
    }
}