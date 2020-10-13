import React from 'react'
import { useHistory } from 'react-router-dom'
import Container from './Container'
import Button from './Button'

const MainMenu = (props) => {
  const history = useHistory()

  return (
    <Container>
      <Button block onClick = {() => history.push('/newGame') }>New Game</Button>
      {window.localStorage.getItem('savedGameKey') !== null &&
          <Button block onClick = {() => history.push('/play') }>Resume Game</Button>
      }
      <Button block onClick = {() => history.push('/leaderboard')}>Leaderboard</Button>
    </Container>
  )
}

export default MainMenu
