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
      <Container />
    </div>
  );
}

/*
// Constants
const interval = 10000;
const dataPath = '/data';
const containerId = 'grid';

// Fetch and parse sensor values from the server and display it in the container
function fetchAndParseSensorValues() {
  fetch(dataPath)
    .then(response => response.json())
    .then(data => {
      console.log(data);

      const container = document.getElementById(containerId);
      if (container == null) {
        console.error('Failed to get container element by id of', containerId);
        return;
      }

      const sensorValues = Array.from(data);
      sensorValues.forEach(value => {
        container.appendChild(
          <div>
            
          </div>
        )
      });
    })
    .catch(err => {
      // TODO: handle error
      console.error(err);
    });
  
}

// At a fixed interval, fetch data from the server
setInterval(() => {
  console.log('Fetching data...');
  fetchAndParseSensorValues();
}, interval);
*/