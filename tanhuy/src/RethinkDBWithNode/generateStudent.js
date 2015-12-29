var r = require('rethinkdb');

async function queryStudent() {
    const conn = await r.connect('localhost');
    const subjectsCursor = await r.db('test').table('subject').run(conn);
    const subjects = await subjectsCursor.toArray();
    const cursor = await r.db('test').table('students').run(conn);
    return {conn, subjects, cursor};
}

function generateSubjects({conn, subjects,cursor}) {
    let count = 0;
    return new Promise((resolve, reject)=>{
        cursor.each(
            (err,student) => {
                count++;
                console.log("will update subjects: ", getRandomSubjects(subjects));
                r.db('test').table('students').get(student.id).update({
                    subjects: getRandomSubjects(subjects)
                }).run(conn).then(null, err => console.error('Error', err));
            },
            ()=>resolve({count})
        );
    });
}

function getRandomSubjects(subjects) {
    const n = Math.floor(Math.random()*3)+5;
    const res = [];
    for (var i = 0; i < n; i++) {
        const index = Math.floor(Math.random()*subjects.length);
        res.push(subjects[index].id);
    }
    return res;
}

async function main() {
    const {conn, subjects, cursor}=await queryStudent();
    const {count}=  await generateSubjects({conn, subjects,cursor});
    console.log('Updates ${count} students');
}

main().then(
    () => console.log('Done!'),
    err => console.error('Error:', err.stack)
);