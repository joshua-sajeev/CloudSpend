import { useState } from "react";
import "./App.css";

function App() {
    const [title, setTitle] = useState("");
    const [amount, setAmount] = useState("");
    const [message, setMessage] = useState("");

    const fetchData = () => {
        fetch("http://localhost:8080/")
            .then(response => response.text())
            .then(data => setMessage(data))
            .catch(error => console.error("Error fetching data:", error));
    };

    const handleSubmit = async (event) => {
        event.preventDefault();

        const data = { title, amount: parseFloat(amount) };

        try {
            const response = await fetch("http://localhost:8080/api/submit", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            });

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }

            const result = await response.json();
            setMessage(result.message); // Update message state
        } catch (error) {
            console.error("Error submitting data:", error);
            setMessage("Error submitting data.");
        }
    };

    return (
        <div>
            <button onClick={fetchData}>Click to fetch from Go server</button>

            {message && (
                <div>
                    <h2>Server Response:</h2>
                    <p>{message}</p>
                </div>
            )}

            <form onSubmit={handleSubmit}>
                <label>
                    Title:
                    <input
                        type="text"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                    />
                </label>
                <label>
                    Enter your amount:
                    <input
                        type="number"
                        value={amount}
                        onChange={(e) => setAmount(e.target.value)}
                    />
                </label>
                <input type="submit" value="Submit" />
            </form>
        </div>
    );
}

export default App;
