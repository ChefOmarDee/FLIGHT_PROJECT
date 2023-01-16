import React, { useEffect, useState } from 'react';

let GetData=()=>{
    let [obj, setObj] = useState([])
    useEffect(() => {
      async function getData() {
        const actualData = await fetch('http://127.0.0.1:8080/gettickets')
        let jsonActualData=await actualData.json()
        setObj(jsonActualData)
      }
      getData()
    }, [])
    console.log(obj)
    //////////////////////////////////////////
    /////////////////////////////////////////
    return(
        <React.Fragment>
        </React.Fragment>
    )

}
export default GetData