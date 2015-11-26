import React from 'react';
import ReactDOM from 'react-dom';
import Graphic from './drawer';

var App=React.createClass({
  getInitialState:function(){
    return {
      mode:'line',
      color:'black'
    };
  },
  changeMode:function(e){
    this.setState({mode:e.target.value});
  },
  changeColor:function(e){
    this.setState({color:e.target.value});
  },
  clear:function(){
    this.refs.canvas.clear();
  },
  render:function(){
    return (
      <div>
        <div>
          <select defaultValue='line' onChange={this.changeMode}>
            <option value='line'>Line</option>
            <option value='rec'>Rectangular</option>
          </select>
          <select defaultValue='black' onChange={this.changeColor}>
            <option value='black'>Black</option>
            <option value='blue'>Blue</option>
            <option value='green'>Green</option>
            <option value='red'>Red</option>
            <option value='yellow'>Yellow</option>
          </select>
          <button ref='clear' onClick={this.clear}>Clear</button>
        </div>
        <div>
          <Graphic mode={this.state.mode} color={this.state.color} ref='canvas'/>
        </div>
      </div>);
  }
});

ReactDOM.render(<App />, document.getElementById('app'));