import { store } from './redux.js'

const log = console.log

window.onload = ()=>{
  

  let body = document.querySelector("body")
  
  const addThree = document.createElement('input')
  addThree.setAttribute("type","button")
  addThree.setAttribute("value","+3")
  addThree.addEventListener('click',()=>{
    store.dispatch( {type:"add", payload:3} )
  })
  body.appendChild(addThree)
  
  const minusTwo = document.createElement('input')
  minusTwo.setAttribute("type","button")
  minusTwo.setAttribute("value","-2")
  minusTwo.addEventListener('click',()=>{
    store.dispatch( {type:"minus", payload:2} )
  })
  body.appendChild(minusTwo)

  const view = document.createElement('div')
  body.appendChild(view)
  
  
  store.subscribe(()=>{
    render( view, store)
  })

  log(store.getState())


  


}

function render( target, store ){

  target.innerHTML = `
  <div> count: ${store.getState()}</div>
`
}
