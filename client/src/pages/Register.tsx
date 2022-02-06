import { useEffect, useState } from 'react'
import styled from "styled-components"
import axios from 'axios'
import { useNavigate } from 'react-router-dom'


const RegisterWrapper = styled.div`
  position: relative;
  // width: 100vw;
  height: calc(var(--vh, 1vh) * 97);
`

const RegisterInnerWrap = styled.div`
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 150px;
`

const RegisterBanner = styled.div`
    font-size: 40px;
    font-weight: 700;
    color: #37b24d;
`

const RegisterForm = styled.form`
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
`

const RegisterInput = styled.input`
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



const Register = () => {

    const navigate = useNavigate()

    useEffect(() => {
        const vh = window.innerHeight * 0.01;
        document.getElementById("layout-wrapper")!.style.setProperty('--vh', `${vh}px`)
    }, [])

    const submitHandler = (e: any) => {
        e.preventDefault()
        axios({
            method: 'post', 
            url: 'http://localhost:8000/register', 
            data: new FormData( e.target ), 
            withCredentials: true
          })
        .then(response => {
            console.log(response)
            e.target.reset()
            alert("Account created successfully")
            navigate("/login")
        })
        .catch(err => {
            console.log(err.Error())
            alert("Username duplicated. Try other ones!")
          })
    }

    return (
        <RegisterWrapper id="layout-wrapper">
            <RegisterInnerWrap>
                <RegisterBanner>Register</RegisterBanner>
                <RegisterForm onSubmit={submitHandler}>
                    <RegisterInput placeholder='username' type='text' name='username'/>
                    <RegisterInput placeholder='password' type='password' name='password'/>
                    {/* <LoginInput type='submit' name='password'></LoginInput> */}
                    <SubmitButton>Submit</SubmitButton>
                </RegisterForm>
            </RegisterInnerWrap>
        </RegisterWrapper>
    )
}

export default Register

