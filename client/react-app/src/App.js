import './App.css';
import React, { useEffect, useState } from 'react';
//import axios from 'axios';
function App() {
  let obj = {}
  let [arr, setArr] = useState([])
  var b;
  useEffect(() => {
    async function getData() {
      const response = await fetch(
        `http://127.0.0.1:8080/tickets`
      )
      let actualData = await response.json();
      //console.log(actualData[0])
      obj = {
        sourceLocation: actualData[0].sourcelocation,
        destinationLocation: actualData[0].destinationlocation,
        sourceDate: actualData[0].sourcedate,
        destinationDate: actualData[0].destinationdate,
        passengerCount: actualData[0].passengercount,
      }
      setArr([...arr, obj])
    }
    getData()
  }, [])
  console.log(arr[0])
  return (
    <React.Fragment>
      <p>afl123{console.log(arr[0])}</p>
    </React.Fragment>
  );
}

export default App;
