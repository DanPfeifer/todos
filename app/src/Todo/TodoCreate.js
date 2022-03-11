import React, { useEffect, useState } from 'react';
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import moment from 'moment';
import API from '../API.js'

const labelStyles = {"fontWeight": "bold"};

const TodoCreate = (props) => {

    const [title, setTitle] = useState('');
    const [message, setMessage] = useState('');
    const [dueDate, setDueDate] = useState(new Date());
    
    useEffect(() => {
        if (props.todo) {
            setTitle(props.todo.title)
            setMessage(props.todo.message)
            setDueDate(new Date(props.todo.due_date))
        }
    }, [])
    

    const handleUpdateAfter = (data) => {
        if(!data.error) {
            //close modal and reload todo list or push
            API.getTodos(props.handleGetTodos)
        }
    }
    const handleClick = () => {
        if (props.todo) {
            API.updateTodo({
                title: title, 
                message: message, 
                priority: props.todo.priority, 
                status: props.todo.status,
                due_date: moment(props.todo.due_date).format('YYYY-MM-DD H:mm:ss')
            }, props.todo.id, handleUpdateAfter)
            return    
        }
        API.createTodo({
            title: title, 
            message: message, 
            priority: 0, 
            status: 0, 
            due_date: moment(dueDate).format('YYYY-MM-DD H:mm:ss')
        }, handleUpdateAfter)
    }

    return (
        <div className="card mb-4">
            <div className="card-body">
                <form >
                    <div className="mb-3">
                        <label htmlFor="title" className="form-label" style={labelStyles}>Title</label>
                        <input
                            onChange={e => setTitle(e.target.value)}
                            value={title}
                            type="text" 
                            className="form-control"
                            id="title"
                            placeholder="Title"/>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="message" className="form-label" style={labelStyles}>Message</label>
                        <textarea 
                            onChange={e => setMessage(e.target.value)}
                            value={message}
                            type="text" 
                            className="form-control" 
                            id="message" 
                            placeholder="..."
                            rows="3"></textarea>
                    </div>            
                    <div className="mb-3">
                        <label htmlFor="due-date" className="form-label" style={labelStyles}>Due Date</label>
                        <DatePicker className="form-control" selected={dueDate} onChange={(date) => setDueDate(date)} />
                    </div>
                    <div className="d-flex justify-content-between">
                        <button type="button" className="btn btn-primary" onClick={handleClick}>Save</button>
                        <button type="button" className={`btn btn-outline-info ${props.todo ? "show" : "d-none"}`} onClick={props.handleClickEdit}>Close</button>
                    </div>
                </form>
            </div>
        </div>
    )
}


export {TodoCreate};