function reducer(state = 0, action){
  if( action && action.type === "add"){
    return state + action.payload
  } 
  if ( action && action.type === "minus"){
    return state - action.payload
  }

  return state
}


function createStore( reducer ){
  let state = reducer()
  let eventhub = []

  function getState(){
    return state
  }

  function dispatch( action ){
    state = reducer(state, action)
    eventhub.forEach( (fn) => fn() )
  }

  function subscribe(fn){
    eventhub.push(fn)
  }

  return {
    getState,
    dispatch,
    subscribe
  }
}

const store = createStore(reducer)

export { store }