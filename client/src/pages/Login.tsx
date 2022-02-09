import { useEffect, useState } from 'react'
import styled from "styled-components"
import axiosInstance from '../axiosInstance'
import { useNavigate } from 'react-router-dom'

// const { REACT_APP_SERVER_DOMAIN, NODE_ENV } = process.env

const LoginWrapper = styled.div`
  position: relative;
  // width: 100vw;
  height: calc(var(--vh, 1vh) * 97);
`

const LoginInnerWrap = styled.div`
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 150px;
`

const LoginBanner = styled.div`
    font-size: 40px;
    font-weight: 700;
    color: #37b24d;
`

const LoginForm = styled.form`
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
`

const LoginInput = styled.input`
    height: 35px;
    padding: 10px;
    margin-top: 20px;
    border: 2px solid #ced4da;
    border-radius: 10px;
    font-size: 20px;
    &:focus {
        outline: none;
    }
`

const SubmitButton = styled.button`
  width: 120px;
  height: 40px;
  margin-top: 20px;
  border: none;
  border-radius: 15px;
  color: white;
  background-color: #37b24d;
  font-size: 20px;
  &:hover {
    cursor: pointer;
  }
`

const RegisterLink = styled.a`
  margin-top: 30px;
  font-size: 15px;
`


const Login = () => {

    const navigate = useNavigate()

    useEffect(() => {
        const vh = window.innerHeight * 0.01;
        document.getElementById("layout-wrapper")!.style.setProperty('--vh', `${vh}px`)
    }, [])

    const submitHandler = (e: any) => {
        e.preventDefault()
        axiosInstance({
            method: 'post', 
            url: `/login` , 
            data: new FormData( e.target ), 
            withCredentials: true
          })
        .then(response => {
            console.log(response)
            e.target.reset()
            navigate("/todo-list")
        })
        .catch(err => {
            if(err.response) {
              console.log(err.response.status == 403) 
            } else if(err.request) {
              console.log("server did not responded")
            } else {  
              console.log("error in settings before request")
            }
            alert("Wrong username or password")
          })
    }

    return (
        <LoginWrapper id="layout-wrapper">
            <LoginInnerWrap>
                <LoginBanner>Login</LoginBanner>
                <LoginForm onSubmit={submitHandler}>
                    <LoginInput placeholder='username' type='text' name='username'/>
                    <LoginInput placeholder='password' type='password' name='password'/>
                    {/* <LoginInput type='submit' name='password'></LoginInput> */}
                    <SubmitButton>Login</SubmitButton>
                </LoginForm>
                <RegisterLink href="/register">Click here to register!</RegisterLink>
            </LoginInnerWrap>
        </LoginWrapper>
    )
}

export default Login

