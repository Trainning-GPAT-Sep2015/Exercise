'use strict';

delete Array.prototype.forEach;

(function() {
    if (typeof Array.prototype.forEach === 'function') {
        array.forEach(callback);
    } else {
        Array.prototype.forEach = function(callback) {
            for(var i = 0; i < this.length; i++) {
                callback(this[i]);
            }
        }
    }
})();

[1, 2, 3].forEach(function(a) {
    console.log(a);
});