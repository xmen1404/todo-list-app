import { useEffect, useState } from 'react'
import styled from "styled-components"
import axios from 'axios'
import DoneIcon from '@mui/icons-material/Done';
import DeleteIcon from '@mui/icons-material/Delete';

const TaskItemWrapper = styled.li<{taskStatus: boolean}>`
  position: relative;
  margin: auto;
  margin-top: 15px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 1px solid rgb(104, 103, 100);
  border-radius: 10px;
  padding: 10px;
  opacity: ${props => props.taskStatus?'0.2':'1'};
`

const TaskItemName = styled.div`
  flex-grow: 1;
`

const TaskItemControl = styled.button` 
  width: 40px;
  height: 40px;
  background-color: white;
  border: 1px solid rgb(104, 103, 100);
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
      url: 'http://localhost:8000/todo-list/change-task-status', 
      data: nData
    }).then(response => {
      reloadData()
    })
  }

  const removeTask = () => {
    var nData = new FormData()
    nData.append('taskid', taskID)
    axios({
      method: 'post', 
      url: 'http://localhost:8000/todo-list/remove-task', 
      data: nData
    }).then(response => {
      reloadData()
    })
  }

  return (
      <TaskItemWrapper taskStatus={taskStatus}>
        <TaskItemName>
          {taskStatus ? <s>{taskName}</s> : taskName}
        </TaskItemName>
        <TaskItemControl onClick={changeTaskStatus}>
          { taskStatus ? <DoneIcon/> : <></> }
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
