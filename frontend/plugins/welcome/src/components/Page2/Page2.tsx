import React, { FC } from 'react';
import {
  Content,
  Header,
  InfoCard,
  Page,
  pageTheme,
  } from '@backstage/core';
import Create from '../Create';
import { Grid } from '@material-ui/core';

const Page2: FC<{}> = () => {
  const profile = { givenName: '' };

  return (
    <Page theme={pageTheme.home}>
     <Header
         title={`${profile.givenName || 'ระบบจองห้องพัก'}`}
         subtitle="Blue Moon"        
     ></Header>
     <Content>
       <Grid container justify="center">
         <Grid item>
       <InfoCard>
     <Create></Create>
     </InfoCard>
     </Grid>
     </Grid>
     </Content>
   </Page>
  );
};
export default Page2;
