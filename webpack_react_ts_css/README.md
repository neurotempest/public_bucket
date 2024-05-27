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
