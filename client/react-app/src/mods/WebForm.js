import React,{useState, useEffect} from 'react'
import '../cssForMods/Form.css';
import USAStates from './USAStates'
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
                    <label >Enter Flight Origin City</label>
                    <br/>
                    <input type="text" className='input'></input>
                    <br/>
                    <label>Enter Flight Origin State</label>
                    <br/>
                    <USAStates/>
                    <br/>
                    <label >Enter Flight Destination City</label>
                    <br/>
                    <input type="text" className='input'></input>
                    <br/>
                    <label>Enter Flight Destination State</label>
                    <br/>
                    <USAStates/>
                    <br/>
                    <label >Enter Flight Date</label>
                    <br/>
                    <input type="date" className='input' defaultValue= {new Date().toISOString().split('T')[0].toString()} max={new Date(new Date().setFullYear(new Date().getFullYear() + 1)).toISOString().split('T')[0].toString()} min={new Date().toISOString().split('T')[0].toString()} required pattern="\d{2}-\d{2}-\d{4}"/>
                    <br/>
                    <input type="submit" className='submit'></input>
                </form>  
                <p className='failed-output'>
                    Could not find results, this could be because there are no direct flights available or because improper Spelling/Capitalization has been used
                </p>
                </React.Fragment>
                )
      }
        else if(isFormSent>=1){
            return (
                <React.Fragment >
                <form className='aftersubmit'  onSubmit={submitHandler}>
                    <label >Enter Flight Origin City</label>
                    <br/>
                    <input type="text" className='input'></input>
                    <br/>
                    <label>Enter Flight Origin State</label>
                    <br/>
                    <USAStates/>
                    <br/>
                    <label >Enter Flight Destination City</label>
                    <br/>
                    <input type="text" className='input'></input>
                    <br/>
                    <label>Enter Flight Destination State</label>
                    <br/>
                    <USAStates/>
                    <br/>
                    <label >Enter Flight Date</label>
                    <br/>
                    <input type="date" className='input' defaultValue= {new Date().toISOString().split('T')[0].toString()} max={new Date(new Date().setFullYear(new Date().getFullYear() + 1)).toISOString().split('T')[0].toString()} min={new Date().toISOString().split('T')[0].toString()} required pattern="\d{2}-\d{2}-\d{4}"/>
                    <br/>
                    <input type="submit" className='submit'></input>
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
            <label >Enter Flight Origin City</label>
            <br/>
            <input type="text" className='input'></input>
            <br/>
            <label>Enter Flight Origin State</label>
            <br/>
            <USAStates/>
            <br/>
            <label >Enter Flight Destination City</label>
            <br/>
            <input type="text" className='input'></input>
            <br/>
            <label>Enter Flight Destination State</label>
            <br/>
            <USAStates/>
            <br/>
            <label >Enter Flight Date</label>
            <br/>
            <input type="date" className='input' defaultValue= {new Date().toISOString().split('T')[0].toString()} max={new Date(new Date().setFullYear(new Date().getFullYear() + 1)).toISOString().split('T')[0].toString()} min={new Date().toISOString().split('T')[0].toString()} required pattern="\d{2}-\d{2}-\d{4}"/>
            <br/>
            <input type="submit" className='submit'></input>
        </form>  
    </React.Fragment>
    )
}

export default WebForm