import EnemyList from './components/enemylist/EnemyList'
import Login from './components/login/Login'
import {  Routes, Route  } from "react-router-dom";
import { createContext, useReducer } from 'react';
import { reducer,initialState } from './reducer/useReducer'

export const UserContext = createContext()

function App() {
  
  const [state, dispatch] = useReducer(reducer,initialState)

  return (
    <>
      <UserContext.Provider value={{state, dispatch}}>
        <Routes>
          <Route path="*" element={<Login />}></Route>
          <Route path="/enemylist" element={<EnemyList/> } />
        </Routes>
      </UserContext.Provider> 
    </>
  )
}

export default App
