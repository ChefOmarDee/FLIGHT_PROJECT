// import './App.css';
import React, { useEffect, useState } from 'react';
import WebForm from './mods/WebForm';
import GetData from './mods/GetData';

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
    <React.Fragment>
      <WebForm/>
      <GetData/>
    </React.Fragment>
  );
}

export default App;