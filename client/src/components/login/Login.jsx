import { useRef } from "react"
import { useState } from "react"
import { useLog, useUser } from "../../context/userContext"

const Login = () => {
    const usernameRef = useRef(null)
    const passwordRef = useRef(null)
    const [, setLogStatus] = useLog()
    const [, setUser] = useUser()
    const handleSubmit = async (event) => {
        event.preventDefault()
        const response = await fetch("https://tom-notes.herokuapp.com/login?username=" + usernameRef.current.value, {
            method: "post",
            credentials: "include",
            body: JSON.stringify({
                "username": usernameRef.current.value,
                "password": passwordRef.current.value
            })
        })

        switch (response.status) {
            case 202: setLogStatus(true);
                setUser(usernameRef.current.value);
                break;
            case 401: setLogStatus(true);
                setUser(usernameRef.current.value);
                break;
            default: setLogStatus(false);
                setUser("");
                break;
        }
    }

    return (
        <div>
            <h1>Login Form</h1>
            <form onSubmit={handleSubmit} name="login">
                <label>
                    Username:
                    <input
                        type="text"
                        ref={usernameRef}
                    />
                </label>
                <label>
                    Password:
                    <input
                        type="password"
                        ref={passwordRef}
                    />
                </label>
                <input className="submit" type="submit" value="Submit" />
            </form>
        </div>
    )
}

export default Login