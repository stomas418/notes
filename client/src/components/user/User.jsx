import { useUser } from "../../context/Context"
import Notes from '../notes/Notes'

const User = () => {
    const [user, setUser] = useUser()
    //TODO: create interface and logic to change username and password

    const changeUsername = async () => {
        setUser("test")
    }

    const changePassword = async () => {
        setUser("testn't")
    }

    return (
        <div id="user-page">
            <h1>Welcome {user}!</h1>
            <div id="settings-tab">
                <div className="change">
                    <span>Your current username: {user}</span>
                    <button onClick={changeUsername}>Edit</button>
                </div>
                <div className="change">
                    <span id="password">Your current password: {user}</span>
                    <button onClick={changePassword}>Edit</button>
                </div>
            </div>
            <div id="notes">
                <Notes />
            </div>
        </div>
    )
}

export default User