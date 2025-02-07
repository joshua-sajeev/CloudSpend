import { useState } from "react";
import InputComponent from "./components/InputComponent.jsx"
function App() {

    return (
        <div className="">
            <div className="text-center">
                <label>Cloud Spend</label>
            </div>
            <div className="border-3 rounded-lg border-indigo-700 p-5">
                <InputComponent/>
            </div>
        </div>
    );
}

export default App;
