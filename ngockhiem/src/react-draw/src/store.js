export const LINE = "LINE";
export const RECT = "RECT";
export const FREE = "FREE";
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
	RECTcolor:[],
	FREElist: [],
	FREE: [],
	FREEcolor: []
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
	}	else if (state.mode === FREE){
		state.FREE.push([x,y]);
		state.FREEcolor.push(state.color);
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
		} else if (state.mode === FREE){
			state.FREE.push([x,y]);
			state.FREElist.push(state.FREE);			
			state.FREE = [];
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
		} else if (state.mode === FREE){
			state.FREE.push([x,y]);
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
	state.FREElist = [];
	state.FREE = [];
	state.FREEcolor = [];
	state.context.clearRect(0, 0, state.context.canvas.width, state.context.canvas.height);
}

function drawFREE() {
	for ( let i = 0; i < state.FREElist.length; i++) {
		state.context.beginPath();
		state.context.moveTo(state.FREElist[i][0][0],state.FREElist[i][0][1]);
		for ( let j = 1; j<state.FREElist[i].length;j++){
			state.context.lineTo(state.FREElist[i][j][0],state.FREElist[i][j][1]);
		}
		state.context.strokeStyle = state.FREEcolor[i];
		state.context.stroke();
	};
}

function drawLINE() {
	for ( let i = 0; i < state.LINEstart.length; i++) {
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
	if (state.mode === FREE) {
		state.context.clearRect(0, 0, state.context.canvas.width, state.context.canvas.height);
		drawFREE();
		drawLINE();
		drawRECT();

		state.context.beginPath();
		state.context.moveTo(state.FREE[0][0], state.FREE[0][1]);
		for ( let i = 1; i<state.FREE.length;i++){
			state.context.lineTo(state.FREE[i][0], state.FREE[i][1]);
		}
		state.context.strokeStyle = state.FREEcolor[state.FREEcolor.length - 1];
		state.context.stroke();
	}
	if (state.mode === LINE) {
		state.context.clearRect(0, 0, state.context.canvas.width, state.context.canvas.height);
		drawFREE();
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
		drawFREE();
		drawLINE();
		drawRECT();

		const last = state.RECTstart.length - 1;
		state.context.beginPath();
		state.context.rect(state.RECTstart[last][0], state.RECTstart[last][1],state.RECTend[last][0] - state.RECTstart[last][0], state.RECTend[last][1] - state.RECTstart[last][1]);
		state.context.strokeStyle = state.RECTcolor[last];
		state.context.stroke();
	}
}