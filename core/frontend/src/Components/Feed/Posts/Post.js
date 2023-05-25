import axios from 'axios';
import React, {Component} from 'react';
import {Card, Button, ListGroup} from 'react-bootstrap';
import {Trash} from 'react-bootstrap-icons';
import Comment from '../Comments/Comment';
import AddCommentForm from '../Comments/AddCommentForm';
import {COMM_URL} from '../../../config';
import {createComment, createPost, deleteCommentAPI, getCommentsByPostId} from '../../../service/api';
import {getUsername, getUsernameByParseJWT} from "../../../utils";

class Post extends Component {
    constructor(props) {
        super(props);

        const {id, username, title, description} = this.props.data;
        const owner_id = this.props.data.user_id;
        const {user_id, group_id} = this.props.user;
        const comments = this.props.comments;

        this.state = {
            comments: comments,
            id: id,
            username: username,
            title: title,
            description: description,
            owner_id: owner_id,
            user_id: user_id,
            group_id: group_id,
        };
    }

    componentDidMount() {
        //this.fetchData();
    }

    async fetchComments() {
        const response = await getCommentsByPostId(this.state.post_id);
    }

    async addComment(item) {
        const post_id = this.state.id;
        const response = await createComment(post_id, item.content);
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
                const newComment = {
                    id: post_id,
                    user_id: this.state.user_id,
                    user_name: getUsernameByParseJWT(),
                    content: item.content
                }
                let prevState = this.state.comments;
                this.setState(({
                    comments: [...prevState, newComment],
                    alert: 'success',
                    error: 'Everything has been updated.',
                    needUpdate: false
                }));
                break;
            case 500:
                this.setState({
                    alert: 'error',
                    error: `Internal Server Error\n${response.error}`
                });
                break;
            default:
                this.setState({
                    alert: 'error',
                    error: `Unknown Error\n${response.error}`
                });
                break;
        }
        this.setState({showAlert: true});
        setTimeout(() => this.setState({showAlert: false}), 3000);

    }

    async removeComment(comment_id) {
        const response = await deleteCommentAPI(comment_id);
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
                let comments = this.state.comments;
                this.setState(prevState => ({
                    comments: comments.filter((comment) => comment.id !== comment_id),
                    alert: 'success',
                    error: 'Successfully removed',
                    needUpdate: false
                }));

                break;
            case 500:
                this.setState({
                    alert: 'error',
                    error: `Internal Server Error\n${response.error}`
                });
                break;
            default:
                this.setState({
                    alert: 'error',
                    error: `Unknown Error\n${response.error}`
                });
                break;
        }
        this.setState({showAlert: true});
        setTimeout(() => this.setState({showAlert: false}), 3000);
    }

    deletePost() {
        this.props.deletePost(this.state.id);
    }

    render() {
        const {comments, username, title, description, group_id, user_id, owner_id} = this.state;
        return (
            <div className='mb-4'>
                <Card style={{width: '65%', margin: 'auto'}} variant="light">
                    <Card.Header>
                        {((1 <= group_id && group_id < 3) || (user_id === owner_id && user_id >= 1)) && (
                            <Button
                                style={{float: 'right'}}
                                size="sm"
                                className="float-right"
                                variant="danger"
                                onClick={() => this.deletePost()}>
                                <Trash/>
                            </Button>
                        )}
                    </Card.Header>
                    <Card.Body>
                        <Card.Title as="h3">{title}</Card.Title>
                        <Card.Subtitle className="mb-2 mt-1 text-muted">@{username}</Card.Subtitle>
                        <Card.Text>{description}</Card.Text>
                    </Card.Body>
                    <Card.Footer className='px-0 py-0 pt-0'>
                        <ListGroup>
                            {comments &&
                                comments.map((i, pos) => {
                                    return (
                                        <Comment
                                            key={i.id}
                                            user={this.props.user}
                                            commentOwner_id={i.user_id}
                                            username={i.user_name}
                                            content={i.content}
                                            comment_id={i.id}
                                            deleteComment={(comment_id) => this.removeComment(comment_id)}
                                            position={pos}
                                        />
                                    );
                                })}
                        </ListGroup>
                    </Card.Footer>
                </Card>
                {1 <= group_id && group_id <= 3 && <AddCommentForm addComment={(item) => this.addComment(item)}/>}
            </div>
        );
    }
}

export default Post;
