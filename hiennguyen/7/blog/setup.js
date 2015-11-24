const r=require('rethinkdb');

async function connect(){
	return await r.connect('localhost');
}

async function create(){
	const conn=await connect();
	const db=await r.dbCreate('blog').run(conn);
	const articleTable=await r.db('blog').tableCreate('article').run(conn);
	const authorTable=await r.db('blog').tableCreate('author').run(conn);
	const tagTable=await r.db('blog').tableCreate('tag').run(conn);
	return await Promise.all([db,articleTable,authorTable,tagTable]);
}

// create().then(
// 	(result)=>console.log('Done!'),
// 	(error)=>console.log(error));

async function insert(){
	const conn=await connect();
	const articles=await r.db('blog').table('article').insert([
		{
			title:'Article 1',
			content:'This is the very first article'
		},
		{
			title:'Article 2',
			content:'This is the second sample article'
		},
		{
			title:'Article 3',
			content:'Another article'
		},
		{
			title:'Article 4',
			content:'Just another sample'
		},
		{
			title:'Article 5',
			content:'The last article'
		}]).run(conn);
	const authors=await r.db('blog').table('author').insert([
		{
			name:'Hung Vuong'
		},
		{
			name:'Dam Khiem'
		}]).run(conn);
	return await Promise.all([articles,authors]);
}

// insert().then(
// 	(result)=>console.log('Done!'),
// 	(error)=>console.log(error));

async function index(){
	const conn=await connect();
	const article_author=await r.db('blog').table('article')
		.indexCreate('article_author',r.row('author'))
		.run(conn);
	const article_tags=await r.db('blog').table('article')
		.indexCreate('article_tags',r.row('tags'),{multi:true})
		.run(conn);
	const author_articles=await r.db('blog').table('author')
		.indexCreate('author_articles',r.row('articles'),{multi:true})
		.run(conn);
	const tag_articles=await r.db('blog').table('tag')
		.indexCreate('tag_articles',r.row('articles'),{multi:true})
		.run(conn);
	return await Promise.all([article_author,article_tags,author_articles,tag_articles]);
}

index().then(
	(result)=>console.log('Done!'),
	(error)=>console.log(error));