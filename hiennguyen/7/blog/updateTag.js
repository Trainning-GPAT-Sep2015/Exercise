const r=require('rethinkdb');

async function connect(){
	return await r.connect('localhost');
}

async function queryArticles(){
	const conn=await connect();
	return await (await r.db('blog').table('article').run(conn)).toArray();
}

async function updateTagArticles(){
	const conn=await connect();
	const articles=await queryArticles();
	const cursor=await r.db('blog').table('tag').run(conn);
	return new Promise((resolve,reject)=>{
		cursor.each((err,tag)=>{
			r.db('blog').table('tag').get(tag.id).update({
				articles:containers(tag,articles)
			}).run(conn);
		});
		resolve();
	});
}

function containers(tag,articles){
	const results=[];
	articles.forEach(function (article){
		var content=article.content.split(' ');
		content=content.map((word)=>word.toLowerCase());
		if (content.indexOf(tag.word)>-1){
			results.push(article);
		}
	});
	var objs=[];
	results.forEach(function(article,i){
		objs[i]={};
		objs[i]['id']=article.id;
		objs[i]['title']=article.title;
	});
	return objs;
}

updateTagArticles().then(
	(result)=>console.log('Done!'),
	(error)=>console.log(error));