import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import API from '../API.js'


const Login = ({setUser}) => {
    
    const [email, setEmail] = useState('test@test.com');
    const [password, setPassword] = useState('123');
    const navigate = useNavigate();

    const handleLogin = (data) => {
        if (!data.error) {
            localStorage.setItem('token', data.token)
            setUser(data)
            navigate('/todos')
        }
    }
    const handleClick = () => {
        API.login(email, password, handleLogin)
    }
    

    return (
        <div className="row">
            <div className="col-md-4 offset-md-4">
                <h1 className="mb-4 mt-4">Login</h1>
                <form>
                    <div className="form-group mb-4">
                        <label htmlFor="email">Email Address</label>
                        <input
                            onChange={e => setEmail(e.target.value)}
                            value={email}
                            type="email" 
                            className="form-control" 
                            id="email" 
                            aria-describedby="email" 
                            placeholder="Enter email"/>
                        <small id="email-help" className="form-text text-muted">Email is "test@test.com"</small>
                    </div>
                    <div className="form-group mb-4">
                        <label htmlFor="password">Password</label>
                        <input
                            onChange={e => setPassword(e.target.value)}
                            value={password}
                            type="password" 
                            className="form-control" 
                            id="password" 
                            aria-describedby="password-help" 
                            placeholder="Password"/>
                        <small id="password-help" className="form-text text-muted">Password is "123"</small>
                    </div>
                    <button type="button" className="btn btn-primary col-12" onClick={handleClick}>Login</button>
                </form>
            </div>
        </div>
    );
}

export {Login};