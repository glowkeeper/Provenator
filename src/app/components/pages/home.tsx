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

            <Grid item container xs={6}>

                <NavLink to={Local.addUser} className={themeClasses.homeLink}>
                    <Grid item>
                         <FaceTwoToneIcon color="primary"/>
                    </Grid>
                    <Grid item>
                        {Paths.addUser}
                    </Grid>
                </NavLink>

             </Grid>

             <Grid item container xs={6}>

                 <NavLink to={Local.addFile} className={themeClasses.homeLink}>
                     <Grid item>
                          <AttachFileTwoToneIcon color="primary"/>
                     </Grid>
                     <Grid item>
                         {Paths.addFile}
                     </Grid>
                 </NavLink>

              </Grid>

       </Grid>

    </>


  )
}
