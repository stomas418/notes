import Note from './Note'
import { useEffect } from "react"
import { useState } from "react"
import { useUser } from "../../context/Context"

const Notes = () => {
    const [user,] = useUser()
    const [notes, setNotes] = useState({})
    const url = import.meta.env.VITE_API_URL
    useEffect(() => {
        const getNotes = async () => {
            const response = await fetch(`${url}/${user}/notes/`, {
                method: "GET",
                credentials: "include"
            })
            const data = await response.json()
            setNotes(data)
        }
        getNotes()

    }, [])

    const LoadingNotes = () => {
        return (
            <div>
                Your notes are loading or there aren't any
            </div>
        )
    }

    return (
        notes ? (<div>
            {Object.entries(notes).map((note, index) => (
                <Note key={index} note={note[1]} />
            ))}
        </div>) :
            <LoadingNotes />

    )
}

export default Notes