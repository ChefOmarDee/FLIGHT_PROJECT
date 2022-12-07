import './App.css';
import React from 'react';
import axios from 'axios';
function App() {
  axios
    .get("http://127.0.0.1:8080/tickets")
    .then((res) => console.log(res))
    .catch((err) => console.log(err));
  return (
    <React.Fragment>
      <p>jojo</p>
    </React.Fragment>
  );
}

export default App;
