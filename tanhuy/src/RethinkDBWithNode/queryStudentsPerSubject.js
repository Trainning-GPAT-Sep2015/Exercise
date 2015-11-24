import r from 'rethinkdb';

async function connect () {
	return await r.connect('localhost');
}

async function querySubjects(){
	const conn =await connect();
	const subjectCursor = await r.db('test').table('subject').run(conn);
	const subjects = await subjectCursor.toArray();
	return subjects;
}

async function query(){
	const conn = await connect();
	const subjects = await querySubjects();
	const promises = subjects.map(async (subject)=>
		(async ()=>{
		const cursor = await r.db('test').table('students')
			.getAll(subject.id,{index:'student_subject'})
			.distinct()
			.map((row)=>({
				id: row('id'),
				name: row('name')
			})).run(conn);
		const students = await cursor.toArray();
		return students;
	})());
	const students = await Promise.all(promises);

	return subjects.map((subject,i)=>({
		subject,
		students: students[i]
	}));

}

query().then(
	()=> {
		console.log('Done!');
		result.forEach((info)=>{
			console.log('Subject ${info.subject.id} has ${info.students.length} students');
			console.log(info.students.map(s=>'${a.id}\t${s.name}').join('\n'));
		})
	},
	err=> console.error('Error: ', err.stack)
);