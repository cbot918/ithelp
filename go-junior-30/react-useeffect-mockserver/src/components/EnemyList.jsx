import { useState, useEffect} from 'react'
import Enemy from './Enemy'
function EnemyList(){

  const [enemy, setEnemy] = useState([{}])

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
    <>{
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