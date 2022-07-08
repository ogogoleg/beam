import React, { MouseEvent } from 'react';
import logo from './logo.svg';
import './App.css';

const buttonStyle = {
  width: 200
};

function handleClick(event: MouseEvent) {

const frame = document.getElementById('frame') as HTMLIFrameElement;

if (!frame) {
  // Handle case where iframe not found
  return;
}

const contentWindow = frame.contentWindow;
if (!contentWindow) {
  return;
}

const snippet = document.getElementById('snippet') as HTMLInputElement;
const data = snippet.value

const message = {
  type: 'SetMultiContentMessage',
  content: [
      {
          sdk: 'java',
          content: data,
      },
      {
          sdk: 'python',
          content: data,
      },
  ],
};

  contentWindow.postMessage(message, '*');
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <div>
          <button style={buttonStyle} onClick={handleClick}>Pass content to playground</button>
          <input id='snippet'></input>
        </div>

        <iframe id="frame" width="600" height="600" src="http://localhost:3002/embedded?enabled=true"></iframe>
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
