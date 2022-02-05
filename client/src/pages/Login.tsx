import { useEffect, useState } from 'react'
import styled from "styled-components"
import axios from 'axios'

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
`



const Login = () => {

    useEffect(() => {
        const vh = window.innerHeight * 0.01;
        document.getElementById("layout-wrapper")!.style.setProperty('--vh', `${vh}px`)
    }, [])

    const submitHandler = (e: any) => {

        e.preventDefault()
        e.target.reset()
    }

    return (
        <LoginWrapper id="layout-wrapper">
            <LoginInnerWrap>
                <LoginBanner>Login</LoginBanner>
                <LoginForm onSubmit={submitHandler}>
                    <LoginInput placeholder='username' type='text' name='username'/>
                    <LoginInput placeholder='password' type='password' name='password'/>
                    {/* <LoginInput type='submit' name='password'></LoginInput> */}
                    <SubmitButton>Submit</SubmitButton>
                </LoginForm>
            </LoginInnerWrap>
        </LoginWrapper>
    )
}

export default Login

