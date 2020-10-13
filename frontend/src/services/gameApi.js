import axios from 'axios'
const baseUrl = 'https://powerful-journey-30384.herokuapp.com'

const delay = retryCount =>
  new Promise(resolve => setTimeout(resolve, 10 ** retryCount))

const getGame = async (game_key, retryCount = 0, lastError = null) => {
  try {
    const response = await axios.get(`${baseUrl}/game/get?game_key=${encodeURIComponent(game_key)}`)
    return response.data
  } catch (e) {
    await delay(retryCount)
    return getGame(game_key, retryCount + 1, e)
  }
}

const newGame = async (gameOptions, retryCount = 0, lastError = null) => {
  try {
    const response = await axios.post(`${baseUrl}/game/new`, gameOptions)
    return response.data
  } catch (e) {
    await delay(retryCount)
    return getGame(gameOptions, retryCount + 1, e)
  }
}

const makeMove = async (move, retryCount = 0, lastError = null) => {
  try {
    const response = await axios.post(`${baseUrl}/game/make_move`, move)
    return response.data
  } catch (e) {
    await delay(retryCount)
    return makeMove(move, retryCount + 1, e)
  }
}

const highscores = async (page, retryCount = 0, lastError = null) => {
  try {
    const response = await axios.get(`${baseUrl}/highscores?offset=${page * 10}&limit=10`)
    return response.data
  } catch (e) {
    await delay(retryCount)
    return makeMove(page, retryCount + 1, e)
  }
}

export default { getGame, newGame, makeMove, highscores }
