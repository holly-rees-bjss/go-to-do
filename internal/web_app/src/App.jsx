import { useEffect, useState } from 'react'
import { v4 as uuidv4 } from 'uuid';

import { getTodos } from './api/apiCalls';

import './App.css'
import ToDoCard from './components/ToDoCard';
import PostTodo from './components/PostTodo';

function App() {
  const [count, setCount] = useState(0)

  const [todos, setTodos] = useState([])

  useEffect(() => {
    getTodos().then((todos) => {
      setTodos(todos)
    })
  }, [])


  return (
    <>
      <div>

      </div>
      <h1>To Do App</h1>
      <div className="card">
        <h2>Your Todos:</h2>
        <ol>
          {todos.map((todo, i) => (
            < li key={uuidv4()} >
              <ToDoCard todo={todo} index={i}></ToDoCard>
            </li>
          ))}
        </ol>

      </div >
      <PostTodo />
    </>
  )
}

export default App
