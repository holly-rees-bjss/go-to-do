import axios from "axios";

function getTodos() {
    return axios
        .get(`http://localhost:8080/api/todos`)
        .then(({ data }) => {
            console.log("todos>>> ", data)
            return data;

        })
        .catch((err) => { });
}

function postTodo(todo) {
    return axios
        .post(
            `http://localhost:8080/api/todo`,
            todo
        )
        .then(({ data }) => {
            const { todo } = data;
            return todo;
        });
}

function deleteTodo(index) {
    return axios
        .delete(`http://localhost:8080/api/todo/${index}`);
}


function patchTodo(index, status) {
    const patch = { Task: "", Status: status }
    return axios
        .patch(
            `http://localhost:8080/api/todo/${index}`,
            patch
        )
        .then(({ data }) => {
            const { todo } = data;
            return todo;
        });
}


export { getTodos, postTodo, deleteTodo, patchTodo };