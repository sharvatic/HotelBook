// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthContext';
import Home from './components/Home';
import Signup from './components/Signup';
import Signin from './components/Signin'
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import UserDashboard from './components/UserDashboard';

const App = () => {
    return (
        <AuthProvider>
          <Router>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/signup" element={<Signup />} />
                <Route path="/login" element={<Login />} />
                <Route path="/signin" element={<Signin/>} />
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/userdashboard/*" element={<UserDashboard />} />
            </Routes>
        </Router>
        </AuthProvider>
    );
};

export default App;
