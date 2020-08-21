import React from 'react'

import { withStyles } from '@material-ui/core/styles'

import Button from '@material-ui/core/Button'

export const Okay = withStyles({
  root: {
    boxShadow: 'none',
    textTransform: 'none',
    fontSize: 16,
    padding: '6px 12px',
    border: '1px solid',
    lineHeight: 1.5,
    background:'linear-gradient(#9e9e9e, #9e9e9e)',
    backgroundColor: '#9e9e9e',
    borderColor: '#212121',
    fontFamily: "\"Barlow\", \"Arial\", \"sans-serif\", \"Roboto\"",
    '&:hover': {
      backgroundColor: '#494949',
      borderColor: '#424242',
      boxShadow: 'none',
    },
    '&:active': {
      boxShadow: 'none',
      backgroundColor: '#424242',
      borderColor: '#212121',
    },
    '&:focus': {
      boxShadow: '0 0 0 0.2rem rgba(0,123,255,.5)',
    },
  },
})(Button)
