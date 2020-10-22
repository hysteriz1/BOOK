import { createPlugin } from '@backstage/core';
import Page2 from './components/Page2';
import Pagemain from './components/Pagemain';
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/Page2', Page2);
    router.registerRoute('/', Pagemain);
  },
});
