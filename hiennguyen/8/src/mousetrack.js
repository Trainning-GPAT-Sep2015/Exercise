import React from 'react';
import ReactDOM from 'react-dom';

var Coords=React.createClass({
	getInitialState:function(){
		return {
			x:0,
			y:0
		};
	},
	showCoords:function(e){
		this.setState({
			x:e.clientX,
			y:e.clientY
		});
	},
	render:function(){
		return (
			<div onMouseMove={this.showCoords} style={{width:1000,height:700}}>
				<p>{this.state.x}</p>
				<p>{this.state.y}</p>
			</div>
			);
	}
});

ReactDOM.render(<Coords/>,document.getElementById('app'));