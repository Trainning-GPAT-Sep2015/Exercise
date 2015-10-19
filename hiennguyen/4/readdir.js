const fs=require('fs');
const crypto=require('crypto');

const path=process.argv[2];

// Test: babel-node readdir.js input

function promisify(fn){
	return (...args)=>{
		return new Promise((resolve,reject)=>{
			try{
				fn(...args,(error,result)=>{
					if (error){reject(error);}
					else {resolve(result);}
				});
			} catch(e){
				reject(e);
			}
		});
	}
}

var readdirPromise=promisify(fs.readdir);
var readfilePromise=promisify(fs.readFile);
readdirPromise(path).then(
	(result)=>{
		for (var file of result){
			file=path+"/"+file;
			readfilePromise(file,'utf-8').then(
				(file_result)=>{
					const obj={'original':file_result,'hashed':crypto.createHash('md5').update(file_result).digest('hex')};
					console.log(obj.original,'|',obj.hashed);
				},
				(file_error)=>console.log(file_error));
		}
	},
	(error)=>console.log(error));