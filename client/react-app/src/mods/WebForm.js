import React,{useState, useEffect} from 'react'
import '../cssForMods/Form.css';
import FormOptions from './FormOptions';
let WebForm =()=>{
    let [isFormSent,setIsFormSent]=useState(0)
    let [obj, setObj] = useState([])

    async function getData() {
        const actualData = await fetch('http://127.0.0.1:8080/gettickets')
        let jsonActualData=await actualData.json()
        await setObj(jsonActualData)
      }

    async function submitHandler(event){
        event.preventDefault();

        var formData={
            destinationlocation:(event.target[2].value+event.target[3].value),
            sourcedate:(event.target[4].value),
            sourcelocation:(event.target[0].value+event.target[1].value)
        }

        let submitForm = async(formData)=>{
            let userFlightData= await fetch('http://localhost:8080/tickets',{
                method:'POST',
                headers:{
                   'Content-Type':'application/json'
                },
                body: JSON.stringify(formData)
            })
            setIsFormSent(isFormSent+=1)
        }   
        await submitForm(formData)
        await getData()
        await console.log(obj)
    }
    useEffect(() => {
        getData()
      }, [])
      if(isFormSent>=1&&(obj.flightprice===0||obj.flighttime[0]==='E'||(obj.destinationlocation).length===2||(obj.sourcelocation).length===2)){
            return(
            <React.Fragment >
                <form className='aftersubmit'  onSubmit={submitHandler}>
                    <FormOptions/>
                </form>  
                <p className='failed-output'>
                    Could not find results, this could be because there are no direct flights available or because improper spelling/capitalization has been used
                </p>
                </React.Fragment>
                )
      }
        else if(isFormSent>=1){
            return (
                <React.Fragment >
                <form className='aftersubmit'  onSubmit={submitHandler}>
                <FormOptions/>
                </form>  
                <p className='output'>
                <span className='output-text'>
                <br/>
                    <p className='output-text'>
                            Origin Iata:    
                        {  obj.sourceiata}
                    </p>
                    <p className='output-text'>
                    Destination Iata:  
                        {obj.destiata}
                    </p>
                        <p className='output-text'>
                        Cheapest Airline:
                        {obj.flightairline}
                        </p>
                    <p className='output-text'>
                        Cheapest Flight Price:
                        {obj.flightprice}
                    </p>
                    <p className='output-text'>
                        Flight Time:
                        {obj.flighttime}
                    </p>
                </span>
                </p>
                </React.Fragment>
                )
    }

    return (
    <React.Fragment >
        <form className='form' onSubmit={submitHandler}>
            <FormOptions/>
        </form>  
    </React.Fragment>
    )
}

export default WebForm