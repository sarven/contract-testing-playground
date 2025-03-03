module.exports = {
    testEnvironment: 'jsdom',
    testMatch: ["**/?(*.)+(spec|test).[jt]s?(x)"],
    testPathIgnorePatterns: ["/node_modules/", "/src/pact/"],
};
