import { Pact } from "@pact-foundation/pact";
import axios from "axios";
import { getUser, User } from "./api/userApi";

const pact = new Pact({
    consumer: "ReactFrontend",
    provider: "PHPBackend",
    port: 1234,
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
                path: "/users/1",
            },
            willRespondWith: {
                status: 200,
                headers: { "Content-Type": "application/json" },
                body: { id: 1, name: "Alice" },
            },
        });

        // Point Axios to the Pact mock server
        axios.defaults.baseURL = "http://localhost:1234";

        // Make a request and verify the response
        const response: User = await getUser(1);
        expect(response).toEqual({ id: 1, name: "Alice" });

        await pact.verify();
    });
});
