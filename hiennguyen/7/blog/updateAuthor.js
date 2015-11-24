const r=require('rethinkdb');

async function connect(){
	return await r.connect('localhost');
}

async function queryAuthors(){
	const conn=await connect();
	return await (await r.db('blog').table('author').run(conn)).toArray();
}

async function generateArticleAuthor(){
	const conn=await connect();
	const authors=await queryAuthors();
	const cursor=await r.db('blog').table('article').run(conn);
	return new Promise((resolve,reject)=>{
		cursor.each((err,article)=>{
			r.db('blog').table('article').get(article.id).update({
				author:randomAuthor(authors)
			}).run(conn);
		});
		resolve();
	});
}

function randomAuthor(authors){
	const index=Math.floor(Math.random()*authors.length);
	return {
		id:authors[index].id,
		name:authors[index].name
	};
}

generateArticleAuthor().then(
	(result)=>console.log('Done!'),
	(error)=>console.log(error));