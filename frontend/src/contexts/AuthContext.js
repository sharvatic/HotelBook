import React, { createContext, useState, useEffect } from 'react';
import { jwtDecode } from 'jwt-decode'; // Ensure the correct import for jwtDecode

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            try {
                const decoded = jwtDecode(token);
                setUser({ id: decoded.userID, role: decoded.role });
            } catch (error) {
                console.error('Failed to decode token:', error);
                // Optionally clear the token if it's invalid
                localStorage.removeItem('token');
            }
        }
    }, []);

    const login = (token) => {
        try {
            const decoded = jwtDecode(token);
            setUser({ id: decoded.userID, role: decoded.role });
            localStorage.setItem('token', token);
        } catch (error) {
            console.error('Failed to decode token during login:', error);
        }
    };

    const logout = () => {
        setUser(null);
        localStorage.removeItem('token');
    };

    return (
        <AuthContext.Provider value={{ user, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

export default AuthContext;
