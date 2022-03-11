class API {

    /**
     * 
     * @returns string
     */
    getToken() {
        return localStorage.getItem('token');
    }
    /**
     * 
     * @param {string} route 
     * @param {object} data 
     * @param {string} method 
     * @returns Response Object
     */
    async sendRequest(route, data, method, callback) {
        let token = this.getToken()
        let config ={
            headers: {
                "content-type":"application/json; charset=UTF-8",
                "X-Session-Token": token
            },
            method: method
        }
        if (["POST", "PUT"].includes(method)) {
            config.body = JSON.stringify(data);
        }
        
        const response = await fetch(`http://localhost:88${route}`, config)
        if (response.ok) {
            const json = await response.json();
            callback(json);
            return
        }
        const text = await response.text();
        callback({error: text});     
    }
    /**
     * 
     * @param {string} email 
     * @param {string} password 
     * @returns json
     */
    login(email, password, callback) {
        this.sendRequest('/login', {email: email, password: password}, 'POST',  callback)  
    }
    /**
     * 
     * @returns Response Object
     */
    getUser(callback) {
        this.sendRequest('/user', {}, 'GET',  callback) 
    }
    /**
     * 
     * @returns Response Object
     */
    getTodos(callback) {
        this.sendRequest('/todo', {}, 'GET',  callback)   
    }

    /**
     * 
     * @param {string} id 
     * @returns Response Object
     */
    getTodo(id, callback) {
        this.sendRequest(`/todo/${id}`, {}, 'GET',  callback)   
    }
    /**
     * 
     * @param {object} todo 
     * @returns Response Object
     */
    createTodo(todo, callback) {
        this.sendRequest('/todo', todo, 'POST',  callback)   
    }
    /**
     * 
     * @param {object} todo 
     * @returns Response Object
     */
    updateTodo(todo, id, callback) {
        this.sendRequest(`/todo/${id}`, todo, 'PUT',  callback)   
    }
    /**
     * 
     * @param {string} id 
     * @returns Response Object
     */
    deleteTodo(id, callback) {
        this.sendRequest(`/todo/${id}`, {}, 'DELETE',  callback)   
    }
}

export default new API();