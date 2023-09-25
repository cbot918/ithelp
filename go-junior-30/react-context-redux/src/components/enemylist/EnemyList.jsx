import { useState, useEffect, useContext} from 'react'
import { UserContext } from '../../App'
import Enemy from './Enemy'
function EnemyList(){

  const [enemy, setEnemy] = useState([{}])
  const {state, dispatch} = useContext(UserContext)
  useEffect(()=>{
    
    fetch("http://localhost:8886/mock")
    .then(res=>res.json())
    .then(data=>{
      // console.log(data)
      setEnemy(data)
    })
    .catch(err=>{
      console.log(err)
    })

  }, [])

  return(
    <>
    <h2 className="title"> {state} Challenge List </h2>
    {
      enemy.map((e)=>{
        return (         
          <Enemy props = {e} key={e.id}/>
        )
      })
    }
    </>
  )
  

}

export default EnemyList