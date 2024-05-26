
* Webpack setup from https://webpack.js.org/guides/getting-started/

# Steps

## First: manual typescript builds which are bundled with webpack:

* yarn deps
```
yarn init
yarn add --dev webpack webpack-cli typescript
```

* add `dist/index.html`:
```
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Getting Started</title>
  </head>
  <body>
    <script src="main.js"></script>
  </body>
</html>
```

* add `src/index.ts`
```
import _ from 'lodash';

function component() {
  const element = document.createElement('div');

  element.innerHTML = _.join(['Hello', 'webpack'], ' ');

  return element;
}

document.body.appendChild(component());

```

* in `package.json`:
  * Add `private: true`
  * Remove `main: index.js`

* make typescript compiler config:
```
yarn tsc --init
```

* add lodash libs and types:
```
yarn add lodash
yarn add --dev @types/lodash
```

* At this point, we should be able to get everything running by first building the typescript into js, and then using webpack to bundle all the js:
```
yarn tsc
```
(should create a `./src/index.js` file)
```
yarn webpack
```
(should create a `./dist/main.js file)

  * Now opening `./dist/index.html` in a browser should show "Hello webpack"

## Second: webpack to build the ts and bundle the js all together

* yarn add ts loader:
```
yarn add --dev ts-loader
```

* set the following ts config:
```
{
  "compilerOptions": {
    "outDir": "./dist/",
    "noImplicitAny": true,
    "module": "es6",
    "target": "es5",
    "jsx": "react",
    "allowJs": true,
    "moduleResolution": "node"
  }
}
```

* Add webpack config:
```
const path = require('path');

module.exports = {
  entry: './src/index.ts',
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js'],
  },
  output: {
    filename: 'main.js',
    path: path.resolve(__dirname, 'dist'),
  },
};
```

* change the import statement in `./src/index.ts` (otherwise webpack wont compile):
```
import * as _ from 'lodash';
```

* now when we run webpack it will bundle everything into `./dist/main.js`
```
yarn webpack
```
    * Opening `./dist/index.html` will show "Hello webpack" again
