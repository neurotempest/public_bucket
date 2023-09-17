const vertexShaderSource = `#version 300 es

uniform float uPointSize;
uniform vec2 uPosition;

void main() {
  gl_PointSize = uPointSize;
  gl_Position = vec4(uPosition, 0.0, 1.0);
}`;

const fragmentShaderSource = `#version 300 es

precision mediump float;

out vec4 fragColor;

uniform int uIndex;
uniform vec4 uColors[3];

void main() {
  fragColor = uColors[uIndex];
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

const uPositionLoc = gl.getUniformLocation(program, 'uPosition');
const uPointSizeLoc = gl.getUniformLocation(program, 'uPointSize');
const uIndexLoc = gl.getUniformLocation(program, 'uIndex');
const uColorsLoc = gl.getUniformLocation(program, 'uColors');

gl.uniform4fv(uColorsLoc, [
  1.0, 0.0, 0.0, 1.0, // red
  0.0, 0.0, 1.0, 1.0, // blue
  0.0, 1.0, 0.0, 1.0, // green
]);


gl.uniform1i(uIndexLoc, 0);
gl.uniform1f(uPointSizeLoc, 100);
gl.uniform2f(uPositionLoc, 0.2, -0.6);

gl.drawArrays(gl.POINTS, 0, 1);

gl.uniform1i(uIndexLoc, 1);
gl.uniform1f(uPointSizeLoc, 30);
gl.uniform2f(uPositionLoc, -0.8, -0.1);

gl.drawArrays(gl.POINTS, 0, 1);

gl.uniform1i(uIndexLoc, 2);
gl.uniform1f(uPointSizeLoc, 45);
gl.uniform2f(uPositionLoc, 0.2, 0.5);

gl.drawArrays(gl.POINTS, 0, 1);
