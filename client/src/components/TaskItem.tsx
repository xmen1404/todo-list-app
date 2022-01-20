import { useEffect, useState } from 'react'
import styled from "styled-components"
import axios from 'axios'
import AddTaskIcon from '@mui/icons-material/AddTask';
import DeleteIcon from '@mui/icons-material/Delete';

const TaskItemWrapper = styled.div`
  position: relative;
  margin: auto;
  margin-top: 15px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 1px solid rgb(204, 203, 200);
  border-radius: 10px;
  padding: 10px;
`

const TaskItemName = styled.div`
  flex-grow: 1;
`

const TaskItemControl = styled.button` 
  width: 40px;
  height: 40px;
  background-color: white;
  border: 1px solid rgb(204, 203, 200);
  border-radius: 50%;
  margin-left: 5px;

  &:hover {
    cursor: pointer;
    opacity: 0.7;
  }
`

const TaskItem = (props: Props) => {

  const { taskID, taskName, taskStatus, reloadData } = props 

  const changeTaskStatus = () => {
    var nData = new FormData()
    nData.append('taskid', taskID)
    axios({
      method: 'post', 
      url: 'http://localhost:5000/todo-list/change-task-status', 
      data: nData
    })
    reloadData()
  }

  const removeTask = () => {
    var nData = new FormData()
    nData.append('taskid', taskID)
    axios({
      method: 'post', 
      url: 'http://localhost:5000/todo-list/remove-task', 
      data: nData
    })
    reloadData()
  }

  return (
      <TaskItemWrapper>
        <TaskItemName>{taskName}</TaskItemName>
        <TaskItemControl onClick={changeTaskStatus}>
          { taskStatus ? <AddTaskIcon/> : <></> }
        </TaskItemControl>
        <TaskItemControl onClick={removeTask}>
          <DeleteIcon/>
        </TaskItemControl>
      </TaskItemWrapper>
  )
}

export default TaskItem

type Props = {
  taskID: string,
  taskName: string, 
  taskStatus: boolean, 
  reloadData: () => void
}
