import { useEffect, useState } from 'react'
import styled from "styled-components"
import axios from 'axios'
import SendIcon from '@mui/icons-material/Send';
import { TaskItem } from '../components/index'
import { Navigate } from 'react-router-dom'

const LayoutWrapper = styled.div`
  position: relative;
  // width: 100vw;
  height: calc(var(--vh, 1vh) * 97);
`

const LayoutInnerWrap = styled.div`
  position: relative;
  height: 100%;
  // margin: 25px;
  margin: auto;
  width: 80%;
  max-width: 500px;
  display: flex;
  flex-direction: column;
`

const AddTaskForm = styled.form`
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
  border: 1px solid rgb(204, 203, 200);
  border-radius: 10px;
`

const AddTaskInput = styled.input`
  height: 35px;
  flex-grow: 1;
  border: none;
  border-radius: 10px;
  font-size: 20px;
  &:focus {
    outline: none;
  }
`

const AddTaskButton = styled.button`
  width: 40px;
  height: 40px;
  border: none;
  background-color: white;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;

  &:hover {
    opacity: 0.5;
    cursor: pointer;
  }
`
const TaskList = styled.ul`
  position: relative;
  padding:0px;
  height: 5%;
  flex-grow: 1;
  overflow: scroll;
  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;  /* IE and Edge */
  scrollbar-width: none;  /* Firefox */
`

const TodoList = () => {

  const [todoListData, setTodoListData] = useState<[TodoItem]|null>(null)

  const submitHandler = (e: any) => {
    axios({
      method: 'post', 
      url: 'http://localhost:8000/todo-list/add-task', 
      data: new FormData( e.target )
    })
    .then(response => {
      console.log(response)
      loadData()
    })
    .catch(err => {
      if(err.response) {
        console.log(err.response.status == 403) 
      } else if(err.request) {
        console.log("server did not responded")
      } else {
        console.log("error in settings before request")
      }
    })
    e.preventDefault()
    e.target.reset()
  }

  const loadData = () => {
    axios({
      method: 'get', 
      url: 'http://localhost:8000/todo-list/get-task-list'
    })
    .then(response => {
      setTodoListData(response.data.todolist)
    })
    .catch(err => {
      if(err.response) {
        console.log(err.response.status == 403) 
      } else if(err.request) {
        console.log("server did not responded")
      } else {
        console.log("error in settings before request")
      }
    })
  }

  const authUser = () => {
    let result = false
    axios({
      method: 'get', 
      url: 'http://localhost:8000/todo-list/get-task-list'
    })
    .then(response => {
      result = true
    })
    .catch(err => {
      result = false 
    })
    return result
  }
 
  useEffect(() => {
    const vh = window.innerHeight * 0.01;
    document.getElementById("layout-wrapper")!.style.setProperty('--vh', `${vh}px`)

    loadData()
  }, [])

  if(authUser()) {
    return (
      <LayoutWrapper id="layout-wrapper">
        <LayoutInnerWrap>
          <AddTaskForm onSubmit={submitHandler}> 
            <AddTaskInput type="text" name="taskname"/>  
            <AddTaskButton type="submit">
              <SendIcon />
            </AddTaskButton>
          </AddTaskForm>
          <TaskList>
            {
              todoListData?.map((item: TodoItem) => (
                <TaskItem 
                  taskID={item.taskid} 
                  taskName={item.taskname} 
                  taskStatus={item.taskstatus}
                  reloadData={loadData}/>
              ))
            }
          </TaskList>
        </LayoutInnerWrap>
      </LayoutWrapper>
    )
  } else {
    return <Navigate to='/' /> 
  }
}

type TodoItem = {
  taskid: string, 
  taskname: string, 
  taskstatus: boolean
}

type TodoListData = {
  todolist: [TodoItem]
}

export default TodoList

