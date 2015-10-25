async

function fetchInfo(userid) {
    const user = await fetchUser(userid);
    const _class = await fetchClass(user.classid);
    return {user, class: _class};
}

fetchInfo.then(
    result => console.log('Class:', result),
    error => console.log('Error:', error);
)