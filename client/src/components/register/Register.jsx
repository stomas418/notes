import React, { useState } from 'react'

const Register = () => {
    const usernameRef = useRef()
    const passwordRef = useRef()
    const url = import.meta.env.VITE_API_URL
    const handleSubmit = async (event) => {
        event.preventDefault()
        const response = await fetch(`${url}/register?username=${usernameRef.current.value}`, {
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

export default Register