import './App.css'
import User from './components/user/User'
import { AppContext } from './context/Context'

const App = () => {
  return (
    <AppContext>
      <User />
    </AppContext>
  )
}

export default App
