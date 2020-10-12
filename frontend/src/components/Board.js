import React from 'react'
import styled from 'styled-components'
import BoardCell from './BoardCell'

const BoardTable = styled.table`
  margin 0 auto;
  max-width: 420px;
  max-height: 420px;
  width: 85vw; 
  height: 85vw;
  border-spacing: 2px 2px; 
`

const Board = ({boardState,handleMakeMove}) => {
    return (    
      <BoardTable>    
      <tbody>
              <tr>
                <BoardCell handleMakeMove={handleMakeMove} cell_id={0} boardState={boardState}/>
                <BoardCell handleMakeMove={handleMakeMove} cell_id={1} boardState={boardState}/>
                <BoardCell handleMakeMove={handleMakeMove} cell_id={2} boardState={boardState}/>
              </tr>
              <tr>
                <BoardCell handleMakeMove={handleMakeMove} cell_id={3} boardState={boardState}/>
                <BoardCell handleMakeMove={handleMakeMove} cell_id={4} boardState={boardState}/>
                <BoardCell handleMakeMove={handleMakeMove} cell_id={5} boardState={boardState}/>
              </tr>
              <tr>
                <BoardCell handleMakeMove={handleMakeMove} cell_id={6} boardState={boardState}/>
                <BoardCell handleMakeMove={handleMakeMove} cell_id={7} boardState={boardState}/>
                <BoardCell handleMakeMove={handleMakeMove} cell_id={8} boardState={boardState}/>
              </tr>
            </tbody>
        </BoardTable>
    )
}

export default Board