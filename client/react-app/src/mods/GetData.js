import React, { useEffect, useState } from 'react';

let GetData=(props)=>{
    let [obj, setObj] = useState([])
    let [isFormSubmitted,setIsFormSubmitted]=useState(0)
    function timeout(delay) {
      return new Promise( res => setTimeout(res, delay) );
  }
  async function getData() {
    await timeout(6000)
    const actualData = await fetch('http://127.0.0.1:8080/gettickets')
    let jsonActualData=await actualData.json()
    setObj(jsonActualData)
    setIsFormSubmitted(isFormSubmitted+=1)
    await console.log(isFormSubmitted)

  }
    useEffect(() => {
      getData()
    }, [])

    if(isFormSubmitted>=2){
      console.log(obj)
      return(
        <React.Fragment>
        <button onClick={getData}>
        click here for your data
        </button>
        <p>
          {obj.sourcelocation}
        </p>
        </React.Fragment>
        
      )
    }
    // console.log(isFormSubmitted)
    //////////////////////////////////////////
    /////////////////////////////////////////
    return(
        <React.Fragment>
          <button onClick={getData}>
          click here for your data
          </button>
        </React.Fragment>
    )

}
export default GetData