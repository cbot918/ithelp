import './login.css'

import { useState, useContext} from 'react'
import { useNavigate } from 'react-router-dom'
import { UserContext } from '../../App'

function Login(){
  
  const {state, dispatch} = useContext(UserContext)

  const [name, setName] = useState("")

  const navigate = useNavigate()

  return (
    <div className="outer">
      <div> enter your name </div>

      <input 
        className="username"
        type="text"
        placeholder='username'
        onChange={(e)=>{
          setName(e.target.value)
        }}
      />

      <div>
        <input 
          className="submit"
          type="button" 
          value="登入" 
          onClick={(e)=>{
            dispatch({type:"USER", payload:name})
            navigate("/enemylist")
          }}
        />
      </div>
      

    </div>
  )
}

export default Login