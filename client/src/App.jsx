import './App.css'
import User from './components/user/User'
import Login from './components/login/Login'
import { AppContext } from './context/userContext'

const App = () => {

  return (
    <AppContext>
      <Login />
      <User />
    </AppContext>
  )
}

export default App
