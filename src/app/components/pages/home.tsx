import React from 'react'
import { NavLink } from 'react-router-dom'

import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import FaceTwoToneIcon from '@material-ui/icons/FaceTwoTone';
import AttachFileTwoToneIcon from '@material-ui/icons/AttachFileTwoTone'

import { themeStyles } from '../../styles'

import { Paths, Local } from '../../config'

export const Home = () => {

  const themeClasses = themeStyles()

  return (

     <>

       <Grid container>

            <Grid item container xs={6} justify="center">

                <NavLink to={Local.addUser} className={themeClasses.homeLink}>
                    <Grid item>
                        <Paper className={themeClasses.home} elevation={0}>
                           <FaceTwoToneIcon color="primary"/>
                        </Paper>
                    </Grid>
                    <Grid item>
                        <Paper className={themeClasses.home} elevation={0}>
                            {Paths.addUser}
                        </Paper>
                    </Grid>
                </NavLink>

             </Grid>

             <Grid item container xs={6} justify="center">

                 <NavLink to={Local.addFile} className={themeClasses.homeLink}>
                     <Grid item>
                        <Paper className={themeClasses.home} elevation={0}>
                          <AttachFileTwoToneIcon color="primary"/>
                        </Paper>
                     </Grid>
                     <Grid item>
                        <Paper className={themeClasses.home} elevation={0}>
                            {Paths.addFile}
                        </Paper>
                     </Grid>
                 </NavLink>

              </Grid>

       </Grid>

    </>


  )
}
