import React from 'react';
import ReactDOM from 'react-dom';
import TodoList from './pages/TodoList'
import Login from './pages/Login'
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

const Root = () => {
  return (
      <Router>
        <Routes>
          <Route path="/todo-list" element={<TodoList/>}/>
          <Route path="/" element={<Login/>}/>
        </Routes>
      </Router>
  )
}

const IndexHtml = document.getElementById('root')

ReactDOM.render(
  <React.StrictMode>
    <Root />
  </React.StrictMode>,
  document.getElementById('root')
);
