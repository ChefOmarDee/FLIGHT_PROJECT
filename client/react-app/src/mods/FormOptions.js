import React from "react"
import USAStates from "./USAStates"
let FormOptions=()=>{
return(
    <React.Fragment>
            <label >
                    Enter Flight Origin City</label>
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
                    </React.Fragment>
)
}
export default FormOptions