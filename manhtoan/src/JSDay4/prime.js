function prime(n, callback) {
    for (var i = 2; i < n; i++) {
        var flag = true;
        for (var j = 2; j <= Math.floor(Math.sqrt(i)); j++) {
            if (i % j === 0) {
                flag = false;
                break;
            }
        }
        if (flag) {
            callback(i);
        }
    }
}

prime(7, func => console.log(func));

// SetInterval

// try catch pattern 
/*
try {
    ///////////
    callback(null, i)
} catch(e) {
    callback(e.toString())
}

calc(10, (err, result) => {
   if(err) {
    console.log(err.toString())
   }
   console.log(result)
});
*/