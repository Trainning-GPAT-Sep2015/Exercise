import React, {Component} from 'react';
export const LINE="LINE";
export const RECT ="RECT";
const MOUSE_UP = "UP";
const MOUSE_DOWN="DOWN";

export const state={
	context:null,
	colors:["black","red","yellow","pink","purple","green"],
	color:"black",
	pointColor:[],
	mouse:MOUSE_UP,
	mode:LINE,
	pointMode:[],
	beginX:[],
	beginY:[],
	endX:[],
	endY:[],
	current:-1
};

//handle mouse down event
export function mouseDown(x,y){
	state.mouse=MOUSE_DOWN;
	state.current++;
	state.pointMode[state.current]=state.mode;
	state.pointColor[state.current]=state.color;
	state.beginX[state.current]=x;
	state.beginY[state.current]=y;
}

//handle mouse up event
export function mouseUp(x,y){
	if(state.mouse===MOUSE_DOWN){
		state.mouse=MOUSE_UP;
		state.endX[state.current]=x;
		state.endY[state.current]=y;
	}
}

//handle mouse move event
export function mouseMove(x,y){
	if (state.mouse===MOUSE_DOWN) {
		state.endX[state.current]=x;
		state.endY[state.current]=y;
		draw();
	};
}

//clear the canvas
export function clearCanvas(){
	state.beginX=[];
	state.beginY=[];
	state.endX=[];
	state.endY=[];
	state.current=-1;
	state.context.clearRect(0,0,state.context.canvas.width,state.context.canvas.height);
}

//do draw
function draw(){
	if (state.mode===LINE) {
		state.context.clearRect(0,0,state.context.canvas.width,state.context.canvas.height);
		for(var i=0;i<=state.current;i++){
			state.context.beginPath();
			if(state.pointMode[i]===LINE){
				state.context.moveTo(state.beginX[i]/1.65,state.beginY[i]/1.65);
				state.context.lineTo(state.endX[i]/1.65,state.endY[i]/1.65);
			}else{
				state.context.rect(state.beginX[i]/1.65 , state.beginY[i]/1.65,state.endX[i]/1.65- state.beginX[i]/1.65, state.endY[i]/1.65 - state.beginY[i]/1.65 );
			}
			state.context.strokeStyle= state.pointColor[i];
			state.context.stroke();
		}
		state.context.beginPath();
		state.context.moveTo(state.beginX[i]/1.65,state.beginY[i]/1.65);
		state.context.lineTo(state.endX[i]/1.65,state.endY[i]/1.65);
		state.context.strokeStyle=state.pointColor[state.current];
		state.context.stroke();
	}else{
		// console.log(state.pointMode);
		state.context.clearRect(0,0,state.context.canvas.width,state.context.canvas.height);
		for(var i=0;i<=state.current;i++){
			state.context.beginPath();
			if(state.pointMode[i]===LINE){
				state.context.moveTo(state.beginX[i]/1.65,state.beginY[i]/1.65);
				state.context.lineTo(state.endX[i]/1.65,state.endY[i]/1.65);
			}else{
				state.context.rect(state.beginX[i]/1.65 , state.beginY[i]/1.65,state.endX[i]/1.65- state.beginX[i]/1.65, state.endY[i]/1.65 - state.beginY[i]/1.65 );
			}
			state.context.strokeStyle= state.pointColor[i];
			state.context.stroke();
		}
		state.context.beginPath();
		state.context.rect(state.beginX[state.current]/1.65 , state.beginY[state.current]/1.65,state.endX[state.current]/1.65- state.beginX[state.current]/1.65, state.endY[state.current]/1.65 - state.beginY[state.current]/1.65 );
		state.context.strokeStyle = state.pointColor[state.current];
		state.context.stroke();
	}
}