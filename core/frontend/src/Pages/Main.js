import React, {Component, useRef, useEffect, useState} from 'react';
import { Col, Container, Form, Row, Button, ListGroup } from 'react-bootstrap';
import Message from '../Components/Chat/Message';
import AddMessageForm from '../Components/Chat/AddMessageForm';
import {CHAT_URL} from "../external/config";

function ChatPage() {
    const [messages, setMessages] = React.useState([]);
    const [isConnectionOpen, setConnectionOpen] = React.useState(false);
    const [messageBody, setMessageBody] = React.useState("");

    const [error, setError] = useState('');
    const [alert, setAlert] = useState('');
    const [showAlert, setShowAlert] = useState(false);

    const ws = React.useRef(null);

    const sendMessage = (messageBody) => {
        if (messageBody.msg && messageBody.name) {
            ws.current.send(
                JSON.stringify({
                    name: messageBody.name,
                    message: messageBody.msg,
                })
            );
            setMessageBody("");
        } else {
            setError('Message is empty')
            setAlert('warning');
            setShowAlert(true);
            setTimeout(() => setShowAlert(false), 1000);
        }
    };

    React.useEffect(() => {
        ws.current = new WebSocket(CHAT_URL);

        ws.current.onopen = () => {
            setError('Connected')
            setAlert('primary');
            setConnectionOpen(true);
            setShowAlert(true);
            setTimeout(() => setShowAlert(false), 3000);
        };

        ws.current.onmessage = (event) => {
            const data = JSON.parse(event.data);
            if (data !== "channel is filled") {
                setError('Limit of messages per day')
                setAlert('info')
                setMessages((_messages) => [..._messages, data]);
                setShowAlert(true);
                setTimeout(() => setShowAlert(false), 3000);
            }

        };

        return () => {
            ws.current.close();
        };
    }, []);

    return (
        <Container>
            <Row>
                <Col className="py-3">
                    <ListGroup className="mb-0">
                        {messages.map((item, index) => (
                            <Message key={index} name={item.name} content={item.message} />
                        ))}
                    </ListGroup>

                    <AddMessageForm addNewMsg={sendMessage} messageBody={messageBody} setMessageBody={setMessageBody} />

                </Col>
            </Row>
        </Container>
    );
}

export default ChatPage;