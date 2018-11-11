import React from 'react';
import styled from "styled-components"

const Button = styled.button`
  border: none;
  border-radius: 50px;
  padding: 10px 25px;
  font-size: 16px;
  cursor: pointer;
  box-shadow: 0 1px 4px 0 #aaa;
  background-color: #4353fe;
  color: white;

  &:focus {
    box-shadow: none;
    /*background-color: #2537fe;*/
    outline: none;
  }
`

export default Button
