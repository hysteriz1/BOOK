import React from 'react';
import { render } from '@testing-library/react';
import Page3 from './Page2';
import { ThemeProvider } from '@material-ui/core';
import { lightTheme } from '@backstage/theme';

describe('WelcomePage', () => {
  it('should render', () => {
    const rendered = render(
      <ThemeProvider theme={lightTheme}>
        <Page3 />
      </ThemeProvider>,
    );
    expect(rendered.baseElement).toBeInTheDocument();
  });
});
