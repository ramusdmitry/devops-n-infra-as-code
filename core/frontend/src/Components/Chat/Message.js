import React, { Component } from 'react';
import { ListGroup } from 'react-bootstrap';

class Message extends Component {
    render() {
        const { name, content } = this.props
        return (
            <ListGroup.Item className='py-2' style={{ width: "65%", margin: "auto"}}>
                <strong>{name}: </strong>
                <span>{content}</span>
            </ListGroup.Item>

        );
    }
}

export default Message;