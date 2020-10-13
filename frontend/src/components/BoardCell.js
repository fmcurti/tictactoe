import React from 'react'
import styled from 'styled-components'

const Slot = styled.td`
  padding: 0;
  width: 33%; 
  height: 33%;
`

const Cell = styled.button`
  font-size: 4rem;
  height: 100%; 
  width: 100%;
`

const BoardCell = ({ cell_id, boardState, handleMakeMove }) => {
  return (
    <Slot>
      <Cell onClick={() => handleMakeMove(cell_id)}>{boardState[cell_id]}</Cell>
    </Slot>
  )
}

export default BoardCell
