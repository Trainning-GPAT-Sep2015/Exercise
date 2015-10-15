var string = '.12345678.87654321.12345678.87654321.12345678.87654321.12345678.87654321.12345678';

var sudoku = string.match(/.{9}/g).map(function(x) {
	return x.split('').join(' ');
}).join('\n');

console.log(sudoku);