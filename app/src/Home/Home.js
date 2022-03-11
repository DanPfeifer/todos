import React from 'react';
import { Link } from "react-router-dom";

const Home = () => {
    return(
        <div className="row">
            <div className="col-md-4 offset-md-4 text-center">
                <h1 className="m-4">Welcome to Task Kepper the keeper of tasks.</h1>
                <Link to="login"><h3 className="m-4">Login</h3></Link>
            </div>
        </div>
    )
}

export {Home}