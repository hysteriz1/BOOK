import React, { FC } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  } from '@backstage/core';
import { Grid } from '@material-ui/core';
import Link from '@material-ui/core/Link';
import { Link as RouterLink } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import Table from '../Table';


const Pagemain: FC<{}> = () => {
  const profile = { givenName: '' };

  return (
    <Page theme={pageTheme.home}>
     <Header
         title={`${profile.givenName || 'ระบบจองห้องพัก'}`}
         subtitle="Blue Moon"        
     ></Header>
     <Content>
       <Grid container >
         <Grid item xs={12} md={12}>
             <Table></Table>
           </Grid>
           <br/>
           <Grid item xs={12} md={12} align="center">       
           <Link component={RouterLink} to="/Page2">
           <Button variant="contained" color="primary">
           จองห้องพัก           
           </Button></Link>
            </Grid> 
     </Grid>
     </Content>
   </Page>
  );
};
export default Pagemain;
