import styled from 'styled-components'

const Button = styled.button`
  font-weight: bold;
  text-align: center;
  background: transparent;
  display: ${props => props.block ? 'block' : 'inline-block'};
  border:none;
  border-radius: 5px;
  border: 2px solid transparent;
  :hover {
    border-radius: 5px;
    border: 2px solid palevioletred;
  }

  color: palevioletred;
  margin: 1em 1em;
  padding: 0.25em 1em;

  ${props => props.disabled ? 
    `
    background: lightgray;
    ` 
    : ''
  }
`

export default Button