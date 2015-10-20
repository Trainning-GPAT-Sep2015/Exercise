function sum(a, b, callback) {

}

function double(x, callback) {

}

function promisify(fn) {
    return (...args) => {
        return new Promise((resolve, reject) => {
            try {
                fn(...args, (error, result) => {
                    if (error) {
                        reject(error);
                    } else {
                        resolve(result);
                    }
                });
            } catch (e) {
                reject(e);
            }
        });
    };
}

const sumPromise = promisify(sum);
const doublePromise = promisify(double);

async
function calc() {
    const total = await sumPromise(2, 3);
    const d = await doublePromise(total);
    return d;
}

// 

calc().then(
    result => console.log('calc'.result),
    error => console.error('cacl error', error.toString()));



function * fibo() {
    let [a, b] = [0, 1];
    while (true) {
        yield a;
        [a, b] = [b, a + b];
    }
}

for(const n of fibo()) {
    if (n > 1000) {
        break;
    }
    console.log(n);
}