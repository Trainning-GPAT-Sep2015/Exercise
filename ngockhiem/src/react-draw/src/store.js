export const LINE = "LINE";
export const RECT = "RECT";
const MOUSE_UP = "UP";
const MOUSE_DOWN = "DOWN";
export const state = {
	context: null,
	listColor: ["black","red","purple","green","pink","yellow","blue"],
	color: "black",
	mouse: MOUSE_UP,
	mode: LINE,
	LINEstart:[],
	LINEend:[],
	LINEcolor:[],
	RECTstart:[],
	RECTend:[],
	RECTcolor:[]
};

export function mouseDown(x, y) {
	state.mouse = MOUSE_DOWN;
	if (state.mode === LINE){
		state.LINEstart.push([x,y]);
		state.LINEend.push([x,y]);
		state.LINEcolor.push(state.color);
	} else if (state.mode === RECT){
		state.RECTstart.push([x,y]);
		state.RECTend.push([x,y]);
		state.RECTcolor.push(state.color);
	}
}

export function mouseUp(x, y) {
	if (state.mouse === MOUSE_DOWN) {
		if (state.mode === LINE){
			state.LINEend.pop();
			state.LINEend.push([x,y]);
		} else if (state.mode === RECT){
			state.RECTend.pop();
			state.RECTend.push([x,y]);
		}

		state.mouse = MOUSE_UP;
	}
}

export function mouseMove(x, y) {
	if (state.mouse === MOUSE_DOWN) {
		if (state.mode === LINE){
			state.LINEend.pop();
			state.LINEend.push([x,y]);
		} else if (state.mode === RECT){
			state.RECTend.pop();
			state.RECTend.push([x,y]);
		}
		draw();
	}
}

export function clearCanvas(){
	state.LINEstart = [];
	state.LINEend = [];
	state.LINEcolor = [];
	state.RECTstart = [];
	state.RECTend	 = [];
	state.RECTcolor = [];
	state.context.clearRect(0, 0, state.context.canvas.width, state.context.canvas.height);
}

function drawLINE() {
	for ( let i = 0; i <= state.LINEstart.length-1; i++) {
		state.context.beginPath();
		state.context.moveTo(state.LINEstart[i][0],state.LINEstart[i][1]);
		state.context.lineTo(state.LINEend[i][0],state.LINEend[i][1]);
		state.context.strokeStyle = state.LINEcolor[i];
		state.context.stroke();
	};
}

function drawRECT(){
	for ( let i = 0; i <= state.RECTstart.length-1; i++) {
		state.context.beginPath();
		state.context.rect(state.RECTstart[i][0], state.RECTstart[i][1],state.RECTend[i][0] - state.RECTstart[i][0], state.RECTend[i][1] - state.RECTstart[i][1]);
		state.context.strokeStyle = state.RECTcolor[i];
		state.context.stroke();
	};
}

function draw(){
	if (state.mode === LINE) {
		state.context.clearRect(0, 0, state.context.canvas.width, state.context.canvas.height);
		drawLINE();
		drawRECT();

		state.context.beginPath();
		state.context.moveTo(state.LINEstart[state.LINEstart.length-1][0],state.LINEstart[state.LINEstart.length-1][1]);
		state.context.lineTo(state.LINEend[state.LINEend.length-1][0],state.LINEend[state.LINEend.length-1][1]);
		state.context.strokeStyle = state.LINEcolor[state.LINEstart.length-1];
		state.context.stroke();
	}
	if (state.mode === RECT) {
		state.context.clearRect(0, 0, state.context.canvas.width, state.context.canvas.height);
		drawLINE();
		drawRECT();

		const last = state.RECTstart.length - 1;
		state.context.beginPath();
		state.context.rect(state.RECTstart[last][0], state.RECTstart[last][1],state.RECTend[last][0] - state.RECTstart[last][0], state.RECTend[last][1] - state.RECTstart[last][1]);
		state.context.strokeStyle = state.RECTcolor[last];
		state.context.stroke();
	}
}