


function removeJwtToken() {
    console.log("removeToken");
    localStorage.removeItem('mtgoUser');
}

function addJwtToken(user) {
    console.log("addJwtToken");
    localStorage.setItem('mtgoUser', JSON.stringify(user));

}

function createUser(e, t) {
    return {
        email : e,
        token : t
    }
}