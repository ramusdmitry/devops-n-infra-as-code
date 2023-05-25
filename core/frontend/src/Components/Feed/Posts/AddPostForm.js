import { Card, Form, Button } from "react-bootstrap";
import { Component } from "react";
import { getUserId } from "../../../utils";

export default class AddPostForm extends Component {
    state = {
        //userId: 0,
        title: "",
        description: "",
    };

    handlePostTitleChange = (e) => {
        this.setState({
            title: e.target.value,
        });
    };

    handlePostDescChange = (e) => {
        this.setState({
            description: e.target.value,
        });
    };

    createPost = () => {

        if (this.state.title === "" || this.state.description === "") return

        const post = {
            username: localStorage.getItem('user_name'),
            userId: getUserId(),
            title: this.state.title,
            description: this.state.description,
        }
        this.props.addPost(post)
    }

    render() {
        return (
            <>
                <Card className='mb-4' style={{ width: "65%", margin: "auto"}}>
                    <Card.Header>Create a new post</Card.Header>
                    <Card.Body>
                        <Form className='mb-3'>
                            <Form.Group controlId="forTitle">
                                <Form.Control type="text" placeholder='Title' onChange={this.handlePostTitleChange} />
                            </Form.Group>
                        </Form>

                        <Form className='mb-3'>
                            <Form.Group controlId="forContent">
                                <Form.Control as="textarea" rows={3} placeholder='Content' onChange={this.handlePostDescChange} />
                            </Form.Group>
                        </Form>
                        <Button variant="secondary" className='container-fluid' onClick={this.createPost}>Publish</Button>
                    </Card.Body>
                </Card>
            </>
        )
    };

}