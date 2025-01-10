import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import LoginForm from "./components/LoginForm.jsx";
import HeaderNavbar from "./components/HeaderNavbar.jsx";

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
        <HeaderNavbar />
       <LoginForm />
    </>
  )
}

export default App
