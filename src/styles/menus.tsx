import React from 'react'

import { withStyles } from '@material-ui/core/styles'

import Menu from '@material-ui/core/Menu'
import MenuItem from '@material-ui/core/MenuItem'

export const MainMenu = withStyles({
  paper: {
    background:  'linear-gradient(#e0e0e0, #e0e0e0)',
    border: '1px solid #373737',
  },
})((props: any) => (
  <Menu
    elevation={0}
    getContentAnchorEl={null}
    anchorOrigin={{
      vertical: 'top',
      horizontal: 'center',
    }}
    transformOrigin={{
      vertical: 'top',
      horizontal: 'center',
    }}
    {...props}
  />
))

export const MainMenuItem = withStyles((theme) => ({
  root: {
    '&:focus': {
      backgroundColor: '#cfcfcf',
      '& .MuiListItemIcon-root, & .MuiListItemText-primary': {
        color: theme.palette.common.white,
      },
    },
  },
}))(MenuItem)
