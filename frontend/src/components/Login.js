// src/components/Login.js

import React, { useState, useContext } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom'; // Import useNavigate for redirection
import '../styles/Auth.css';
import AuthContext from '../contexts/AuthContext';

const Login = () => {
    const { login } = useContext(AuthContext) || {};
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate(); // Initialize navigate for redirect

    const handleLogin = async (e) => {
        e.preventDefault();
        setError('');

        try {
            const response = await axios.post('http://localhost:8090/api/login', {
                username,
                password,
            });

            // Store token in localStorage
            const token = response.data.token;
            login(token);
            console.log(token) ;
            //localStorage.setItem('token', token);

            alert('Login successful!');
            navigate('/userdashboard');
        } catch (err) {
            setError(err.response?.data?.error || 'Login failed');
        }
    };

    return (
        <div className="container">
            <h1>Login</h1>
            {error && <div className="error">{error}</div>}
            <form onSubmit={handleLogin}>
                <div className="input-field">
                    <input
                        type="text"
                        placeholder="Username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required
                    />
                </div>
                <div className="input-field">
                    <input
                        type="password"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <button type="submit" className="button">Login</button>
            </form>
        </div>
    );
};

export default Login;
