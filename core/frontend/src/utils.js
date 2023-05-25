export function parseJwt(token) {

    if (token == null || token === "") return null

    let base64Url = token.split('.')[1];
    let base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    let jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function (c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
}

export function getInfoByJWT(token) {
    if (token == null || token === "") return [null, null, null]
    return [token.user_id, token.group_id, token.user_name];
}

export function makeRandomName(length) {
    let result = '';
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    const charactersLength = characters.length;
    for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}

export function getAuth() {
    const token = localStorage.getItem('jwt_token');

    let currentTime = new Date().getTime();

    if (!token) {
        return [false, null];
    }

    const tokenData = parseJwt(token);
    if (tokenData.exp <= currentTime / 1000) {
        return [false, null];
    }

    if (!tokenData.user_name || tokenData.user_id === 0 || !(1 <= tokenData.group_id && tokenData.group_id <= 4)) {
        return [false, null];
    }
    return [true, tokenData];

}

export function getUsername(){
    const token = localStorage.getItem('jwt_token');
    if (!token) {
        const user_name = localStorage.getItem('user_name');
        if (!user_name) {
            localStorage.setItem('user_name', makeRandomName(8));
        }
        return localStorage.getItem('user_name');
    }
    return parseJwt(token)['user_name'];
}

export function getUsernameByParseJWT(){
    const token = localStorage.getItem('jwt_token');
    return parseJwt(token)['user_name'];
}

export function isOwnerOrAdmin() {
    const token = localStorage.getItem('jwt_token')
    if (token) {
        const data = parseJwt(token);
        return [data['user_id'], data['group_id']]
    }
    return [null, null];
}


// maybe unused
export function getUserId() {
    const token = localStorage.getItem('jwt_token')
    if (token) {
        return parseJwt(token)['user_id']
    } return 0
}