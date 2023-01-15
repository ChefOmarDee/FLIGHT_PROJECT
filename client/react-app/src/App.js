// import './App.css';
// import React, { useEffect, useState } from 'react';
import WebForm from './mods/WebForm';

function App() {

  // let [obj, setObj] = useState([])
  // useEffect(() => {
  //   async function getData() {
  //     const actualData = await fetch(
  //       `http://127.0.0.1:8080/tickets`
  //     ).then(response => response.json())
  //     setObj(JSON.stringify(actualData))
  //   }
  //   getData()
  // }, [])

  // console.log(obj)
  return (
    // <React.Fragment>
    // { <p>{obj}</p> }
      <WebForm>
      </WebForm>
    // </React.Fragment>
  );
}

export default App;