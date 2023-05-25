import {Button, Nav, NavDropdown} from "react-bootstrap"

import React, {useState} from 'react'
import {useNavigate} from "react-router-dom";
import {getAuth} from "../utils";

const groups = ['Admin', 'Moderator', 'Journalist', 'Guest'];   


export default function Header(props) {

    const [currentTab, setCurrentTab] = useState(window.location.pathname);

    const nav = useNavigate();

    const [isLoggedIn, user] = getAuth();
    const auth = isLoggedIn;

    const handleSelect = (eventKey) => {
        switch (eventKey) {
            case "main":
                setCurrentTab("/");
                nav('/');
                break;
            case "Feed":
                setCurrentTab("/Feed")
                nav('/Feed');
                break;
            case "Dashboard":
                setCurrentTab("/Dashboard");
                nav('/Dashboard');
                break;
            default:
                break;
        }
    }

    const tabs = (auth) => {
        if (!auth) {
            return (
                <>
                    <Nav.Item className="mx-1">
                        <Nav.Link eventKey="/" href="/">
                            Main
                        </Nav.Link>
                    </Nav.Item>
                    <Nav.Item className="mx-1">
                        <Nav.Link eventKey="/feed" href="/feed">
                            Feed
                        </Nav.Link>
                    </Nav.Item>
                    <Button variant="dark" className="mx-2">Guest</Button>
                    <Nav.Item className="mx-2">
                        <Button variant="outline-primary" onClick={() => nav('/Auth/login')}>Login</Button>
                    </Nav.Item>
                    <Nav.Item className="mx-2">
                        <Button variant="outline-secondary" onClick={() => nav('/Auth/register')}>Sign up</Button>
                    </Nav.Item>
                </>
            );
        }

        if (auth && user.group_id === 1) {
            return (
                <>
                    <Nav.Item className="mx-1">
                        <Nav.Link eventKey="/" href="/">
                            Main
                        </Nav.Link>
                    </Nav.Item>
                    <Nav.Item className="mx-1">
                        <Nav.Link eventKey="/feed" href="/feed">
                            Feed
                        </Nav.Link>
                    </Nav.Item>
                    <Nav.Item className="mx-1">
                        <Nav.Link eventKey="/dashboard" href="/dashboard">
                            Dashboard
                        </Nav.Link>
                    </Nav.Item>
                    {dropdown(auth, user)}
                </>
            );
        }

        return (
            <>
                <Nav.Item className="mx-1">
                    <Nav.Link eventKey="/" href="/">
                        Main
                    </Nav.Link>
                </Nav.Item>
                <Nav.Item className="mx-1">
                    <Nav.Link eventKey="/feed" href="/feed">
                        Feed
                    </Nav.Link>
                </Nav.Item>
                {dropdown(user)}
            </>
        );
    }

    const dropdown = () => {
        return (
            <NavDropdown title={`@${user.user_name} (${groups[user.group_id - 1]})`} id="nav-dropdown" className="mx-1">
                <p className="text-left px-3">Signed as <strong>{groups[user.group_id - 1]}</strong></p>
                <NavDropdown.Divider />
                <NavDropdown.Item eventKey="logout" onClick={() => nav('/logout')}>Log out</NavDropdown.Item>
            </NavDropdown>
        )
    }





    return (
        <div className='my-4'>
            <Nav variant="pills" className="justify-content-center" lg activeKey={currentTab} onSelect={handleSelect}>
                {tabs(auth)}
            </Nav>
        </div>
    )
}