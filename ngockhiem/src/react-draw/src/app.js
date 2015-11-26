import React from 'react';
import ReactDOM from 'react-dom';
import { FREE,LINE,RECT, state, mouseDown, mouseUp, mouseMove, clearCanvas } from './store'; 

var Canvas = React.createClass({
	getInitialState(){
		return { selected: "LINE" };
	},
	mouseMove(e){
		mouseMove(e.pageX - this.refs.canvas.offsetLeft, e.pageY - this.refs.canvas.offsetTop);
		this.setState({x:e.pageX, y:e.pageY});
	},
	componentDidMount(){
		state.context = this.refs.canvas.getContext("2d");
		this.refs.canvas.width = window.innerWidth - this.refs.canvas.offsetLeft;
		this.refs.canvas.height = window.innerHeight - this.refs.canvas.offsetTop;
	},
	mouseDown(e){
		mouseDown(e.pageX - this.refs.canvas.offsetLeft, e.pageY - this.refs.canvas.offsetTop);
	},
	mouseUp(e){
		mouseUp(e.pageX - this.refs.canvas.offsetLeft, e.pageY - this.refs.canvas.offsetTop);
	},
	mouseLeave(e){
		mouseUp(e.pageX - this.refs.canvas.offsetLeft, e.pageY - this.refs.canvas.offsetTop );
	},
	clearCanvas(){
		clearCanvas();
	},
	modeFREE(){
		state.mode = FREE;
		this.setState({ selected : "FREE" });
	},
	modeLINE(){
		state.mode = LINE;
		this.setState({ selected : "LINE" });
	},
	modeRECT(){
		state.mode = RECT;
		this.setState({ selected : "RECT" });
	},
	setColor(){
		state.color = this.refs.color.value;
		this.setState({});
	},
	render(){
		return (
			<div>
				<div id="control">
					<button onClick={this.clearCanvas}>CLEAR</button>
					<button onClick={this.modeFREE} id={this.state.selected === "FREE" ? "active" : ""} ref="btnF">FREE</button>
					<button onClick={this.modeLINE} id={this.state.selected === "LINE" ? "active" : ""} ref="btnL">LINE</button>
					<button onClick={this.modeRECT} id={this.state.selected === "RECT" ? "active" : ""} ref="btnR">RECT</button>
					<select onChange={this.setColor} ref="color">
						{state.listColor.map((c) => <option key={c} value={c}>{c.toUpperCase()}</option>)}
					</select> 				
					<span ref="coor"> X: {this.state.x},Y: {this.state.y}</span>
				</div>
				<div>
					<canvas id="canvas" ref='canvas' onMouseLeave={this.mouseLeave} onMouseMove={this.mouseMove} 
					onMouseDown={this.mouseDown} onMouseUp={this.mouseUp} style={{border:"1px solid black"}}
					></canvas>
				</div>
			</div>

		);
	}
});
ReactDOM.render(<Canvas />, document.getElementById('app'));