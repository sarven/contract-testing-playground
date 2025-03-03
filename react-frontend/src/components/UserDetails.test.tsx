import React from 'react'
import { render, screen, waitFor } from '@testing-library/react'
import '@testing-library/jest-dom'
import UserDetails from './UserDetails'
import { jest } from '@jest/globals'

jest.mock('../api/userApi', () => {
    const { fakeGetUserDetails } = require('../api/fakeUserApi')
    return {
        getUserDetails: fakeGetUserDetails,
    }
})

describe('UserDetails component', () => {
    it('fetches and displays user details', async () => {
        render(<UserDetails userId={1} />)

        expect(screen.getByText('Loading...')).toBeInTheDocument()

        await waitFor(() => expect(screen.getByText('John Doe')).toBeInTheDocument())
        expect(screen.getByText('john.doe@example.com')).toBeInTheDocument()
    })

    it('displays an error message on failure', async () => {
        render(<UserDetails userId={2} />)

        expect(screen.getByText('Loading...')).toBeInTheDocument()

        await waitFor(() => expect(screen.getByText('Failed to fetch user details')).toBeInTheDocument())
    })
})
