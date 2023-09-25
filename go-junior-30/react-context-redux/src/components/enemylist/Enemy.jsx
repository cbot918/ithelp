
import  './enemy.css'

function Enemy(props){

  const data = {props}.props.props

  return (
    <>
    { console.log( data ) }
    
    <div className="container">
      <div className="name"> Name: {data.name} </div>
      <div className="name"> Job: {data.job} </div>
      <div className="name"> Level: {data.level} </div>
      <div className="name"> Weapon: {data.weapon} </div>
    </div>

    </>
  )

}

export default Enemy