import { useState, useContext } from "react";

import { postTodo } from "../api/apiCalls";

function PostTodo({ handleRefresh }) {

    const [userInput, setUserInput] = useState({ Task: "", Status: "Not Started", DueDate: "" });
    const [feedbackMsg, setFeedbackMsg] = useState("");

    function handlePostTodo(event) {
        event.preventDefault();
        const regex = /[a-z]/i;

        postTodo({
            task: userInput.Task,
            status: userInput.Status,
            DueDate: userInput.DueDate ? new Date(userInput.DueDate).toISOString() : null
        })
            .then((res) => {
                setFeedbackMsg("Todo posted!");
                handleRefresh()
            })
            .catch((err) => {
                console.log(err);
                setFeedbackMsg("Oops! Couldn't post your todo. Try again later!");
            });
        setUserInput({ Task: "", Status: "Not Started", DueDate: "" });

    }

    function handleInput(event) {
        const { name, value } = event.target;
        setUserInput(prevState => ({
            ...prevState,
            [name]: value
        }));
    }

    return (
        <form onSubmit={handlePostTodo}>
            <p>{feedbackMsg}</p>
            <label htmlFor="task">Task:</label>
            <div>
                <input
                    style={{
                        width: "80vw",
                        height: "50px",
                        textAlign: "center",
                        border: "1px solid black",
                        background: "white",
                        color: "black",
                    }}
                    onChange={handleInput}
                    value={userInput.Task}
                    name="Task"
                    id="task"
                    placeholder="What do you need to do?"
                    type="text"
                    required
                />
            </div>
            <label htmlFor="dueDate">Due Date:</label>
            <div>
                <input
                    style={{
                        width: "80vw",
                        height: "50px",
                        textAlign: "center",
                        border: "1px solid black",
                        background: "white",
                        color: "black",
                    }}
                    onChange={handleInput}
                    value={userInput.DueDate}
                    name="DueDate"
                    id="dueDate"
                    placeholder="Due Date (dd-mm-yyyy)"
                    type="date"
                />
            </div>
            <label htmlFor="status">Status:</label>
            <div>
                <select
                    style={{
                        width: "80vw",
                        height: "50px",
                        textAlign: "center",
                        border: "1px solid black",
                        background: "white",
                        color: "black",
                    }}
                    onChange={handleInput}
                    value={userInput.Status}
                    name="Status"
                    id="status"
                >
                    <option value="Not Started">Not Started</option>
                    <option value="In Progress">In Progress</option>
                    <option value="Completed">Completed</option>
                </select>
            </div>
            <button
                style={{ backgroundColor: "#eea0a2", margin: "0.5rem", color: "black" }}
            >
                Add
            </button>
        </form>
    );
}


export default PostTodo;