import React, { Component } from 'react';
import { Button, Row, Col, ListGroup } from 'react-bootstrap';
import { Trash } from 'react-bootstrap-icons';
import {getInfoByJWT} from "../../../utils";

class Comment extends Component {

    constructor(props) {
        super(props);

        const token = localStorage.getItem('jwt_token');
        const { username, content } = this.props;
        const { user_id, group_id } = this.props.user;
        const owner_id = this.props.commentOwner_id;

        this.state = {
            error: '',
            alert: '',
            needUpdate: true,
            showAlert: false,
            user_id: user_id,
            group_id: group_id,
            user_name: username,
            content: content
        };
    }

    deleteComment = () => {
        this.props.deleteComment(this.props.comment_id, this.props.position);
    }


    render() {

        const { user_id, group_id } = this.props.user;
        const owner_id = this.props.commentOwner_id;
        return (
            <ListGroup.Item>
                <Row>
                    <Col>
                        <strong>{this.state.user_name}</strong>
                        <span>: {this.state.content}</span>
                    </Col>
                    <Col>
                        {
                            ( (1 <= group_id && group_id < 3) || (user_id === owner_id && user_id >= 1)) &&
                            
                            <Button
                                style={{ float: 'right' }}
                                size="sm"
                                className="float-right"
                                variant="danger"
                                onClick={this.deleteComment}><Trash />
                            </Button>
                        }
                    </Col>
                </Row>
            </ListGroup.Item>
        );
    }
}

export default Comment;