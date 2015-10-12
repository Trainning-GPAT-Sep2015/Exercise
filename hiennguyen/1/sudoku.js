const st='3.16.34.5314.5.45.13.643.46.3.345.14.51.45.36.356.31.5.34.53.24.3152524.4525.22..';
const a=st.match(/.{9}/gi);
const sudoku=a.map(function(x){return x.split('');});
console.log(sudoku);