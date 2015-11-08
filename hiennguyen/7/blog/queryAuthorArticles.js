const r=require('rethinkdb');

async function connect(){
	return await r.connect('localhost');
}

async function queryAuthors(){
	const conn=await connect();
	return await (await r.db('blog').table('author').run(conn)).toArray();
}

async function queryAuthorArticles(){
	const conn=await connect();
	const authors=await queryAuthors();
	const promises=await authors.map((author)=>
		(async ()=>{
			const cursor=await r.db('blog').table('article').getAll(author.id,{index:'article_author'}).run(conn);
			return await cursor.toArray();
		})()
	);
	const articles=await Promise.all(promises);
	return authors.map((author,i)=>({
		author,
		articles:articles[i]
	}));
}

queryAuthorArticles().then(
	(result)=>{
		result.forEach((info)=>{
			console.log('List of articles written by author',info.author.name);
			info.articles.map((article)=>console.log(article.title));
		});
	},
	(error)=>console.log(error));