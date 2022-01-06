import React from 'react';
import { Container } from './Container';

export const App = (): JSX.Element => {
  return (
    <div
      style={{
        width: '100vw',
        height: '100vh',
        display: 'grid',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'black',
        color: 'white',
        fontFamily: 'Arial, Helvetica, sans-serif',
      }}
    >
      <Container url="/data-test" />
    </div>
  );
}
