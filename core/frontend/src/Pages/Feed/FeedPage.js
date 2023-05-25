import React, { Component } from 'react';
import {Col, Row, Container, Alert} from 'react-bootstrap';
import Post from '../../Components/Feed/Posts/Post';
import AddPostForm from '../../Components/Feed/Posts/AddPostForm';
import {getInfoByJWT, getUserId, getUsername, isOwnerOrAdmin} from '../../utils';
import {createPost, deletePost, getPosts} from '../../service/api';

class FeedPage extends Component {
    constructor(props) {
        super(props);

        const token = localStorage.getItem('jwt_token');

        const [user_id, group_id, user_name] = getInfoByJWT(token);

        this.state = {
            posts: [],
            error: '',
            alert: '',
            needUpdate: true,
            showAlert: false,
            user_id: user_id,
            group_id: group_id,
            user_name: user_name
        };
    }

    componentDidMount() {
        this.fetchPosts();
    }

    // shouldComponentUpdate(nextProps, nextState) {
    //     const { posts, showAlert, user_id, group_id } = this.state;
    //     const { posts: nextPosts, showAlert: nextShowAlert, user_id: nextUserId, group_id: nextGroupId } = nextState;
    //
    //     return posts !== nextPosts ||
    //         showAlert !== nextShowAlert ||
    //         user_id !== nextUserId ||
    //         group_id !== nextGroupId;
    //
    //      // Не обновлять компонент, если ни одно из указанных свойств или состояний не изменилось
    // }


    fetchPosts = async () => {
        const response = await getPosts();

        if (!response) {
            this.setState({
                error: 'Server refused connection',
                alert: 'error',
                showAlert: true
            });
            setTimeout(() => this.setState({ showAlert: false }), 3000);
            return;
        }

        const status = response.status;

        switch (status) {
            case 200:
                const data = response.data.data;
                if (data) {
                    this.setState({
                        posts: data,
                        alert: 'success',
                        error: 'Everything has been updated.',
                        needUpdate: false
                    });
                } else {
                    this.setState({
                        alert: 'primary',
                        error: 'Posts not found.'
                    });
                }
                break;
            case 500:
                this.setState({
                    alert: 'error',
                    error: 'Internal Server Error'
                });
                break;
            default:
                this.setState({
                    alert: 'error',
                    error: 'Unknown Error'
                });
                break;
        }
        this.setState({ showAlert: true });
        setTimeout(() => this.setState({ showAlert: false }), 3000);
    };


    addPost = async (post) => {

        const response = await createPost(post.title, post.description)
        if (!response) {
            this.setState({
                error: 'Server refused connection',
                alert: 'error',
                showAlert: true
            });
            setTimeout(() => this.setState({ showAlert: false }), 3000);
            return;
        }

        const status = response.status;

        switch (status) {
            case 200:
                const {post_id, user_id} = response.data;
                if (post_id) {
                    const newPost = { id: post_id,
                        title: post.title,
                        description: post.description,
                        user_id: user_id,
                        username: getUsername(),
                        comms: { data: [] } };
                    this.setState(prevState => ({
                        posts: [...prevState.posts, newPost],
                        alert: 'success',
                        error: 'Everything has been updated.',
                        needUpdate: false
                    }));
                } else {
                    this.setState({
                        alert: 'primary',
                        error: 'Posts was not create'
                    });
                }
                break;
            case 500:
                this.setState({
                    alert: 'error',
                    error: `Internal Server Error${response.error}`
                });
                break;
            default:
                this.setState({
                    alert: 'error',
                    error: 'Unknown Error'
                });
                break;
        }
        this.setState({ showAlert: true });
        setTimeout(() => this.setState({ showAlert: false }), 3000);
    };

    removePost = async (id) => {

        const response = await deletePost(id)

        if (!response) {
            this.setState({
                error: 'Server refused connection',
                alert: 'error',
                showAlert: true
            });
            setTimeout(() => this.setState({showAlert: false}), 3000);
            return;
        }

        const status = response.status;

        switch (status) {
            case 200:
                const delete_status = response.data;
                if (delete_status) {
                    let posts = this.state.posts;
                    this.setState(prevState => ({
                        posts: posts.filter((post) => post.id !== id),
                        alert: 'success',
                        error: 'Successfully removed',
                        needUpdate: false
                    }));
                } else {
                    this.setState({
                        alert: 'primary',
                        error: 'Failed to delete post'
                    });
                }
                break;
            case 500:
                this.setState({
                    alert: 'error',
                    error: `Internal Server Error${response.error}`
                });
                break;
            default:
                this.setState({
                    alert: 'error',
                    error: 'Unknown Error'
                });
                break;
        }
        this.setState({showAlert: true});
        setTimeout(() => this.setState({showAlert: false}), 3000);
    };

    render() {
        const { posts, showAlert, username, user_id, group_id, alert, error } = this.state;
        const user = { user_id, group_id };
        const [ownerId, groupId] = isOwnerOrAdmin();
        const currentUser = { user_id: ownerId, group_id: groupId };
        return (
            <Container style={{ margin: 'auto' }}>
                <Row>
                    <Col className="py-5">
                        {showAlert && (
                            <Alert variant={alert} key={error} className="mb-3" style={{ width: "65%", margin: "auto"}}>
                                {error}
                            </Alert>
                        )}
                        {(groupId && (groupId === 3 || groupId === 1))  && (
                            <AddPostForm addPost={this.addPost} />
                        )}
                        {posts &&
                            posts.map((post) => {
                                const comments = post.comms.data !== null ? post.comms.data : [];

                                return (
                                    <Post
                                        key={post.id}
                                        data={post}
                                        username={username}
                                        user={currentUser}
                                        comments={comments}
                                        deletePost={this.removePost}
                                        editPost={this.editPost}
                                    />
                                );
                            })}

                    </Col>
                </Row>
            </Container>
        );
    }
}

export default FeedPage;
