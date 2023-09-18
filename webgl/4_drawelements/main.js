const vertexShaderSource = `#version 300 es

layout(location=0) in float aSize;
layout(location=1) in vec4 aPos;
layout(location=2) in vec4 aColor;

out vec4 vColor;

void main() {
  vColor = aColor;
  gl_PointSize = aSize;
  gl_Position = aPos;
}`;

const aSizeLoc = 0;
const aPosLoc = 1;
const aColorLoc = 2;

const fragmentShaderSource = `#version 300 es

precision mediump float;

out vec4 fragColor;

in vec4 vColor;

void main() {
  fragColor = vColor;
}`;

const canvas = document.querySelector('canvas');
const gl = canvas.getContext('webgl2');

const program = gl.createProgram();

const vertexShader = gl.createShader(gl.VERTEX_SHADER);
gl.shaderSource(vertexShader, vertexShaderSource);
gl.compileShader(vertexShader);
gl.attachShader(program, vertexShader);

const fragmentShader = gl.createShader(gl.FRAGMENT_SHADER);
gl.shaderSource(fragmentShader, fragmentShaderSource);
gl.compileShader(fragmentShader);
gl.attachShader(program, fragmentShader);

gl.linkProgram(program);

if (!gl.getProgramParameter(program, gl.LINK_STATUS)) {
  console.log(gl.getShaderInfoLog(vertexShader));
  console.log(gl.getShaderInfoLog(fragmentShader));
}

gl.useProgram(program);

const elemVertexData = new Float32Array([
  0.00, 0.00,
  0.00, 1.00,
  0.95, 0.31,
  0.58, -.81,
  -.58, -.81,
  -.95, 0.31,
])

const elemIndiciesData = new Uint8Array([
  0,1,2,
  0,2,3,
  0,3,4,
  0,4,5,
  0,5,1,
])

const vertBuf = gl.createBuffer();
gl.bindBuffer(gl.ARRAY_BUFFER, vertBuf);
gl.bufferData(gl.ARRAY_BUFFER, elemVertexData, gl.STATIC_DRAW);

const indexBuf = gl.createBuffer();
gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuf,);
gl.bufferData(gl.ELEMENT_ARRAY_BUFFER, elemIndiciesData, gl.STATIC_DRAW);

gl.vertexAttribPointer(aPosLoc, 2, gl.FLOAT, false, 0, 0);

gl.vertexAttrib4f(aColorLoc, 1, 0, 0, 1);

gl.enableVertexAttribArray(aPosLoc);

gl.drawElements(gl.TRIANGLES, 15, gl.UNSIGNED_BYTE, 0);
