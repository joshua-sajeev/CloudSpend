import { useState } from "react";

export default function FormComponent() {
    const [formData, setFormData] = useState({
        title: "",
        amount: "",
    });
    const [message, setMessage] = useState("");

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value,
        });
    };

    const fetchData = () => {
        fetch("http://localhost:8080/")
            .then(response => response.text())
            .then(data => setMessage(data))
            .catch(error => console.error("Error fetching data:", error));
    };

    const handleSubmit = async (event) => {
        event.preventDefault();

        const data = {
            title: formData.title,
            amount: parseFloat(formData.amount),
        };

        try {
            const response = await fetch("http://localhost:8080/", {
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
            console.log(result.message)
        } catch (error) {
            console.error("Error submitting data:", error);
            setMessage("Error submitting data.");
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Title</label>
                <input 
                    type="text" 
                    name="title" 
                    value={formData.title} 
                    onChange={handleChange} 
                    placeholder="Enter your title" 
                />
            </div>

            <div>
                <label>Amount:</label>
                <input 
                    type="number" 
                    name="amount" 
                    value={formData.amount} 
                    onChange={handleChange} 
                    placeholder="Enter amount" 
                />
            </div>

            <div className="text-center border-2 mt-4">
                <button type="submit">Submit</button>
            </div>

            {message && <p>{message}</p>}
        </form>
    );
}
