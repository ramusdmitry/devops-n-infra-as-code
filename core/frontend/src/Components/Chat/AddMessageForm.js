import { Component } from "react";
import { Form, Button } from "react-bootstrap";
import { getUsername } from "../../utils";

export default class AddMessageForm extends Component {

    state = {
        name: getUsername(),
        msg: "",
    };

    handleNameChange = (e) => {
        this.setState({
            name: e.target.value,
        });
    };

    handleMsgChange = (e) => {
        this.setState({
            msg: e.target.value,
        });
    };

    createMessage = () => {
        const item = {
            name: this.state.name,
            msg: this.state.msg,
        }

        this.props.addNewMsg(item)
        this.setState({
            msg: ''
        })
    }

    render() {
        return (
            <Form style={{ width: "65%", margin: "auto"}}>
                <Form.Group controlId="forMsg" className='my-3'>
                    <Form.Control placeholder='...' onChange={this.handleMsgChange} value={this.state.msg} />
                </Form.Group>
                <Button variant="secondary" className='container-fluid mt-3' onClick={this.createMessage}>Send</Button>
            </Form>
        );
    }


}