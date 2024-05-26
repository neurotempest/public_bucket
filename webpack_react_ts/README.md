
* yarn deps

```
yarn init
yarn add react react-dom lodash
yarn add --dev webpack webpack-cli typescript ts-loader @types/react @types/react-dom @types/lodash
```

* change `packge.json`:
  * remove `entry: index.js`
  * add `private: true`

* add `dist/index.html`:
```
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Getting Started</title>
  </head>
  <body>
    <div id="root"></div>
    <script src="main.js"></script>
  </body>
</html>
```

* add `src/index.tsx`:
```
import React, { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import App from "./App";


const root = createRoot(document.getElementById("root"));
root.render(
  <StrictMode>
    <App />
  </StrictMode>
);
```

* add `src/App.tsx`:
```
import { useState } from 'react';

export default function Form() {
  const [value, setValue] = useState("Change me");

  function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
    setValue(event.currentTarget.value);
  }

  return (
    <>
      <input value={value} onChange={handleChange} />
      <p>Value: {value}</p>
    </>
  );
}
```

* add tsconfig.json:
```
import React, { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import App from "./App";


const root = createRoot(document.getElementById("root"));
root.render(
  <StrictMode>
    <App />
  </StrictMode>
);
```

* add `webpack.config.js`:
```
const path = require('path');

module.exports = {
  entry: './src/index.tsx',
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

* At this point running `yarn webpack` will bundle everyhting and running `open dist/index.html` will run website
