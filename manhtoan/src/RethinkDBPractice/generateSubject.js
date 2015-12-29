'use strict';

import r from 'rethinkdb';

/////////////////// TASK 1 ///////////////////////////

// async function generate(){
//     const conn = await r.connect('localhost');
//     const result = await r.db('test').table('student').run(conn);

//     console.log('Students', await result.toArray());
// }

// generate().then(
//     () => console.log('Done!'),
//     err => console.log('Error:', err.stack));


/////////////////// TASK 2 /////////////////////////////

async function queryStudent(){
    const conn = await r.connect('localhost');

    const subjectsCursor = await r.db('test').table('subject').run(conn);
    const subjects = await subjectsCursor.toArray();

    const cursor = await r.db('test').table('student').run(conn);
    return {conn, subjects, cursor};
}

function generateSubjects({conn, subjects, cursor}) {
    let count = 0;
    return new Promise((resolve, reject) => {
        cursor.each(student => {
            count++;
            console.log('Will update subjects', getRandomSubjects(subjects));
            // r.db('test').table('student').get(student.id)
            //     .update({
            //     subjects: getRandomSubjects(subjects)
            //     }).run(conn);
        }, () => resolve({ count }));
    });
}

function getRandomSubjects(subjects) {
    const n = Math.floor(Math.random() * 3) + 5;
    const res = [];
    for(let i = 0; i < n; i++) {
        const index = Math.floor(Math.random() * subjects.length);
        res.push(subjects[index].id);
    }
}

async function main() {
    const {conn, subjects, cursor} = await queryStudent();
    const { count } = await generateSubjects({conn, subjects, cursor});
    console.log('Update ${count} students');
}

main().then(
    () => console.log('Done!'),
    err => console.log('Error:', err.stack));
