import { BrowserRouter as Router, Route } from "react-router-dom";
import TodoList from '../pages/TodoList'

const Root = () => {
  return (
    <>
      <Router>
        <Route path="/todo-list" exact component={TodoList}/>
      </Router>
    </>
  )
}

export default Root