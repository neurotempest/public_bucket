const vertexShaderSource = `#version 300 es

layout(location=0) in float aSize;
layout(location=1) in vec4 aPos;
layout(location=2) in vec2 aTexCoord;

out vec2 vTexCoord;

void main() {
  vTexCoord = aTexCoord;
  gl_PointSize = aSize;
  gl_Position = aPos;
}`;

const aSizeLoc = 0;
const aPosLoc = 1;
const aTexCoordLoc = 2;

const fragmentShaderSource = `#version 300 es

precision mediump float;

uniform sampler2D uImageSampler;
uniform sampler2D uPixelSampler;

in vec2 vTexCoord;

out vec4 fragColor;

void main() {
  fragColor = texture(uImageSampler, vTexCoord);
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

const vertexBufferData = new Float32Array([
  -0.9,-0.9,
  0,0.9,
  0.9,-0.9,
]);

const texCoordBufferData = new Float32Array([
  0,0,
  .5,1,
  1,0,
]);

const pixels = new Uint8Array([
  255,145,255,   230, 25, 15,    60,180,75,    225,255,25,
  155, 55,225,   130,125,175,    50,255,75,    225,255,25,
   55,145, 15,    30, 25, 35,    30,160,15,    225,255,25,
   25,215, 15,    30,225,275,    10, 80,15,    225,255,25,
]);


const vertBuf = gl.createBuffer();
gl.bindBuffer(gl.ARRAY_BUFFER, vertBuf);
gl.bufferData(gl.ARRAY_BUFFER, vertexBufferData, gl.STATIC_DRAW);
gl.vertexAttribPointer(aPosLoc, 2, gl.FLOAT, false, 0, 0);
gl.enableVertexAttribArray(aPosLoc);

const texCoordBuf = gl.createBuffer();
gl.bindBuffer(gl.ARRAY_BUFFER, texCoordBuf);
gl.bufferData(gl.ARRAY_BUFFER, texCoordBufferData, gl.STATIC_DRAW);
gl.vertexAttribPointer(aTexCoordLoc, 2, gl.FLOAT, false, 0, 0);
gl.enableVertexAttribArray(aTexCoordLoc);


const loadImage = () => new Promise(resolve => {
  const image = new Image();
  image.addEventListener('load', () => resolve(image));
  image.src = './image.png';
});

const run = async () => {
  const image = await loadImage();

  gl.pixelStorei(gl.UNPACK_FLIP_Y_WEBGL, true);

  const pixelTextureUnit = 0;
  const imageTextureUnit = 3;

  gl.uniform1i(gl.getUniformLocation(program, 'uPixelSampler'), pixelTextureUnit);
  gl.uniform1i(gl.getUniformLocation(program, 'uImageSampler'), imageTextureUnit);

  const pixelTexture = gl.createTexture();
  gl.activeTexture(gl.TEXTURE0 + pixelTextureUnit);
  gl.bindTexture(gl.TEXTURE_2D, pixelTexture);
  gl.texImage2D(gl.TEXTURE_2D, 0, gl.RGB, 4, 4, 0, gl.RGB, gl.UNSIGNED_BYTE, pixels);
  gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST);
  gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST);

  const imageTexture = gl.createTexture();
  gl.activeTexture(gl.TEXTURE0 + imageTextureUnit);
  gl.bindTexture(gl.TEXTURE_2D, imageTexture);
  gl.texImage2D(gl.TEXTURE_2D, 0, gl.RGB, 500, 300, 0, gl.RGB, gl.UNSIGNED_BYTE, image);

  gl.generateMipmap(gl.TEXTURE_2D);

  gl.drawArrays(gl.TRIANGLES, 0, 3);
};

run();
