/*
  @argument name string
  @return exports object
*/
function require(name) {

}

/*
  @argument name string
  @module function (exports, require, module, __filename, __dirname)
    @argument exports object
    @argument require function
    @argument module object
    @argument __filename string
    @argument __dirname string
*/
function define(name, module) {

}

define('/home/i/w/training/js/calc',
  function(exports, require, module, __filename, __dirname) {

    var add = require('./calc/add');
    var mul = require('./calc/mul');

    console.log(add.add(2,3,3));

});

define('/home/i/w/training/js/calc/add',
  function(exports, require, module, __filename, __dirname) {

    var mul = require('./mul');

    function add(a, b, count) {
      if (count <= 0) {
        return a + b;
      }
      return mul.mul(a, a+b, count-1);
    }

    exports.add = add;

});

define('/home/i/w/training/js/calc/mul',
  function(exports, require, module, __filename, __dirname) {

    var add = require('./add');

    function mul(a, b, count) {
      if (count <= 0) {
        return a * b;
      }
      return add.add(a, a*b, count-1);
    }

    exports.mul = mul;

});

require('/home/i/w/training/js/calc');
