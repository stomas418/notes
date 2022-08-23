import React, { useState } from 'react'

const Register = () => {
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const handleSubmit = async (event) => {
        event.preventDefault()
        const response = await fetch("https://tom-notes.herokuapp.com/register?username=" + username, {
            method: "post",
            credentials: "include",
            body: JSON.stringify({
                "username": username,
                "password": password
            })
        })
        console.log(response.status)
    }

    return (
        <div>
            <h1>Register Form</h1>
            <form onSubmit={handleSubmit}>
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

export default Register