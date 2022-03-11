import React, { useState, useEffect } from 'react';
import { Routes, Route, BrowserRouter} from "react-router-dom";
import { Todo } from './Todo/Todo';
import { Login } from './Login/Login';
import { Home } from './Home/Home';
import API from './API.js'

import "bootswatch/dist/minty//bootstrap.min.css";
import './App.css';


const Nav = (props) => {
  return (
    <nav className="navbar navbar-light bg-light justify-content-between p-4">
      <h1 className="navbar-brand">Tasks Keeper</h1>
      <h4 className="navbar-brand">{`${props.user ? props.user.first_name : ''} ${props.user ? props.user.last_name : ''}`}</h4>
    </nav>
  );
}

const App = () => {

  const [user, setUser] = useState(null);
  const [checkedUser, setCheckedUser] = useState(false);

  const handleGetUser = (data) => {
    if (!data.erorr) {
      setUser(data);
    }
    setCheckedUser(true)
  }
  
  useEffect(() => {
    if (user) {
      return
    }
    if (API.getToken()) {
      API.getUser(handleGetUser);
    }
  })
  return (
    <div className="App">
      <Nav user={user}/>
      <div className="container">
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login setUser={setUser} />} />
            <Route path="/todos" element={<Todo user={user} checkedUser={checkedUser}/>} />
          </Routes>
        </BrowserRouter>
      </div>
    </div>
  );
}

export default App;
