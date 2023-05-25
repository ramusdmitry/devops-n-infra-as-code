
import { Form } from 'react-bootstrap'
import React, { Component } from 'react'

const groups = ['Admin', 'Moderator', 'Journalist', 'Guest'];   

export default class TableRow extends Component {

    state = {
        user: this.props.user,
        groupValue: this.props.user.group_id - 1
        //groupValue: groups[this.props.user.group_id - 1]
    }

    handleChangeGroup = (e) => {
        this.setState({ groupValue: e.target.value });

        const item = {
            id: this.state.user.id, // это мы всё равно не будем менять (тк зависит от БД)
            name: this.state.user.name,
            username: this.state.user.username,
            group_id: Number(e.target.value) + 1, // поменяли пока только группу
        };
        this.props.groupChanger(item);
    }

    render() {
        const { user } = this.props;

        const options = groups.map((group, i) => {
            return <option value={i}>{group}</option>
        });

        return (
            <tr>
            <th>{user.id}</th>
            <th>{user.name}</th>
            <th>{user.username}</th>
            <th>
                <Form.Select value={this.state.groupValue} onChange={this.handleChangeGroup}>
                {options}
                </Form.Select>
            </th>
        </tr>
        )
    }
}