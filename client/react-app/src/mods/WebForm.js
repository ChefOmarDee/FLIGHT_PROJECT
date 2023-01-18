import React,{useState, useEffect} from 'react'
// import '../cssForMods/Form.css';
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

        let formData={
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
  
        if(isFormSent>=1){
            return (
                <React.Fragment>
                    <form onSubmit={submitHandler}>
                        <label >Enter Flight Origin City</label>
                        <input type="text"></input>
                        <label>Enter Flight Origin State</label>
                        <USAStates/>
                        <label >Enter Flight Destination City</label>
                        <input type="text"></input>
                        <label>Enter Flight Destination State</label>
                        <USAStates/>
                        <label >Enter Flight Date</label>
                        <input type="date" defaultValue= {new Date().toISOString().split('T')[0].toString()} max={new Date(new Date().setFullYear(new Date().getFullYear() + 1)).toISOString().split('T')[0].toString()} min={new Date().toISOString().split('T')[0].toString()} required pattern="\d{2}-\d{2}-\d{4}"/>
                        <input type="submit"></input>
                        </form>  
                        <p>
                        {obj.sourcelocation}
                        <br/>
                        {obj.flightprice}
                        <br/>
                        {obj.flightairline}
                        </p>
                </React.Fragment>
                )
    }

    return (
    <React.Fragment>
        <form Class='form'onSubmit={submitHandler}>
            <label >Enter Flight Origin City</label>
            <input type="text"></input>
            <label>Enter Flight Origin State</label>
            <USAStates/>
            <label >Enter Flight Destination City</label>
            <input type="text"></input>
            <label>Enter Flight Destination State</label>
            <USAStates/>
            <label >Enter Flight Date</label>
            <input type="date" defaultValue= {new Date().toISOString().split('T')[0].toString()} max={new Date(new Date().setFullYear(new Date().getFullYear() + 1)).toISOString().split('T')[0].toString()} min={new Date().toISOString().split('T')[0].toString()} required pattern="\d{2}-\d{2}-\d{4}"/>
            <input type="submit"></input>
        </form>  
    </React.Fragment>
    )
}

export default WebForm