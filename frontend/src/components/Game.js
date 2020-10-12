import React, { useState, useCallback, useLayoutEffect } from 'react'
import gameApi from '../services/gameApi'
import Container from './Container'
import styled, { css, keyframes } from 'styled-components' 
import Board from './Board'
import Button from './Button'
import {useHistory} from 'react-router-dom'
import useInterval from '@use-it/interval';

const BoardContainer = styled.div`
  display:flex;
  justify-content:center; 
`

const Spin = keyframes`
0% {
  transform: rotate(0deg);
}
100% {
  transform: rotate(360deg);
}
`

const Spinner = styled.div`
  height: 64px;
  width: 64px;
  border-radius: 50%;
  border: 3px solid gray;
  border-bottom: none;
  animation: ${Spin} 1.4s infinite linear;
`


const Game = (props) => {
  
    const history = useHistory();
    const [playerTurn,setPlayerTurn] = useState(true)
    const [winner,setWinner]         = useState(0)
    const [boardState,setBoardState] = useState([])
    const [isLoading,setIsLoading] = useState(true)
  
    const handleMakeMove = (cell_id) => {
      if (playerTurn && boardState[cell_id] == '') {
        gameApi
          .makeMove(
            {'cell': cell_id,
            'game_key': props.gameKey})
          .then( _ => {
            updateGameStateCallback()
          })
      }
    }
  
    const gameNotification = () => {
      if (winner === 0){
        return playerTurn ? 'Make your move' : 'Wait for your turn'
      } else {
        switch (winner) {
          case 1:
            return 'You Won!'
          case 2:
            return 'You Lost =('
          case 3:
            return 'Tie'
        }
      }
  
    }
  
    const updateGameStateCallback = useCallback(() => {
      gameApi
      .getGame(props.gameKey)
      .then(gameState => {
        console.log(gameState)
        const newBoard = gameState.board.map(cell => cell === 'E' ? '' : cell )
        setBoardState(newBoard)
        setPlayerTurn(gameState.current_turn ? false : true)
        setWinner(gameState.winner) 
        if (gameState.winner > 0) {
          window.localStorage.removeItem('savedGameKey')
        }
        setIsLoading(false)
      }
    ) },[props.gameKey,setBoardState,setPlayerTurn,setWinner])
  
    useLayoutEffect(() =>  {
      if (props.gameKey === null) {
        history.replace('/')
      } else{
        updateGameStateCallback()
  
      }
    },[props.gameKey,history,updateGameStateCallback])
  
    useInterval(() => {
      if (!playerTurn) {
        updateGameStateCallback()
      }
    },1000)
  
    return (
      <Container>
        <h3>{gameNotification()}</h3>
        <BoardContainer>
          {isLoading ? <Spinner/> :        
          <Board boardState={boardState} handleMakeMove={handleMakeMove}/>}
        </BoardContainer>
        {winner > 0 && <Button onClick={() => history.replace('/') }>Go back</Button>  }
      </Container>
    )
  }
  

  export default Game