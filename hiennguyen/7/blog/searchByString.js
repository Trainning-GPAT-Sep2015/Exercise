const r=require('rethinkdb');

async function connect(){
	return await r.connect('localhost');
}

async function searchByString(str){
	const conn=await connect();
	const cursor=await r.db('blog').table('tag').run(conn);
	const words=str.split(' ').map((word)=>word.toLowerCase());
	return new Promise((resolve,reject)=>{
		var results=[];
		words.forEach(function(word){
			cursor.each((err,tag)=>{
				if (tag.word===word){
					const titles=tag.articles.map((article)=>article.title);
					results=results.concat(titles);
				}
			});
		});
		resolve(removeDup(results));
	});
}

function removeDup(arr){
	var obj={};
	var out=[];
	for (var i=0;i<arr.length;i++){
		obj[arr[i]]=0;
	}
	for (i in obj){
		out.push(i);
	}
	return out;
}

searchByString('sample article').then(
	(result)=>{
		console.log('Done!');
		console.log(result);
	},
	(error)=>console.log(error));