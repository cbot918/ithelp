import { useState } from 'react'

function Hello(){
  const [name, setName] = useState("")
  const [enemy, setEnemy] = useState("")

  function requestData(){
    fetch("http://localhost:8886/mock")
    .then(res=>res.json())
    .then(data=>{
      console.log(data)
      setEnemy(data)
    })
    .catch(err=>{
      console.log(err)
    })
  }

  return (
    <>
        <input 
          type="text" 
          placeholder='your name?' 
          onChange={(e)=>{
            setName(e.target.value)
          }}
        />

        <div> Hello { name } </div>

        <input
          type="button"
          value="require challengers"
          onClick={()=>{
            requestData()
          }}
        />
        <div> You Choose {enemy.name}</div>
    </>
  )

}

export default Hello