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

gl.vertexAttrib1f(aSizeLoc, 150.0);
gl.vertexAttrib4fv(aPosLoc, [0.5, -0.5, 0.0, 1.0]);
gl.vertexAttrib4fv(aColorLoc, [0.5, 0.5, 0.0, 1.0]);

const vertData = new Float32Array([
  -0.1,-0.5,  10, 0.1,0.5,0.1,
  -0.2,-0.4,  20, 0.2,0.4,0.5,
  -0.3,-0.3,  30, 0.3,0.3,0.2,
  -0.4,-0.2,  40, 0.4,0.2,0.3,
  -0.5,-0.1,  50, 0.5,0.1,0.4,
   0.1,-0.5,  80, 0.2,0.5,0.1,
   0.2,-0.2,  60, 0.3,0.3,0.5,
   0.3,-0.4,  30, 0.1,0.2,0.7,
   0.4,-0.1,  40, 0.4,0.2,0.1,
   0.5,-0.2,  20, 0.2,0.5,0.8,
])


const vertBuf = gl.createBuffer();
gl.bindBuffer(gl.ARRAY_BUFFER, vertBuf);
gl.bufferData(gl.ARRAY_BUFFER, vertData, gl.STATIC_DRAW);

gl.vertexAttribPointer(aPosLoc, 2, gl.FLOAT, false, 24, 0);
gl.vertexAttribPointer(aSizeLoc, 1, gl.FLOAT, false, 24, 8);
gl.vertexAttribPointer(aColorLoc, 3, gl.FLOAT, false, 24, 12);

gl.enableVertexAttribArray(aPosLoc);
gl.enableVertexAttribArray(aSizeLoc);
gl.enableVertexAttribArray(aColorLoc);

gl.drawArrays(gl.POINTS, 0, 10);
