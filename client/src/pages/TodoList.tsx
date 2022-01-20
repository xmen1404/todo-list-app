import { useEffect, useState } from 'react'
import styled from "styled-components"
import axios from 'axios'
import SendIcon from '@mui/icons-material/Send';
import { TaskItem } from '../components/index'

const LayoutWrapper = styled.div`
  position: relative;
  width: 100vw;
  height: calc(var(--vh, 1vh) * 100);
`

const LayoutInnerWrap = styled.div`
  position: relative;
  height: 100%;
  // margin: 25px;
  margin: auto;
  width: 80%;
  max-width: 500px;
`

const AddTaskForm = styled.form`
  position: relative;
  margin: auto;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 10px;
  border: 1px solid rgb(204, 203, 200);
  border-radius: 10px;
`

const AddTaskInput = styled.input`
  height: 30px;
  flex-grow: 1;
  border: none;
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

const TodoList = () => {

  useEffect(() => {
    const vh = window.innerHeight * 0.01;
    document.getElementById("layout-wrapper")!.style.setProperty('--vh', `${vh}px`)
  }, [])

  const submitHandler = (e: any) => {
    // console.log(e.target[0].value)
    // var temp = new FormData( e.target )
    // console.log(temp.get('taskname'))
    // axios.post("", temp)
    axios({
      method: 'post', 
      url: 'http://localhost:5000/todo-list/add-task', 
      data: new FormData( e.target )
    })
    .then(response => console.log(response))
    // fetch("http://localhost:5000/todo-list/add-task", {
    //   method: 'POST', 
    //   mode: 'no-cors',
    //   body: new FormData( e.target )
    // }).then(response => console.log(response))
    e.preventDefault()
    e.target.reset()
  }

  

  return (
    <LayoutWrapper id="layout-wrapper">
      <LayoutInnerWrap>
        <AddTaskForm onSubmit={submitHandler}> 
          <AddTaskInput type="text" name="taskname"/>  
          <AddTaskButton type="submit">
            <SendIcon />
          </AddTaskButton>
        </AddTaskForm>
        <TaskItem taskName="bruh" taskStatus={false}/>
      </LayoutInnerWrap>
    </LayoutWrapper>
  )
}

export default TodoList