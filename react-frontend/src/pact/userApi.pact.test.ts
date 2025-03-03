import * as Pact from '@pact-foundation/pact';
import axios from "axios";
import { getUserDetails, User } from "./../api/userApi";
import path from "path";

const pact = new Pact.Pact({
    consumer: "ReactFrontend",
    provider: "Backend",
    port: 1234,
    host: '127.0.0.1',
    dir: path.resolve(process.cwd(), 'pacts'),
    log: path.resolve(process.cwd(), 'logs', 'pact.log'),
    logLevel: "debug",
});

describe("User API Contract Test", () => {
    beforeAll(() => pact.setup());
    afterAll(() => pact.finalize());

    it("should get user by ID", async () => {
        await pact.addInteraction({
            state: "User with ID 1 exists",
            uponReceiving: "a request for user 1",
            withRequest: {
                method: "GET",
                path: "/api/users/1",
            },
            willRespondWith: {
                status: 200,
                headers: { "Content-Type": "application/json" },
                body: Pact.Matchers.somethingLike({
                    id: 1,
                    name: "Alice",
                }),
            },
        });

        // Point Axios to the Pact mock server
        axios.defaults.baseURL = "http://localhost:1234";

        // Make a request and verify the response
        const response: User = await getUserDetails(1);
        expect(response).toEqual({
            id: 1,
            name: "Alice",
        });

        await pact.verify();
    });
});
