import React, { useEffect, useState } from 'react';

let GetData=()=>{
    let [obj, setObj] = useState([])
    function timeout(delay) {
      return new Promise( res => setTimeout(res, delay) );
  }
  async function getData() {
    await timeout(2000)
    const actualData = await fetch('http://127.0.0.1:8080/gettickets')
    let jsonActualData=await actualData.json()
    setObj(jsonActualData)
  }
    useEffect(() => {
      getData()
    }, [])

    //////////////////////////////////////////
    /////////////////////////////////////////
    return(
        <React.Fragment>
          <button onClick={getData}>
          {obj.sourcelocation}
          </button>
        </React.Fragment>
    )

}
export default GetData