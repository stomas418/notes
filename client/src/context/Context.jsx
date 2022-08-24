import { useContext } from "react"
import { createContext, useState } from "react"

const UserContext = createContext()
const LogContext = createContext()

export const useUser = () => {
    return useContext(UserContext)
}

export const useLog = () => {
    return useContext(LogContext)
}

export const UserProvider = ({ children }) => {
    const [user, setUser] = useState("")
    return (
        <UserContext.Provider value={[user, setUser]}>
            {children}
        </UserContext.Provider>
    )
}

export const LogProvider = ({ children }) => {
    const [logStatus, setLogStatus] = useState(false)
    return (
        <LogContext.Provider value={[logStatus, setLogStatus]}>
            {children}
        </LogContext.Provider>
    )
}

export const AppContext = ({ children }) => {
    return (
        <UserProvider>
            <LogProvider>
                {children}
            </LogProvider>
        </UserProvider>
    )
}