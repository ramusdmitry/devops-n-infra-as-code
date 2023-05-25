import React, {useState} from 'react';
import {Col, Container, Form, Row, Button, Alert} from 'react-bootstrap';
import {useNavigate, Link} from 'react-router-dom';
import {registerUser} from "../../service/api";

const RegisterPage = () => {
    const navigate = useNavigate();
    const [name, setName] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorState, setErrorState] = useState(0);
    const [error, setError] = useState('');
    const [alert, setAlert] = useState('');
    const [showAlert, setShowAlert] = useState(false);

    const handleNameChange = (event) => {
        setName(event.target.value);
    };

    const handleUsernameChange = (event) => {
        setUsername(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    const handleRegisterClick = async () => {
        if (!name || !username || !password) {
            setError('Fill in all the fields');
            setAlert('warning');
        } else if (name && username && password) {
            const response = await registerUser(name, username, password);

            if (!response) {
                setError('Unknown error');
                setAlert('primary');
                return;
            }

            const status = response.status;

            switch (status) {
                case 200:
                    setError("Register successfully");
                    setAlert('success');
                    const token = response.data['token'];
                    localStorage.setItem('jwt_token', token);
                    navigate('/');
                    break;
                case 400:
                    setError("Invalid credentials");
                    setAlert('danger');
                    break;
                case 500:
                    setError(`Internal Server Error:\n${response.data.message}`)
                    setAlert('danger');
                    break;
                default:
                    setError("Unknown error");
                    setAlert("danger");
                    break;
            }
            if (status === 200) {
                setShowAlert(true);
                setTimeout(() => setShowAlert(false), 3000);
                navigate('/');
            }
        }

        setShowAlert(true);
        setTimeout(() => setShowAlert(false), 3000);

    };

    return (
        <Container className="d-flex justify-content-center">
            <Row className="py-5">
                <h1 className="d-flex justify-content-center mb-5">Create an account</h1>
                <Col>
                    <Form className="mb-3">
                        <Form.Group controlId="forName">
                            <Form.Control
                                type="text"
                                placeholder="Name"
                                value={name}
                                onChange={handleNameChange}
                                required
                            />
                        </Form.Group>
                    </Form>

                    <Form className="mb-3">
                        <Form.Group controlId="forUserName">
                            <Form.Control
                                type="text"
                                placeholder="@username"
                                value={username}
                                onChange={handleUsernameChange}
                                required
                            />
                        </Form.Group>
                    </Form>

                    <Form className="mb-3">
                        <Form.Group controlId="forPassword">
                            <Form.Control
                                type="password"
                                placeholder="******"
                                value={password}
                                onChange={handlePasswordChange}
                            />
                        </Form.Group>
                    </Form>

                    <Link to="/auth/login" className="d-flex justify-content-center">
                        Log in
                    </Link>

                    <Button
                        variant="secondary"
                        className="container-fluid mt-3"
                        onClick={handleRegisterClick}
                    >
                        Sign Up
                    </Button>
                    {showAlert && (
                        <Alert variant={alert} key={alert} className="mb-3" style={{fontSize: '1.2rem'}}>
                            {error}
                        </Alert>
                    )}
                </Col>
            </Row>
        </Container>
    );
};

export default RegisterPage;