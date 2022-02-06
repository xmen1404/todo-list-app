import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import TodoList from '../pages/TodoList'
import Login from '../pages/Login'

const Root = () => {
  return (
      <Router>
        <Routes>
          <Route path="/todo-list" exact element={<TodoList/>}/>
          <Route path="/" exact element={<Login/>}/>
        </Routes>
      </Router>
  )
}

export default Root