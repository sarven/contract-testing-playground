import React, { useEffect, useState } from 'react'
import { getUserDetails, User } from '../api/userApi'

interface UserDetailsProps {
    userId: number
}

const UserDetails: React.FC<UserDetailsProps> = ({ userId }) => {
    const [user, setUser] = useState<User | null>(null)
    const [loading, setLoading] = useState<boolean>(true)
    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        const fetchUserDetails = async () => {
            try {
                const userDetails = await getUserDetails(userId)
                setUser(userDetails)
            } catch (err) {
                setError('Failed to fetch user details')
            } finally {
                setLoading(false)
            }
        }

        fetchUserDetails()
    }, [userId])

    if (loading) return <div>Loading...</div>
    if (error) return <div>{error}</div>
    if (!user) return null

    return (
        <div>
            <h1>{user.name}</h1>
            <p>{user.email}</p>
        </div>
    )
}

export default UserDetails
