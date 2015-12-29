import r from 'rethinkdb';

async function connect() {
    return await r.connect('localhost');
}

async function query() {
    const conn = await connect();
    const subjectsCursor = await r.db('test').table('subject').run(conn);
    const subjects = await subjectsCursor.toArray();

    return subjects;
}

async function query(subjects) {
    const conn = await connect();

    const students = subjects.map( async (subject) => {
        const cursor = await r.db('test').table('student')
            .getAll(subject.id, {index: 'student_subject'})
            .map((row) => ({
                id: row('id'),
                name: row('name')
            }));
        const students = await cursor.toArray();
        return students;
    });

    return await Promise.all(promises);
}