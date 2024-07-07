# overview
This is a more _complete_ example setup of a webpack react project with css and some other local typescript libs which will be imported into the react App

# steps

* yarn deps
```
yarn init
yarn add react react-dom lodash
yarn add --dev webpack webpack-cli typescript ts-loader @types/react @types/react-dom @types/lodash html-webpack-plugin css-loader style-loader jest ts-jest @types/jest ts-node
```

* add `.gitignore`:
```
node_modules/
dist/
```

* change `packge.json`:
  * remove `entry: index.js`
  * add `private: true`
  * add:
```
  "scripts": {
    "build": "webpack",
    "test": "jest"
  },
```

* add tsconfig.json:
* add webpack.config.js
* add `src/index.html`:
* add `src/index.tsx`:
* add `src/someLib.ts`:
* add `test/someLib.spec.ts`:
* add `jest.config.ts`:


### Adding support for testing react components:

* yarn add react testing lib: `yarn add @testing-library/react --dev`

* yarn add some other libs required by jest:
```
yarn add @testing-library/dom jest-environment-jsdom --dev
```

* At the top of all test files which will test a react component we need to change the jest test enviroment to jsdom (i.e. browser-like test env)
    * See: https://stackoverflow.com/questions/69227566/consider-using-the-jsdom-test-environment
    * And https://jestjs.io/docs/configuration#testenvironment-string
  * Otherwise the test will fail. (THis can also be done for the whole project in the jest config - however not if the testshave a mix of UI and non-UI test files
  * To do this add the following doc-block as a head to each file:
```
/**
 * @jest-environment jsdom
 */

```

* Add an actual test file - e.g. `./test/App.spec.tsx`
```
/**
 * @jest-environment jsdom
 */
import { render, screen } from '@testing-library/react';


import Form from "../src/App"

describe("App", () => {
  it("does something", () => {
    render(<Form />)
    screen.debug()
  })
})
```

* Note on suporting `css` imports;
  * Webpack (specifically css-loader/ts-loader/babel) passes the ts+css to make valid ts/js
  * We need to mock this behaviour in jest
  * One solution is just to ~ignore~ nullify all css imports:
    * Add a module overide to jest config:
```
  moduleNameMapper: {
    "\\.css$": "<rootDir>/__mocks__/style_mock.js",
  },
```
    * Define the `style_mock.js` file:
```
module.exports = {};
```
  * see https://stackoverflow.com/questions/39641068/jest-trying-to-parse-css
