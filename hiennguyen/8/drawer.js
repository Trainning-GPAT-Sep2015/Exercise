import React from 'react';

export default React.createClass({
  getInitialState:function(){
    return {
      firstX:0,
      firstY:0,
      mousePressed:false,
      ctx:null,
      coords:[]
    };
  },
  componentDidMount:function(){
    this.setState({ctx:this.refs.canvas.getContext('2d')});
  },
  mousedown:function(e){
    const d=this.refs.canvas;
    this.setState({mousePressed:true});
    this.setState({
      firstX:e.pageX-d.offsetLeft,
      firstY:e.pageY-d.offsetTop
    });
  },
  mousemove:function(e){
    if (this.state.mousePressed){
      this.clear();
      const d=this.refs.canvas;
      this.draw(e.pageX-d.offsetLeft,e.pageY-d.offsetTop);
      const coords=this.state.coords;
      coords.push({
        x:e.pageX-d.offsetLeft,
        y:e.pageY-d.offsetTop
      });
      this.setState({coords:coords});
    }
  },
  mouseup:function(e){
    this.clear();
    this.setState({mousePressed:false});
    const coords=this.state.coords;
    const d=this.refs.canvas;
    this.draw(coords[coords.length-1].x,coords[coords.length-1].y);
    coords.length=0;
    coords.push({
        x:e.pageX-d.offsetLeft,
        y:e.pageY-d.offsetTop
      });
    this.setState({coords:coords});
  },
  draw:function(x,y) {
    var ctx=this.state.ctx;
    ctx.beginPath();
    ctx.strokeStyle=this.props.color;
    ctx.lineWidth=2;
    ctx.lineJoin='round';
    ctx.moveTo(this.state.firstX,this.state.firstY);
    ctx.lineTo(x,y);
    ctx.closePath();
    ctx.stroke();
  },
  clear:function(){
    var ctx=this.state.ctx;
    ctx.clearRect(0,0,this.refs.canvas.width,this.refs.canvas.height);
  },
  render:function() {
    return <canvas width={1200} height={600} onMouseDown={this.mousedown} onMouseMove={this.mousemove}
      onMouseUp={this.mouseup} ref='canvas' style={{border:'solid black 2px'}}/>;
  }
});