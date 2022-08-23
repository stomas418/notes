import { useState } from "react"

const Login = () => {
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const handleSubmit = async (event) => {
        event.preventDefault()
        const response = await fetch("https://tom-notes.herokuapp.com/login?username=" + username, {
            method: "post",
            credentials: "include",
            body: JSON.stringify({
                "username": username,
                "password": password
            })
        })
    }

    return (
        <div>
            <h1>Login Form</h1>
            <form onSubmit={handleSubmit} name="login">
                <label>
                    Username:
                    <input
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </label>
                <label>
                    Password:
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </label>
                <input className="submit" type="submit" value="Submit" />
            </form>
        </div>
    )
}

export default Login