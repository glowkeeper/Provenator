import React from 'react'

import { createMuiTheme, responsiveFontSizes } from '@material-ui/core/styles'
import { makeStyles } from '@material-ui/styles'
import { withStyles } from '@material-ui/core/styles'

import red from '@material-ui/core/colors/red'
import blue from '@material-ui/core/colors/blue'
import green from '@material-ui/core/colors/blue'
import indigo from '@material-ui/core/colors/indigo'
import orange from '@material-ui/core/colors/orange'
import yellow from '@material-ui/core/colors/yellow'
import grey from '@material-ui/core/colors/grey'

import GoogleFontLoader from 'react-google-font-loader'

//load this first
const fontLoader = () =>
    <GoogleFontLoader
      fonts={[
        {
          font: 'Barlow',
          weights: [400, '400i'],
        },
        {
          font: 'Lato',
          weights: [400, 500],
        },
      ]}
      subsets={['cyrillic-ext', 'greek']}
    />

let theme = createMuiTheme ({
  spacing: 8,
  typography: {
    fontFamily: [
      'Barlow',
      'Lato',
      'sans-serif',
      'Arial',
      'Roboto',
      '-apple-system',
    ].join(','),
    fontSize: 16,
    fontWeightLight: 300,
    fontWeightRegular: 400,
    fontWeightMedium: 500,
    h1: {
      fontSize: "2rem",
      fontWeight: 400,
      fontFamily: "\"Barlow\", \"Arial\", \"sans-serif\", \"Roboto\"",
      color: '#000000'
    },
    h2: {
      fontSize: "1.5rem",
      fontWeight: 400,
      fontFamily: "\"Barlow\", \"Arial\", \"sans-serif\", \"Roboto\"",
      color: '#000000'
    },
    h3: {
      fontSize: "1.25rem",
      fontWeight: 400,
      fontFamily: "\"Barlow\", \"Arial\", \"sans-serif\", \"Roboto\"",
      color: '#000000'
    },
    h4: {
      fontSize: "1.1rem",
      fontWeight: 400,
      fontFamily: "\"Barlow\", \"Arial\", \"sans-serif\", \"Roboto\"",
      color: '#000000'
    },
    subtitle1: {
      fontSize: "0.9rem",
      fontWeight: 400,
      fontFamily: "\"Barlow\", \"Arial\", \"sans-serif\", \"Roboto\"",
      lineHeight: "1.5em",
      color: '#38348B'
    },
    body1: {
      fontSize: "0.875rem",
      fontWeight: 400,
      fontFamily: "\"Lato\", \"Arial\", \"sans-serif\", \"Roboto\"",
      lineHeight: "1.46429em",
      color: '#38348B'
    },
    body2: {
      fontSize: "0.8rem",
      fontWeight: 400,
      fontFamily: "\"Lato\", \"Arial\", \"sans-serif\", \"Roboto\"",
      lineHeight: "1.4em",
      color: '#38348B'
    },
    caption: {
      fontSize: "0.75rem",
      fontWeight: 400,
      fontFamily: "\"Lato\", \"Arial\", \"sans-serif\", \"Roboto\"",
      lineHeight: "1.375em",
      color: orange[900]
    },
    button: {
      fontSize: "0.875rem",
      textTransform: "uppercase",
      fontWeight: 500,
      fontFamily: "\"Lato\", \"Arial\", \"sans-serif\", \"Roboto\"",
      color: '#38348B'
    }
  },
  palette: {
    type: 'dark',
    background: {
      default: '#FFFFFF',
    },
    text: {
      primary: "#424242",
      secondary: "#424242"
    },
    primary: {
      main: '#424242'
    },
    secondary: {
      main: '#bdbdbd'
    },
    error: red,
    warning: orange,
    info: grey,
    success: green,
  }
})

theme = responsiveFontSizes(theme)
theme.spacing(4)

const themeStyles = makeStyles({
  root: {
    background: 'linear-gradient(#FFFFFF, #FFFFFF)',
    color: theme.palette.text.primary,
    height: "100vh",
    width: "100%"
  },
  grid: {
    height: "100vh",
    width: "100%"
  },
  logo: {
    padding: theme.spacing(1),
    margin: theme.spacing(1),
    textAlign: 'left',
  },
  header: {
    padding: theme.spacing(1),
    margin: theme.spacing(1),
    textAlign: 'center',
    background: 'transparent linear-gradient(180deg, #8d8d8d 0%, #bdbdbd 100%) 0% 0% no-repeat padding-box',
    height: "15%",
    width: "100%"
  },
  title: {
    padding: theme.spacing(1),
    margin: theme.spacing(1),
    textAlign: 'center',
    background: 'linear-gradient(#38348B, #38348B)'
  },
  subTitle: {
    padding: theme.spacing(1),
    margin: theme.spacing(1),
    textAlign: 'center',
    background: 'linear-gradient(#38348B, #38348B)'
  },
  content: {
    padding: theme.spacing(2),
    margin: theme.spacing(2),
    color: theme.palette.text.primary,
    background: 'linear-gradient(#FFFFFF, #FFFFFF)',
    width: "100%"
  },
  caption: {
    padding: theme.spacing(1),
    margin: theme.spacing(1),
    textAlign: 'center',
    background: 'linear-gradient(#4CB79C, #4CB79C)'
  },
  footer: {
    padding: theme.spacing(2),
    margin: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
    background: 'transparent linear-gradient(180deg, #8d8d8d 0%, #bdbdbd 100%) 0% 0% no-repeat padding-box',
    height: "10%",
    width: "100%"
  },
  button: {
    padding: theme.spacing(1),
    margin: theme.spacing(1),
    fontSize: "4rem",
    fontWeight: 400,
    textAlign: 'right',
    textTransform: 'none',
    color: theme.palette.common.white,
    background: 'linear-gradient(#424242, #424242)'
  },
  menu: {
    padding: 0,
    margin: 0,
    textAlign: 'left',
    background: 'linear-gradient(#cfcfcf, #cfcfcf)'
  },
  error: {
    color: "#FF0000",
    textAlign: 'left',
  },
  spinner: {
     position: 'absolute',
     left: '50%',
     top: '50%',
     transform: 'translate(-50%, -50%)'
  }
})

export { fontLoader, theme, themeStyles }
