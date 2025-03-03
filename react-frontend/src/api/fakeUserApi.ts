import { User } from './userApi'

export const fakeGetUserDetails = async (userId: number): Promise<User> => {
    if (userId === 1) {
        return {
            id: 1,
            name: 'John Doe',
            email: 'john.doe@example.com',
        }
    } else {
        throw new Error('Failed to fetch user details')
    }
}
