import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'

createRoot(document.getElementById('root')).render(
    <StrictMode>
        <div className="flex h-screen bg-[#FAF5EF] items-center justify-center">
            <div className="border-5 border-purple-500  p-5 rounded-lg">
                <App />
            </div>
        </div>
    </StrictMode>,
)
