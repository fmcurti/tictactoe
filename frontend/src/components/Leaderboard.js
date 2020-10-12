import React, { useState,useEffect,useCallback } from 'react'
import styled from 'styled-components'
import {useHistory} from 'react-router-dom'

import gameApi from '../services/gameApi'
import Container from './Container'
import Button from './Button'



const Table = styled.table`
  text-align: center;
  width: 50%; 
  border-collapse: collapse; 
`

const Entry = styled.td`
  padding: 6px; 
  border: 1px solid #ccc; 
`

const ButtonGroup = styled.div`
  margin-bottom: 10px;  
  display: flex;
  align-items:center;
  justify-content: space-between;
`

const Leaderboard = (props) => {

  const [page,setPage] = useState(0)
  const [leaderboard,setLeaderboard] = useState([])
  const [maxPages,setMaxPages] = useState(0)

  const history = useHistory()

  const getLeaderboardCallback = useCallback(() => {
    gameApi
    .highscores(page)
    .then(res => {
      setLeaderboard(res.entries)
      setMaxPages(res.pages)
    })
  },[setLeaderboard,setMaxPages,page])

  const handlePageChange = newPage => {
    setPage(newPage)
  }

  useEffect(() => {
    getLeaderboardCallback()
  },[getLeaderboardCallback])

  const rows = () => leaderboard.map((entry,i) => 
      <tr key={i}>
        <Entry>{entry.player_name}</Entry>
        <Entry>{entry.play_time}</Entry>
      </tr>
  )
  return (
    <Container>
      <Table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Time played</th>
          </tr>
        </thead>
        <tbody>
            {rows()}
        </tbody>
      </Table>
      <ButtonGroup>
        <Button disabled={page == 0} onClick={() => handlePageChange(page - 1)} >Previous</Button>
        <Button disabled={page == maxPages - 1} onClick={() => handlePageChange(page + 1)}>Next</Button> 
      </ButtonGroup>
      <Button onClick={() => history.push('/')}>Back to Main Menu</Button>
    </Container>
  )

} 


export default Leaderboard