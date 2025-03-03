import axios from 'axios'

export interface User {
    id: number
    name: string
    email: string
}

export const getUserDetails = async (userId: number): Promise<User> => {
    const response = await axios.get<User>(`/api/users/${userId}`)
    return response.data
}
