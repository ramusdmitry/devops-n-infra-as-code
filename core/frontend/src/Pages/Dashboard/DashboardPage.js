import React, {Component} from 'react';
import {Container, Row, Table, Col, Button, Alert} from 'react-bootstrap';
import TableRow from '../../Components/Dashboard/TableRow';
import {applyUsers, deleteUsers, getUsers} from "../../service/api";

export default class Dashboard extends Component {
    constructor(props) {
        super(props);

        this.state = {
            users: [],
            modifiedUsers: new Map(),
            error: '',
            alert: '',
            needUpdate: true,
            showAlert: false,
            user_id: null,
            group_id: null
        };
    }

    fetchUsers = async () => {
        const response = await getUsers();

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
                const users = response.data.users;
                this.setState({
                    alert: 'primary',
                    error: 'Dashboard was updated',
                    users: users
                });
                break;
            case 401:
                this.setState({
                    alert: 'danger',
                    error: 'Unauthorized'
                });
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
    }

    handleDeleteAllBtn = async () => {
        const response = await deleteUsers();

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
                this.setState({
                    alert: 'success',
                    error: 'Everything users has been deleted.',
                    needUpdate: false
                });
                window.location.href = '/logout';
                break;
            case 500:
                this.setState({
                    alert: 'error',
                    error: `Internal Server Error\n${response.message}`
                });
                break;
            default:
                this.setState({
                    alert: 'error',
                    error: `Unknown Error\n${response.message}`
                });
                break;
        }
        this.setState({showAlert: true});
        setTimeout(() => this.setState({showAlert: false}), 3000);
    }

    handleSelectedGroup = (e) => {
        this.state.modifiedUsers.set(e.id, e)
    }

    handleUpdateAllBtn = async () => {
        await this.fetchUsers();
    }

    handleApplyAllBtn = async () => {

        if (this.state.modifiedUsers.size === 0) return

        const updatedUsers = Array.from(this.state.modifiedUsers.values());
        const response = await applyUsers(updatedUsers);

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
                this.setState({
                    alert: 'success',
                    error: 'Everything has been applied.',
                    needUpdate: false
                });
                break;
            case 500:
                this.setState({
                    alert: 'error',
                    error: `Internal Server Error\n${response.message}`
                });
                break;
            default:
                this.setState({
                    alert: 'error',
                    error: `Unknown Error\n${response.message}`
                });
                break;
        }
        this.setState({ showAlert: true });
        setTimeout(() => this.setState({ showAlert: false }), 3000);
    }

    async componentDidMount() {
        await this.fetchUsers();
    }
    render() {

        const arrUsers = this.state.users.map((i) => {
            return (
                <TableRow user={i} groupChanger={this.handleSelectedGroup} />
            )
        })

        const { showAlert,  alert, error } = this.state;

        return (
            <Container>
                <Row>
                    <h1>Dashboard</h1>
                    {showAlert && (
                        <Alert variant={alert} key={alert} className="mb-3" style={{fontSize: '1.2rem'}}>
                            {error}
                        </Alert>
                    )}
                </Row>
                <Row className='d-flex justify-content-end'>
                    <div className='d-flex justify-content-end'>
                        <Button type="submit" variant="danger" className='mx-1' onClick={this.handleDeleteAllBtn}>Delete All</Button>
                        <Button type="submit" variant="primary" className='mx-1' onClick={this.handleUpdateAllBtn}>Update All</Button>
                        <Button type="submit" variant="success" className='mx-1' onClick={this.handleApplyAllBtn}>Apply All</Button>
                    </div>
                </Row>
                <Row>
                    <Col>
                        <Table>
                            <thead>
                                <tr>
                                    <th>ID</th>
                                    <th>Name</th>
                                    <th>Username</th>
                                    <th>Group</th>
                                </tr>
                            </thead><tbody>{arrUsers}</tbody>
                        </Table>
                    </Col>
                </Row>
            </Container >
        );
    }
}