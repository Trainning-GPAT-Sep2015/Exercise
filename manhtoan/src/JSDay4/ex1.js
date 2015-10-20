function foo(callback) {
    console.log("foo")
    setTimeout(() => callback(1), 500);
}

function bar(callback) {
    console.log("bar")
    setTimeout(() => callback(1), 300);
}

// incorrect
/*
function wait(callback, ...actions) {
    for (var i = 0; i < actions.length; i++) {
        actions[i]();
    }
    callback();
}
*/

function wrap(fn , ...args) {
    return (callback) => {
        fn(...args, callback);
    }
}

//correct
function wait(callback, ...actions) {
    var count = 0;
    for (const action of actions) {
        action(() => {
            count++;
            if (count === actions.length) {
                callback();
            }
        })
    }
}

const start = new Date();
wait(() => {
    console.log('Done!', new Date() - start);
}, foo, bar);