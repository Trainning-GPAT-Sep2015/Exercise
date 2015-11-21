export const LINE = "LINE";
export const RECT = "RECT";
const MOUSE_UP = "UP";
const MOUSE_DOWN = "DOWN";
export const state = {
	context: null,
	listColor: ["black","red","purple","green","pink","yellow","blue"],
	color: "black",
	pointColor: [],
	mouse: MOUSE_UP,
	mode: LINE,
	pointMode: [],
	startX: [],
	startY: [],
	endX: [],
	endY: [],
	current: -1
};

export function mouseDown(x, y) {
	state.mouse = MOUSE_DOWN;
	state.current++;
	state.pointMode[state.current] = state.mode;
	state.pointColor[state.current] = state.color;
	state.startX[state.current] = x;
	state.startY[state.current] = y;
}

export function mouseUp(x, y) {
	if (state.mouse === MOUSE_DOWN) {
		state.endX[state.current] = x;
		state.endY[state.current] = y;
		state.mouse = MOUSE_UP;
	}
}

export function mouseMove(x, y) {
	if (state.mouse === MOUSE_DOWN) {
		state.endX[state.current] = x;
		state.endY[state.current] = y;
		draw();
	}
}

export function clearCanvas(){
	state.startX = [];
	state.startY = [];
	state.endX = [];
	state.endY = [];
	state.current = -1;
	state.context.clearRect(0, 0, state.context.canvas.width, state.context.canvas.height);
}


function draw() {
	console.log(state.pointMode);
	if (state.mode === LINE) {
		state.context.clearRect(0, 0, state.context.canvas.width, state.context.canvas.height);
		for (let i = 0; i<=state.current;i++){
			state.context.beginPath();
			if (state.pointMode[i] === LINE){
				state.context.moveTo(state.startX[i] / 1.65, state.startY[i] / 1.65);
				state.context.lineTo(state.endX[i] / 1.65, state.endY[i] / 1.65);
			} else {
				state.context.rect(state.startX[i] / 1.65, state.startY[i] / 1.65,state.endX[i] / 1.65 - state.startX[i] / 1.65, state.endY[i] / 1.65 - state.startY[i] / 1.65);
			}
			state.context.strokeStyle = state.pointColor[i];
			state.context.stroke();
		}				
		state.context.beginPath();
		state.context.moveTo(state.startX[state.current] / 1.65, state.startY[state.current] / 1.65);
		state.context.lineTo(state.endX[state.current] / 1.65, state.endY[state.current] / 1.65);
		state.context.strokeStyle = state.pointColor[state.current];
		state.context.stroke();
	}
	// console.log("--------------------------------------------------")
	// for (let i=0;i<=state.current;i++){
	// 	console.log("("+state.startX[i]+","+state.startY[i]+")" +" "+ "("+state.endX[i]+","+state.endY[i]+")");
	// }
	if (state.mode === RECT) {
		state.context.clearRect(0, 0, state.context.canvas.width, state.context.canvas.height);
		for (let i = 0; i<=state.current;i++){
			state.context.beginPath();
			if (state.pointMode[i] === LINE){
				state.context.moveTo(state.startX[i] / 1.65, state.startY[i] / 1.65);
				state.context.lineTo(state.endX[i] / 1.65, state.endY[i] / 1.65);
			} else  {
				state.context.rect(state.startX[i] / 1.65, state.startY[i] / 1.65,state.endX[i] / 1.65 - state.startX[i] / 1.65, state.endY[i] / 1.65 - state.startY[i] / 1.65);
			}
			state.context.strokeStyle = state.pointColor[i];
			state.context.stroke();
		}				
		state.context.beginPath();
		state.context.rect(state.startX[state.current] / 1.65, state.startY[state.current] / 1.65,state.endX[state.current] / 1.65 - state.startX[state.current] / 1.65, state.endY[state.current] / 1.65 - state.startY[state.current] / 1.65);
		state.context.strokeStyle = state.pointColor[state.current];
		state.context.stroke();
 
	}
}

function drawLINE(index) {
	state.context.beginPath();
	state.context.moveTo(state.startX[index] / 1.65, state.startY[index] / 1.65);
	state.context.lineTo(state.endX[index] / 1.65, state.endY[index] / 1.65);
	state.context.strokeStyle = state.pointColor[index];
	state.context.stroke();
}

function drawRECT(index){
	state.context.rect(state.startX[index] / 1.65, state.startY[index] / 1.65,state.endX[index] / 1.65 - state.startX[index] / 1.65, state.endY[index] / 1.65 - state.startY[index] / 1.65);
	state.context.strokeStyle = state.pointColor[index];
	state.context.stroke(); 
}