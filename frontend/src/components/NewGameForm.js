import React, { useState } from 'react'
import styled from 'styled-components'

import {useHistory} from 'react-router-dom'
import gameApi from '../services/gameApi'
import Container from './Container'


const Input = styled.input`
  display: block;
  border: 0;
  border-bottom: 1px solid darkgray;
  margin: 10% auto;
  font-size: 1.5em;
  background-color: transparent;
  &::placeholder {
    color: darkgray;
    opacity: 0.6;
  }
`;

const Label = styled.label`
  font-size: 1.5em;
`

const Radio  = styled.input`
  display: flex;
  border: 0;
  border-bottom: 1px solid darkgray;
  font-size: 1.5em;
  background-color: transparent;
`;

const Submit = styled.input`
  background-color: ${props => (props.disabled ? 'gray' : 'darkgray')};
  color: white;
  border: 0;
  font-weight: 700;
  padding: .8em;
  display: flex;
  margin: auto;
  font-size: 1em;
  cursor: ${props => (props.disabled ? 'not-allowed' : 'pointer')};
`;

const RadioContainer = styled.div`
  margin-bottom: 10px;  
  display: flex;
  align-items:center;
  justify-content: ${props => props.outer ? 'space-between' : 'flex-start'};
`

const NewGameForm = (props) => {
  const [playerName,setPlayerName] = useState('')

  const history = useHistory();
    
  const handleNewGame = (event) => {
    event.preventDefault()
    gameApi
    .newGame({'name': event.target.name.value ,'difficulty': event.target.difficulty.value})
    .then( (res) => {
      console.log(res)
      props.setGameKey(res.game_key)
      window.localStorage.setItem('savedGameKey',res.game_key)
      history.push('/play')
    })
  }

  return (
    <Container>
      <form onSubmit={handleNewGame}>
        <Input
            type="text"
            name="name"
            value={playerName}
            onChange={(e) => setPlayerName(e.target.value)}
            placeholder="Enter your name"
          />
          <RadioContainer outer>
            <RadioContainer>
              <Radio defaultChecked  name="difficulty" type="radio" value="Easy" id="easy" />
              <Label htmlFor="easy">Easy</Label>
            </RadioContainer>
            <RadioContainer>
              <Radio name="difficulty" type="radio" value="Impossible" id="impossible" />
              <Label htmlFor="impossible">Impossible</Label>
            </RadioContainer>
          </RadioContainer>
          <Submit disabled={playerName == ''} value="New Game" type="submit"/>
      </form>
    </Container>
  )
}

export default NewGameForm