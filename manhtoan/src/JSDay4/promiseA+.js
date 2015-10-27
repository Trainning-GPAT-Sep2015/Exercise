function promiseDouble(x) {
    return {
        // @callback(result)
        then(resultCallback, errorCallback) {
            resultCallback(x * 2);
        }
    }
}

function promiseSum(a, b) {
    return {
        then(resultCallback, errorCallback) {
            resultCallback(a, b);
        }
    }
}

promiseDouble(4).then(
    (result) => console.log('double', result));

promiseSum(2, 3).then(
    (result) => console.log('sum', result));

/////////////////////////////////////////////////////////

function sum3(a, b) {
    return new Promise((resolve, reject) => {
        setTimeout(() => resolve(a + b), 500);
    });
}

function double3(x) {
    return new Promise((resolve, reject) => {
        setTimeout(() => resolve(x*2), 300);
    });
}

// incorrect
// function wait(...promises) {
//     var count = 0;
//     for(const promise of promises) {
//         promise
//             .then((result) => {
//                 console.log(result)
//             })
//     }
// }

// correct
function wait(...promises) {
    let count = 0;
    const results = new Array();
    return new Promise((resolve, reject) => {
        var count = 0;
        for (const promise of promises) {
            promise.then((result) => {
                count++;
                results.push(result);
                if (count === promises.length) {
                    resolve(results);
                }
            }, (error) => {
                reject(error);
            });
        }
    });
}

const start = new Date();

// Replace wait above
Promise.all([sum3(2, 3), double3(4)]);

wait(sum3(2,3), double3(4))
    .then((result) => {
        console.log('Done', new Date() - start)
    });

sum3(5, 6)
    .then((result) => {
        console.log('sum3', result);
        return result * 10;
    }).then((result) => {
        console.log('sum3', result);
    });

/////////////////////////////////////////////////////////
function promisify(fn) {
    return (...args) => {
        return new Promise((resolve, reject) => {
            try {
                fn(...args, (error, result) => {
                    if(error) {
                        reject(error);
                    } else {
                        resolve(result);
                    }
                });
            } catch(e) {
                reject(e);
            }
        });
    };
}