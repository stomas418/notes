import { useState } from "react"
import { useUser } from "../../context/userContext"

const User = () => {
    const [user,] = useUser()
    const [data, setData] = useState("")
    const getUser = async () => {
        const response = await fetch("https://tom-notes.herokuapp.com/" + user, {
            credentials: "include",
            method: "GET",
            mode: "cors"
        })
        const data = await response.json()
        setData(data)
    }
    return (
        <div>
            <button onClick={getUser}>Click me for user info</button>
            <div>{data}</div>
        </div>
    )
}

export default User