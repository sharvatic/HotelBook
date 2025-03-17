// src/components/Signup.js

import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import '../styles/Auth.css';

const Signup = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('')
    const [error, setError] = useState('');
    const navigate = useNavigate(); // Initialize navigate for redirect

    const handleSignup = async (e) => {
        e.preventDefault();
        setError('');

        try {
            const response = await axios.post('http://localhost:8090/api/signup', {
                username,
                password,
                role,
            });

            if (response.status === 201) {
                alert('Signup successful!');
                navigate('/login');

            } else {
                setError('Signup failed');
            }

        } catch (err) {
            setError(err.response.data.error || 'Signup failed');
        }
    };

    return (
        <div className="container">
            <h1>Signup</h1>
            {error && <div className="error">{error}</div>}
            <form onSubmit={handleSignup}>
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
                        type="text"
                        placeholder="Role(user / staff)"
                        value={role}
                        onChange={(e) => setRole(e.target.value)}
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
                <button type="submit" className="button">Signup</button>
            </form>
        </div>
    );
};

export default Signup;
