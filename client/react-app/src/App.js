import './App.css';
import React from 'react';
import axios from 'axios';
function App() {

  async function fetchData() {
    return axios.get("https://randomuser.me/api/")
      .then((response) => console.log(response.data));
  }
  let result = fetchData();
  return (
    <React.Fragment>
      <p>{result}</p>
    </React.Fragment>
  );
}

export default App;
