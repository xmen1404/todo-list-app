import { useEffect, useState } from 'react'
import styled from "styled-components"
import axios from 'axios'
import SendIcon from '@mui/icons-material/Send';
import { TaskItem } from '../components/index'
import { Navigate, useNavigate } from 'react-router-dom'

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
  const navigate = useNavigate()

  const submitHandler = (e: any) => {
    e.preventDefault()
    axios({
      method: 'post', 
      url: 'http://localhost:8000/todo-list/add-task', 
      data: new FormData( e.target ), 
      withCredentials: true, 
      // headers:{Cookie:"029857ec-10c3-4ad3-881b-b82ebf3c9683"}
    })
    .then(response => {
      console.log(response)
      e.target.reset()
      loadData()
    })
    .catch(err => {
      console.log(err)
      navigate('/login')
    })
    
    
  }

  const loadData = () => {
    axios({
      method: 'get', 
      url: 'http://localhost:8000/todo-list/get-task-list', 
      // headers: { Access-Control-Allow-Origin: "https://amazing.site"}
      withCredentials: true
    })
    .then(response => {
      setTodoListData(response.data.todolist)
    })
    .catch(err => {
      console.log(err)
      navigate('/login')
    })
  }
 
  useEffect(() => {
    const vh = window.innerHeight * 0.01;
    document.getElementById("layout-wrapper")!.style.setProperty('--vh', `${vh}px`)

    loadData()
  }, [])

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

