// src/components/Home.js

import React from 'react';
import '../styles/Home.css';
// import { GoogleAuthProvider, signInWithPopup } from "firebase/auth";
// import { auth } from '../firebase/firebaseConfig'
//import { useNavigate } from 'react-router-dom';

const Home = () => {

    //const navigate = useNavigate();

    // const handleGoogle = async () => {
    //     const provider = new GoogleAuthProvider();
    //     try {
    //         const result = await signInWithPopup(auth, provider);
    //         const token = await result.user.getIdToken() ;
    //         localStorage.setItem("token", token) ;

    //         console.log("User signed in with token: ",token) ;
    //         navigate("/userdashboard")

    //         console.log("User signed in:", result.user);
            
    //     } catch (error) {
    //         console.error("Error during sign-in:", error);
    //     }
    // };

    return (
        <div className="home">
            <h1>Welcome to BookMyHotel</h1>
            <p>Discover our services and manage your bookings effortlessly!</p>
            <div className="home-actions">
                <a href="/login" className="home-link">Login</a>
                <a href="/signup" className="home-link">Signup</a>
            </div>
            <div className="home-actions">
                <a href="/signin" className="home-link">Sign In with Google</a>
            </div>
        </div>
    );
};

export default Home;
