import axios from "axios";
import {AUTH_URL, COMMS_URL, GROUP_URL, POST_URL} from '../external/config'
import {COMM_URL} from "../config";

export const loginUser = async (username, password) => {

    const json_data = JSON.stringify(
        {
            username: username,
            password: password
        }
    )

    try {
        return await axios.post(`${AUTH_URL}/sign-in`, json_data);
    } catch (error) {
        return error.response;
    }
}

export const registerUser = async (name, username, password) => {

    const json_data = JSON.stringify(
        {
            name: name,
            username: username,
            password: password
        }
    )

    try {
        return await axios.post(`${AUTH_URL}/sign-up`, json_data);
    } catch (error) {
        return error.response;
    }
};

export const getPosts = async () => {
    try {
        return await axios.get(POST_URL);
    } catch (error) {
        return error.response;
    }
}

export const createPost = async (title, description) => {
    const token = localStorage.getItem('jwt_token');
    if (!token) {
        return {
            data: null,
            status: 401,
            statusText: 'Unauthorized',
            message: 'Please log in to access this resource'
        };
    }

    const json_data = {
        title: title,
        description: description
    }

    const headers = {
        ContentType: 'application/json',
        Authorization: `Bearer ${token}`,
    }

    try {
        return await axios.post(POST_URL, json_data, { headers: headers});
    } catch (error) {
        return error.response;
    }
}

export const editPost = async (id, title, description) => {
    const json_data = {
        title: title,
        description: description
    }

    const headers = {
        ContentType: 'application/json',
        Authorization: `Bearer ${localStorage.getItem('jwt_token')}`,
    }

    try {
        return await axios.put(`${POST_URL}/${id}`, json_data, headers);
    } catch (error) {
        return error.response;
    }
}

export const deletePost = async (id) => {
    const token = localStorage.getItem('jwt_token');
    if (!token) {
        return {
            data: null,
            status: 401,
            statusText: 'Unauthorized',
            message: 'Please log in to access this resource'
        };
    }

    const headers = {
        ContentType: 'application/json',
        Authorization: `Bearer ${token}`,
    }

    try {
        return await axios.delete(`${POST_URL}/${id}`, { headers: headers});
    } catch (error) {
        return error.response;
    }
}

export const getUsers = async () => {

    const token = localStorage.getItem('jwt_token');
    if (!token) {
        return {
            data: null,
            status: 401,
            statusText: 'Unauthorized',
            message: 'Please log in to access this resource'
        };
    }

    const headers = {
        'Content-type': 'application/json',
        Authorization: `Bearer ${token}`,
    }

    try {
        return await axios.get(GROUP_URL, {
            headers: headers
        });
    } catch (error) {
        return error.response;
    }
}

export const applyUsers = async (users) => {
    const token = localStorage.getItem('jwt_token');
    if (!token) {
        return {
            data: null,
            status: 401,
            statusText: 'Unauthorized',
            message: 'Please log in to access this resource'
        };
    }

    const headers = {
        'Content-type': 'application/json',
        Authorization: `Bearer ${token}`,
    }

    const json_data = JSON.stringify({
        data: users
    })

    try {
        return await axios.put(GROUP_URL, json_data,{
            headers: headers
        });
    } catch (error) {
        return error.response;
    }
}

export const deleteUsers = async () => {
    const token = localStorage.getItem('jwt_token');
    if (!token) {
        return {
            data: null,
            status: 401,
            statusText: 'Unauthorized',
            message: 'Please log in to access this resource'
        };
    }

    const headers = {
        'Content-type': 'application/json',
        Authorization: `Bearer ${token}`,
    }

    try {
        return await axios.delete(GROUP_URL, {
            headers: headers
        });
    } catch (error) {
        return error.response;
    }
}


export const getComments = async () => {

}

export const getCommentsByPostId = async(post_id) => {
    try {
        return await axios.get(    `${COMM_URL}/post/${post_id}`);
    } catch (error) {
        return error.response;
    }
}

export const createComment = async (post_id, comment) => {
    const token = localStorage.getItem('jwt_token');
    if (!token) {
        return {
            data: null,
            status: 401,
            statusText: 'Unauthorized',
            message: 'Please log in to access this resource'
        };
    }

    const json_data = {
        post_id: post_id,
        content: comment
    }

    const headers = {
        ContentType: 'application/json',
        Authorization: `Bearer ${token}`,
    }

    try {
        return await axios.post(`${COMMS_URL}/add/${post_id}`, json_data, { headers: headers});
    } catch (error) {
        return error.response;
    }

}

export const editComment = async (comment) => {

}

export const deleteCommentAPI = async (id) => {
    const token = localStorage.getItem('jwt_token');
    if (!token) {
        return {
            data: null,
            status: 401,
            statusText: 'Unauthorized',
            message: 'Please log in to access this resource'
        };
    }

    const headers = {
        ContentType: 'application/json',
        Authorization: `Bearer ${token}`,
    }

    try {
        return await axios.delete(`${COMMS_URL}/${id}`, { headers: headers});
    } catch (error) {
        return error.response;
    }
}