// src/components/UserDashboard.js
import React, { useEffect, useState, useContext } from 'react';
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom';
import axios from 'axios';
import Sidebar from './Sidebar';
import '../styles/Dashboard.css';
import AuthContext from '../contexts/AuthContext';

const UserDashboard = () => {
    const { user, logout } = useContext(AuthContext) || {};
    const [tables, setTables] = useState([]);
    const [menus, setMenus] = useState([]);
    const [menuItems, setMenuItems] = useState([]);
    const [orderArray, setOrderArray] = useState([]);
    const [orders, setOrders] = useState([]);
    const [error, setError] = useState('');

    const navigate = useNavigate();

    useEffect(() => {
        if (!user) {
            navigate('/signin'); // Redirect to login if not authenticated
        } else {
            fetchTables();
    
            console.log("User state:", user); // Log user state
            const token = localStorage.getItem('token');
            console.log("Token:", token); // Log the token
    
            if (window.location.pathname.includes('/userdashboard/menus')) {
                fetchMenus();
            }
            if (window.location.pathname.includes('/userdashboard/orderview')) {
                fetchOrders();
            }
        }
    }, [user, navigate]);
    
    const fetchTables = async () => {
        try {
            const response = await axios.get('http://localhost:8090/api/tables');
            setTables(response.data.tables);
        } catch (err) {
            console.error('Error fetching tables:', err);
        }
    };

    const fetchMenus = async () => {
        try {
            const response = await axios.get('http://localhost:8090/api/menu');
            setMenus(response.data);
        } catch (err) {
            console.error('Error fetching menus:', err);
        }
    };

    const fetchMenuItems = async (menuId) => {
        try {
            const response = await axios.get(`http://localhost:8090/api/menu/${menuId}`);
            setMenuItems(response.data);
        } catch (err) {
            console.error('Error fetching menu items:', err);
        }
    };

    const createOrder = (menuID, name) => {
        const newOrder = {
            menu_item_id: menuID,
            item_name: name,
            quantity: 1
        };

        // Add the new order to the existing order array
        setOrderArray((prevOrderArray) => [...prevOrderArray, newOrder]);
        console.log('Order created:', newOrder);
    };


    const placeOrder = async () => {
        try {
            const token = localStorage.getItem('token');

            // Ensure orderArray is defined and accessible here
            if (!orderArray || orderArray.length === 0) {
                alert("No items in the order to place.");
                return;
            }

            // Send the entire orderArray as the payload
            const response = await axios.post(
                'http://localhost:8090/api/orders',
                orderArray,
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                        'Content-Type': 'application/json', // Adding Content-Type header for JSON
                    },
                }
            );

            alert('Order placed successfully!');
            console.log(response.data);
        } catch (err) {
            alert('Failed to place order')
        }
    };
    const clearOrder = async () => {
        setOrderArray([]) ;
    }

    const fetchOrders = async () => {
        const token = localStorage.getItem('token');
        try {
            const response = await axios.get('http://localhost:8090/api/orders', {
                headers: { Authorization: `Bearer ${token}` },
            });
            console.log("Fetched orders:", response.data);
            setOrders(response.data.orders);
        } catch (error) {
            console.error('Error fetching orders:', error.response ? error.response.data : error.message);
            
            if (error.response) {
                console.error('Response status:', error.response.status);
                console.error('Response headers:', error.response.headers);
            }
        }
    };
    

    const bookTable = async (tableId) => {
        try {
            const token = localStorage.getItem('token');
            await axios.post(`http://localhost:8090/api/tables/${tableId}/book`, {}, {
                headers: { Authorization: `Bearer ${token}` },
            });
            alert('Table booked successfully!');
            fetchTables(); // Refresh tables
        } catch (err) {
            setError(err.response?.data?.error || 'Could not book table');
        }
    };

    const cancelTable = async (tableId) => {
        try {
            const token = localStorage.getItem('token');
            await axios.post(`http://localhost:8090/api/tables/${tableId}/cancel`, {}, {
                headers: { Authorization: `Bearer ${token}` },
            });
            alert('Table cancelled successfully!');
            fetchTables(); // Refresh tables
        } catch (err) {
            setError(err.response?.data?.error || 'Could not cancel table');
        }
    };

    const handleLogout = () => {
        console.log("Current user :", user)
        logout();
        //localStorage.removeItem('token'); // Remove the token from local storage

        navigate('/'); // Redirect to login page
    };

    return (
        <div className="dashboard">
            <Sidebar />
            <div className="content">

                <Routes>
                    <Route path="/" element={<Navigate to="tables" />} /> {/* Default route to tables */}
                    <Route
                        path="tables"
                        element={
                            <div>
                                <h1>Manage Tables</h1>
                                {error && <p className="error-message">{error}</p>}

                                <div className="tables-list">
                                    {tables.map((table) => (
                                        <div key={table.id} className="table-item">
                                            <p>Table Number: {table.table_number}</p>
                                            <p>Seats: {table.seats}</p>
                                            <p>Status: {table.is_booked ? 'Booked' : 'Available'}</p>

                                            {table.is_booked ? (
                                                table.booked_by === user.id ? (
                                                    <button onClick={() => cancelTable(table.id)} className="button cancel-button">
                                                        Cancel Booking
                                                    </button>
                                                ) : (
                                                    <button className="button already-booked-button" disabled>
                                                        Already Booked
                                                    </button>
                                                )
                                            ) : (
                                                <button onClick={() => bookTable(table.id)} className="button book-button">
                                                    Book Table
                                                </button>
                                            )}
                                        </div>
                                    ))}
                                </div>
                            </div>

                        }
                    />
                    <Route
                        path="menus"
                        element={
                            <div>
                                <h1>Menu</h1>

                                {/* Menus Section */}
                                <div className="menu-items">
                                    {menus.map((menu) => (
                                        <button
                                            key={menu.id}
                                            className="table-item"
                                            onClick={() => fetchMenuItems(menu.id)}
                                        >
                                            {menu.name}
                                        </button>
                                    ))}
                                </div>

                                {/* Menu Items Section */}
                                <div className="menu-item-details">
                                    {menuItems.length > 0 ? (
                                        menuItems.map((item) => (
                                            <div key={item.id} className="menu-item-detail">
                                                <div className="menu-item-content">
                                                    <div>
                                                        <p>{item.name}</p>
                                                        <p>Price: {item.price}</p>
                                                    </div>
                                                    <button className="order-button" onClick={() => createOrder(item.id, item.name)}>Order</button>
                                                </div>
                                            </div>
                                        ))
                                    ) : (
                                        <p>No menu items available. Select a menu to view items.</p>
                                    )}
                                </div>
                            </div>

                        }
                    />
                    <Route
                        path="orders"
                        element={
                            <div className="orders-container">
                                <h2>Your Order</h2>
                                <div className="orders-list">
                                    {orderArray.length > 0 ? (
                                        <ul className="order-items">
                                            {orderArray.map((order, index) => (
                                                <li key={index}>
                                                    <span className="order-item-name">Item Name: {order.item_name}</span>
                                                </li>
                                            ))}
                                        </ul>
                                    ) : (
                                        <p className="no-orders-message">No items to display </p>
                                    )}
                                </div>
                                <div>
                                    <button className='button place-order-btn' onClick={() => placeOrder()}>
                                        Place Your Order
                                    </button>
                                    <button className='button clear-order-btn' onClick={() => clearOrder()}>
                                        Clear Your Order
                                    </button>
                                </div>
                            </div>
                        }
                    />
                    <Route
                        path="orderview"
                        element={
                            <div className="orders-container">
                                {console.log("Orders array:", orders)} {/* Debugging the orders array */}
                                {orders.length > 0 ? (
                                    orders.map((order, index) => (
                                        <div key={index} className="order-box">
                                            <h3>Order ID: {order.order.ID}</h3>
                                            <p>Status: {order.order.Status}</p>
                                            <p>Created At: {new Date(order.order.CreatedAt).toLocaleString()}</p>
                                            <div className="order-items">
                                                {order.items.length > 0 ? (
                                                    order.items.map((item, idx) => (
                                                        <div key={idx} className="order-item">
                                                            <p>Menu Item ID: {item.menu_item_id}</p>
                                                            <p>Quantity: {item.quantity}</p>
                                                        </div>
                                                    ))
                                                ) : (
                                                    <p>No items in this order</p>
                                                )}
                                            </div>
                                        </div>
                                    ))
                                ) : (
                                    <p>No orders to display</p>
                                )}
                            </div>
                        }
                    />
                </Routes>
                <button onClick={handleLogout} className="button logout-button">
                    Logout
                </button>
            </div>
        </div>
    );
};

export default UserDashboard;
