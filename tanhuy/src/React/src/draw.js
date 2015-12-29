import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import {LINE, RECT, state,mouseDown,mouseUp,mouseMove,clearCanvas} from './draw_logic';

var Canvas = React.createClass({
	getInitialState(){
		return {};
	},
	componentDidMount(){
		state.context = this.refs.canvas.getContext("2d");
		console.log(this.refs.canvas.getContext('2d'));
	},
	mouseMove(e){
		mouseMove(e.pageX - this.refs.canvas.offsetLeft, e.pageY-this.refs.canvas.offsetTop);
	},

	// mouseDidMount(){
	// 	state.context = this.refs.canvas.getContext("2d");
	// },

	mouseDown(e){
		mouseDown(e.pageX - this.refs.canvas.offsetLeft, e.pageY-this.refs.canvas.offsetTop);
	},

	mouseUp(e){
		mouseUp(e.pageX - this.refs.canvas.offsetLeft, e.pageY-this.refs.canvas.offsetTop);
	},

	clearCanvas(){
		clearCanvas();
	},

	modeRect(){
		state.mode=RECT;
	},

	modeLine(){
		state.mode=LINE;
	},

	setColor(){
		state.color = this.refs.color.value;
		this.setState({});
	},

	render(){
		return(
			<div>
				<button onClick={this.modeLine}>Draw Line</button>
				<button onClick={this.modeRect}>Draw Rectangle</button>
				<button onClick={this.clearCanvas}>Clear</button>
				<select onChange ={this.setColor} ref='color'>
					{state.colors.map((c)=><option value={c}>{c.toUpperCase()}</option>)}
				</select>
				<div>
					<canvas ref='canvas' onMouseMove={this.mouseMove} onMouseUp={this.mouseUp} onMouseDown={this.mouseDown} style={{width:"500", height:"250", border:"1px solid black"}}></canvas>
				</div>
			</div>

		);
	}

});

ReactDOM.render(<Canvas />, document.getElementById('app'));