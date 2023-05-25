import { Component } from "react";
import { Form, Button, Row, Col, Container } from "react-bootstrap";
import { ArrowRight } from "react-bootstrap-icons";

export default class AddCommentForm extends Component {

    state = {
        user_name: localStorage.getItem('username'),
        content: "",
    };

    handleContentChange = (e) => {
        this.setState({
            content: e.target.value,
        });
    };

    createComment = () => {
        const item = {
            user_name: this.state.user_name,
            content: this.state.content,
        }

        this.props.addComment(item)
        this.setState({
            content: ''
        })
    }

    render() {
        return (
            <Row className='mt-3' style={{ width: "65%", margin: "auto"}}>
                <Col className='col'>
                    <Form>
                        <Form.Group controlId="forContent">
                            <Form.Control placeholder='Comment...' onChange={this.handleContentChange} value={this.state.content} />
                        </Form.Group></Form>
                </Col>
                <Col className='col-lg-1'>
                    <Button variant="primary" onClick={this.createComment}>
                        <span>
                            <ArrowRight className="align-middle" />
                        </span>
                    </Button>
                </Col>
            </Row>
        );
    }


}