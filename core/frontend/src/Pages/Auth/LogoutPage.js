import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const LogoutPage = () => {
    const history = useNavigate();

    useEffect(() => {
        localStorage.removeItem('jwt_token');
        const redirectTimeout = setTimeout(() => {
            history('/Auth/login');
        }, 1000);
        return () => clearTimeout(redirectTimeout);
    }, [history]);

    return (
        <div className="align-content-center">
            <h1>Logging out...</h1>
        </div>
    );
};


export default LogoutPage;