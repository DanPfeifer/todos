import React, { useState } from 'react';
import moment from 'moment';
import API from '../API.js'
import { TodoCreate } from './TodoCreate';

const Status = (props) => {
    if (props.status) {

    }
}
const TodoRow = (props) => {
    const [edit, setEdit] = useState(false);

    const handleClickEdit = () => {
        setEdit(!edit)
    }
    const handleEditAfter = () => {
        setEdit(false)
        API.getTodos(props.handleGetTodos)
    }
    const handleClickPrioritize = (todo) => {
        todo.priority = todo.priority == 1 ? 0 : 1;
        API.updateTodo(todo,todo.id, handleEditAfter)
    }
    const handleClickUpdateStatus = (todo) => {
        todo.status = todo.status == 1 ? 0 : 1;
        API.updateTodo(todo, todo.id, handleEditAfter)
    }
    
    if(edit) {
        return (<TodoCreate todo={props.todo} handleGetTodos={handleEditAfter} handleClickEdit={handleClickEdit}/>)
    }
    
    return (
        <div className={`card text-center mb-4 ${props.todo.priority != 0 ? "border-primary" : ''}`}>
            <div className="card-header">
                <div className="d-flex justify-content-between">
                    <h6 className={`card-text m-1 align-self-center ${props.todo.priority != 0 ? "text-primary" : ''}`}>
                        Due: {moment(props.todo.due_date).format("dddd, MMMM Do YYYY, h:mm:ss a")}
                    </h6>
                    <button 
                        className="btn btn-outline-danger"
                        onClick={e =>{props.handleClickDelete(props.todo.id)}}>
                        Remove
                    </button>
                </div>                
            </div>
            <div className="card-body">
                <h5 className="card-title">{props.todo.title}</h5>
                <p className="card-text">{props.todo.message}</p>
                
            </div>
            <div className="card-footer text-muted">
            <div className="d-flex justify-content-between">
                    <button 
                        className="btn btn-outline-info"
                        onClick={handleClickEdit}>
                        Edit
                    </button>
                    <button 
                        className="btn btn-outline-primary"
                        onClick={e =>{handleClickPrioritize(props.todo)}}>
                        {props.todo.priority == 0 ? 'Prioritize': 'Remove Priority'}
                    </button>
                    <button 
                        className={`btn ${props.todo.status == 0 ? 'btn-success' : 'btn-outline-warning'}`}
                        onClick={e =>{handleClickUpdateStatus(props.todo)}}>
                        {props.todo.status == 0 ? 'Complete' : 'Redo'}
                    </button>
                </div>
            </div>
        </div>
    )
}

const TodoTable = (props) =>{


    const handelDeleteTodo = (data) => {
        if (!data.error) {
            API.getTodos(props.handleGetTodos)
        }
    }

    const handleClickDelete = (id) => {
        API.deleteTodo(id, handelDeleteTodo)
    }

    const todoRows = props.todos ? props.todos.map((t, i) => {return <TodoRow handleGetTodos={props.handleGetTodos} handleClickDelete={handleClickDelete} todo={t} key={i}/>;} ) : '';
    return (
        <div>
            {todoRows}
        </div>
    );
}
export {TodoTable}