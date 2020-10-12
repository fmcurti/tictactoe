import React, { useState,useEffect,useCallback } from 'react'
import styled from 'styled-components'
import './App.css';
import {
  BrowserRouter as Router,
  Switch, Route
} from "react-router-dom"

import Game from './components/Game'
import MainMenu from './components/MainMenu'
import NewGameForm from './components/NewGameForm'
import Leaderboard from './components/Leaderboard'


const Title = styled.h1`
  text-align:center;
`

const App = (props) => {
  
  const [gameKey,setGameKey] = useState(() => window.localStorage.getItem('savedGameKey'))

  return (
    <Router>
        <Title>Tic Tac Toe</Title>
        <Switch>
            <Route path="/play">
              <Game gameKey={gameKey} />
            </Route>
            <Route path="/newGame">
              <NewGameForm setGameKey={setGameKey}/>
            </Route>
            <Route path="/leaderboard">
              <Leaderboard/>
            </Route>
            <Route path="/">
                <MainMenu/>
            </Route>
        </Switch>
    </Router>
  )
}

export default App;
