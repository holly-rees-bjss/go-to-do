import { useContext, useEffect, useState } from "react";

import { deleteTodo, patchTodo } from "../api/apiCalls";


import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";

function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-GB', {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric',
    });
}

function ToDoCard({ todo, index }) {


    console.log(todo, "<<<<<")

    const [feedbackMsg, setFeedbackMsg] = useState("");

    function handleDelete(event) {
        const i = parseInt(event.target.value, 10) + 1;
        console.log(i, "<< todo index")
        deleteTodo(i)
            .then(() => {
                setFeedbackMsg("deleted!");

            })
            .catch((err) => {
                console.log(err);
                setFeedbackMsg(
                    "Oops, something went wrong! Couldn't delete your todo."
                );
            });
    }

    function handleMarkComplete(event) {
        const i = parseInt(event.target.value, 10) + 1;
        console.log(i, "<< todo index")
        patchTodo(i, "Completed")
            .then(() => {
                setFeedbackMsg("marked complete!");

            })
            .catch((err) => {
                console.log(err);
                setFeedbackMsg(
                    "Oops, something went wrong! Couldn't update your todo."
                );
            });
    }

    const isDueDateSet = todo.DueDate !== '0001-01-01T00:00:00Z';
    const formattedDate = isDueDateSet ? formatDate(todo.DueDate) : 'Not Set';

    return (
        <Card variant="outlined" style={{
            margin: "2vw", paddingLeft: "2em", paddingRight: "2em", border: "4px solid grey",
            borderRadius: "10px",
        }}>
            <h3>{todo.Task}</h3>
            <p>Status: {todo.Status}</p>
            <p>Due: {formattedDate}{" "}</p>
            <p>{feedbackMsg}</p>
            <div>
                <button
                    value={index}
                    style={{
                        backgroundColor: "#e4e4e4",
                        color: "black",
                        margin: "1vw",
                        padding: "1vw",
                        alignContent: "right",
                    }}
                    onClick={handleMarkComplete}
                >
                    mark complete
                </button>
                <button
                    value={index}
                    style={{
                        backgroundColor: "#e4e4e4",
                        color: "black",
                        margin: "1vw",
                        padding: "1vw",
                        alignContent: "right",
                    }}
                    onClick={handleDelete}
                >
                    delete
                </button>
            </div>
        </Card>
    );
}

export default ToDoCard;