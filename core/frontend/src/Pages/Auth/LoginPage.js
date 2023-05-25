import React, {useState, useEffect} from 'react';
import {Container, Row, Col, Form, Button, Alert} from 'react-bootstrap';
import {Link, useNavigate} from 'react-router-dom';
import {loginUser} from "../../service/api";

const LoginPage = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [alert, setAlert] = useState('');
    const [showAlert, setShowAlert] = useState(false);

    const handleUsernameChange = (event) => {
        setUsername(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    const handleLoginClick = async () => {
        if (!username || !password) {
            setError('Fill in all the fields');
            setAlert('warning');
        } else if (username && password) {
            const response = await loginUser(username, password);

            if (!response) {
                setError('Unknown error');
                setAlert('primary');
                return;
            }

            const status = response.status;

            switch (status) {
                case 200:
                    setError('Login successfully');
                    setAlert('success');
                    const token = response.data['token'];
                    localStorage.setItem('jwt_token', token);
                    break;
                case 400:
                    setError('Invalid username or password');
                    setAlert('danger');
                    break;
                case 401:
                    setError('User is not found')
                    setAlert('danger')
                    break;
                case 500:
                    setError(`Internal Server Error:\n${response.data.message}`);
                    setAlert('danger');
                    break;
                default:
                    setError('Unknown error');
                    setAlert('danger');
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

    useEffect(() => {
    }, [username, password]);

    return (
        <Container className="d-flex justify-content-center" style={{height: '100vh'}}>
            <Row className="align-items-center">
                <Col>
                    <h3 className="mb-3 text-center">Sign In</h3>
                    <Form>
                        <Form.Group className="mb-3">
                            <Form.Control
                                type="text"
                                placeholder="Username"
                                value={username}
                                onChange={handleUsernameChange}
                                style={{fontSize: '1.2rem'}}
                            />
                        </Form.Group>
                        <Form.Group className="mb-3">
                            <Form.Control
                                type="password"
                                placeholder="Password"
                                value={password}
                                onChange={handlePasswordChange}
                                style={{fontSize: '1.2rem'}}
                            />
                        </Form.Group>
                        <Link to="/auth/register" className="d-flex justify-content-center">
                            Create an account
                        </Link>
                        {showAlert && (
                            <Alert variant={alert} key={alert} className="mb-3" style={{fontSize: '1.2rem'}}>
                                {error}
                            </Alert>
                        )}
                        <Button variant="primary" onClick={handleLoginClick} className="w-100"
                                style={{fontSize: '1.2rem'}}>
                            Login
                        </Button>
                    </Form>
                </Col>
            </Row>
        </Container>

    );
};

export default LoginPage;
