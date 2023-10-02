import { useState } from 'react'


function App() {

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [title, setTitle] = useState("")
  const [body, setBody] = useState("")

  function signupPost({email,password}){
    
    fetch("http://localhost:8886/signup",{
      method:"POST",
      headers:{"Content-Type":"application/json"},
      body:JSON.stringify({email,password})
    })
    .then(res=>res.json())
    .then(data=>{
      console.log(data)
    })
    .catch(err=>{
      console.log(err)
    })
  }

  function signinPost({email,password}){
    fetch("http://localhost:8886/signin",{
      method:"POST",
      headers:{"Content-Type":"application/json"},
      body:JSON.stringify({email,password})
    })
    .then(res=>res.json())
    .then(data=>{
      console.log(data)
      localStorage.setItem("id",data.id)
      localStorage.setItem("token",data.token)
    })
    .catch(err=>{
      console.log(err)
    })
  }

  function postPost({title,body}){
    const id = localStorage.getItem("id")
    const token = localStorage.getItem("token")
    fetch("http://localhost:8886/api/post",{
      method:"POST",
      headers:{
        "Content-Type":"application/json",
        "Authorization":token
      },
      body:JSON.stringify({id,title,body})
    })
    .then(res=>res.json())
    .then(data=>{
      console.log(data)
    })
    .catch(err=>{
      console.log(err)
    })
  }

  return (
    <>
      <div>
        註冊
        <input 
          type="text"
          placeholder="email"
          onChange={(e)=>{
            setEmail(e.target.value)
          }}
        />
        <input 
          type="text"
          placeholder="password"
          onChange={(e)=>{
            setPassword(e.target.value)
          }}
        />
        <input 
          type="button"
          value="submit"
          onClick={(e)=>{
            e.preventDefault()
            signupPost({email,password})
          }}
        />
      </div>
      <div>
        登入
        <input 
          type="text"
          placeholder="email"
          onChange={(e)=>{
            setEmail(e.target.value)
          }}
        />
        <input 
          type="text"
          placeholder="password"
          onChange={(e)=>{
            setPassword(e.target.value)
          }}
        />
        <input 
          type="button"
          value="submit"
          onClick={(e)=>{
            e.preventDefault()
            signinPost({email,password})
          }}
        />
      </div>
      <br />
      <div>
        發文
        <input 
          type="text"
          placeholder="title"
          onChange={(e)=>{
            setTitle(e.target.value)
          }}
        />
        <input 
          type="text"
          placeholder="body"
          onChange={(e)=>{
            setBody(e.target.value)
          }}
        />
        <input 
          type="button"
          value="submit"
          onClick={(e)=>{
            e.preventDefault()
            postPost({title,body})
          }}
        />
      </div>
    </>
  )
}

export default App
