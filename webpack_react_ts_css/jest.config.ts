module.exports = {
  preset: "ts-jest",
  testEnvironment: "node",
  moduleNameMapper: {
    "\\.css$": "<rootDir>/__mocks__/style_mock.js",
  },
};
