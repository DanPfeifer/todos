import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { TodoTable } from './TodoTable';
import { TodoCreate } from './TodoCreate';
import API from '../API.js'

const Todo = (props) =>{
    const navigate = useNavigate();
    useEffect(() => {
        if (!props.user && props.checkedUser) {
            navigate(`/login`);
          return
        }
    })

    const [todosToBeDone, setTodosToBeDone] = useState([]);
    const [todosCompleted, setTodosCompleted] = useState([]);
    const [todoCheck, setTodoCheck] = useState(false);
    

    const handleGetTodos = (data) => {
        setTodoCheck(true)
      if (!data.erorr) {
        let todos = data.todos.filter(t => {if (t.status != 1) {return t}}).sort((a, b) => {return b.priority - a.priority})
        let toDons = data.todos.filter(t => {if (t.status == 1) {return t}}).sort((a, b) => {return b.priority - a.priority})
        setTodosToBeDone(todos)
        setTodosCompleted(toDons)
      }
    }
    
    useEffect(() => {
      if (todoCheck) {
        return
      }
      API.getTodos(handleGetTodos);
    })

    return(
        <div>
            <div className="row">
                <div className="col-xs-12 title m-4">
                <h1>Your Task Board</h1>
                </div>
            </div>
            <div className="row">
                <div className="col-md-4 title mb-4"><h2>Tasks</h2></div>
                <div className="col-md-4 title mb-4"><h2>Completed</h2></div>
                <div className="col-md-4 title mb-4"><h2>Create New</h2></div>
            </div>
            <div className="row">
                <div className="col-md-4">
                    <TodoTable todos={todosToBeDone} handleGetTodos={handleGetTodos}/>
                </div>
                <div className="col-md-4">
                    <TodoTable todos={todosCompleted} handleGetTodos={handleGetTodos}/>
                </div>
                <div className="col-md-4">
                    <TodoCreate handleGetTodos={handleGetTodos}/>    
                </div>
            </div>
            
        </div>
    )
}

export {Todo};