import React from 'react';
import ReactDOM from 'react-dom';
import { LINE,RECT, state, mouseDown, mouseUp, mouseMove, clearCanvas } from './store'; 

var Canvas = React.createClass({
	getInitialState(){
		return {};
	},
	mouseMove(e){
		mouseMove(e.pageX - this.refs.canvas.offsetLeft, e.pageY - this.refs.canvas.offsetTop);
	},
	componentDidMount(){
		state.context = this.refs.canvas.getContext("2d");
		console.log(this.refs.canvas.offsetLeft);
	},
	mouseDown(e){
		mouseDown(e.pageX - this.refs.canvas.offsetLeft, e.pageY - this.refs.canvas.offsetTop);
	},
	mouseUp(e){
		mouseUp(e.pageX - this.refs.canvas.offsetLeft, e.pageY - this.refs.canvas.offsetTop);
	},
	mouseLeave(){
		// console.log("mouse leave");
	},
	clearCanvas(){
		clearCanvas();
	},
	modeLINE(){
		state.mode = LINE;
	},
	modeRECT(){
		state.mode = RECT;
	},
	setColor(){
		state.color = this.refs.color.value;
		this.setState({});
	},
	render(){
		return (
			<div>
				<button onClick={this.clearCanvas}>CLEAR</button>
				<button onClick={this.modeLINE}>LINE</button>
				<button onClick={this.modeRECT}>RECT</button>
				<select onChange={this.setColor} ref="color">
					{state.listColor.map((c) => <option value={c}>{c.toUpperCase()}</option>)}
				</select> 				
				<div>
					<canvas ref='canvas' onMouseLeave={this.mouseLeave} onMouseMove={this.mouseMove} onMouseDown={this.mouseDown} onMouseUp={this.mouseUp} style={{border: "1px solid black", width:"500", height:"250"}}></canvas>
				</div>
			</div>
		);
	}
});
ReactDOM.render(<Canvas />, document.getElementById('app'));