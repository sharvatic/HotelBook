// src/components/Signup.js

import React, { useState } from 'react';
//import axios from 'axios'
import { useNavigate } from 'react-router-dom';
import { GoogleAuthProvider, signInWithPopup } from "firebase/auth";
import { auth } from '../firebase/firebaseConfig'
import '../styles/Auth.css';

const Signin = () => {
    const [role, setRole] = useState('')
    const [error, setError] = useState('');
    const navigate = useNavigate(); // Initialize navigate for redirect

    const handleSignin = async (e) => {
        e.preventDefault();
        setError('');

        const provider = new GoogleAuthProvider();
        try {
            const result = await signInWithPopup(auth, provider);
            const token = await result.user.getIdToken() ;
            localStorage.setItem("token", token) ;

            console.log("User signed in with token: ",token) ;

            // const response = await axios.post("http://localhost:8090/api/setClaims", {
            //     role
            // }, {
            //     headers: {
            //         Authorization: `Bearer ${token}`,
            //         'Content-Type': 'application/json'
            //     },
            // })
            // if (response.status === 201) {
            //     alert('Signup successful!');
            //     navigate("/userdashboard");
            //     console.log("User signed in:", result.user);

            // } else {
            //     setError('Signup failed');
            // }

            navigate("/userdashboard");
            console.log("User signed in:", result.user);

        } catch (error) {
            console.error("Error during sign-in:", error);
        }

    };

    return (
        <div className="container">
            <h1>Signin</h1>
            {error && <div className="error">{error}</div>}
            <form onSubmit={handleSignin}>
                <div className="input-field">
                    <input
                        type="text"
                        placeholder="Role(user / staff)"
                        value={role}
                        onChange={(e) => setRole(e.target.value)}
                        required
                    />
                </div>
                
                <button type="submit" className="button">Signin</button>
            </form>
        </div>
    );
};

export default Signin;
