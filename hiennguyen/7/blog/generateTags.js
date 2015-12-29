const r=require('rethinkdb');

async function connect(){
	return await r.connect('localhost');
}

async function generateTags(){
	const conn=await connect();
	const cursor=await r.db('blog').table('article').run(conn);
	return new Promise((resolve,reject)=>{
		var tags=[];
		cursor.each((err,article)=>{
			var words=article.content.split(' ');
			words=words.map((word)=>word.toLowerCase());
			tags=tags.concat(words);
		});
		tags=removeDup(tags);
		var objs=[];
		tags.forEach(function(tag,i){
			objs[i]={};
			objs[i]['word']=tag;
		});
		r.db('blog').table('tag').insert(objs).run(conn);
		resolve();
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

generateTags().then(
	(result)=>console.log('Done!'),
	(error)=>console.log(error));