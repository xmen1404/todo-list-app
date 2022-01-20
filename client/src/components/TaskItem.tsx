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

  const { taskName, taskStatus } = props 

  const changeTaskStatus = () => {
    
  }

  return (
      <TaskItemWrapper>
        <TaskItemName>{taskName}</TaskItemName>
        <TaskItemControl>
          { taskStatus ? <AddTaskIcon/> : <></> }
        </TaskItemControl>
        <TaskItemControl>
          <DeleteIcon/>
        </TaskItemControl>
      </TaskItemWrapper>
  )
}

export default TaskItem

type Props = {
  taskName: string, 
  taskStatus: boolean
}
