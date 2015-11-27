import React from 'react';

export default React.createClass({
  getInitialState:function(){
    return {
      mousePressed:false,
      ctx:null,
      first:{
        x:0,
        y:0
      },
      coords:[],
      start:[],
      end:[],
      colors:[],
      modes:[]
    };
  },
  componentDidMount:function(){
    this.setState({ctx:this.refs.canvas.getContext('2d')});
  },
  mousedown:function(e){
    const d=this.refs.canvas;
    this.setState({
      mousePressed:true,
      first:{
        x:e.pageX-d.offsetLeft,
        y:e.pageY-d.offsetTop
      }
    });
  },
  mousemove:function(e){
    if (this.state.mousePressed){
      this.drawAll();
      const d=this.refs.canvas;
      this.draw(
        this.state.first.x,
        this.state.first.y,
        e.pageX-d.offsetLeft,
        e.pageY-d.offsetTop,
        this.props.color,
        this.props.mode
        );
      const coords=this.state.coords;
      coords.push({
        x:e.pageX-d.offsetLeft,
        y:e.pageY-d.offsetTop
      });
      this.setState({coords:coords});
    }
  },
  mouseup:function(e){
    this.setState({mousePressed:false});
    const coords=this.state.coords;
    const d=this.refs.canvas;
    coords.length=0;
    const start=this.state.start;
    start.push({
      x:this.state.first.x,
      y:this.state.first.y
    });
    const end=this.state.end;
    end.push({
        x:e.pageX-d.offsetLeft,
        y:e.pageY-d.offsetTop
      });
    const colors=this.state.colors;
    colors.push(this.props.color);
    const modes=this.state.modes;
    modes.push(this.props.mode);
    this.setState({
      coords:coords,
      start:start,
      end:end,
      colors:colors,
      modes:modes
    });
    this.drawAll();
  },
  draw:function(x1,y1,x2,y2,color,mode){
    var ctx=this.state.ctx;
    ctx.beginPath();
    if (mode==='line'){
      ctx.strokeStyle=color;
      ctx.lineWidth=2;
      ctx.lineJoin='round';
      ctx.moveTo(x1,y1);
      ctx.lineTo(x2,y2);
      ctx.stroke();
    } else if (mode==='rec'){
      ctx.strokeStyle=color;
      ctx.lineWidth=2;
      ctx.lineJoin='round';
      ctx.rect(x1,y1,x2-x1,y2-y1);
      ctx.stroke();
    }
  },
  drawAll:function(){
    this.clear();
    for (let i=0;i<this.state.end.length;i++){
      this.draw(
        this.state.start[i].x,
        this.state.start[i].y,
        this.state.end[i].x,
        this.state.end[i].y,
        this.state.colors[i],
        this.state.modes[i]
        );
    }
  },
  clear:function(){
    var ctx=this.state.ctx;
    ctx.clearRect(0,0,this.refs.canvas.width,this.refs.canvas.height);
  },
  clearAll:function(){
    this.setState({
      start:[],
      end:[],
      colors:[],
      modes:[]
    });
    this.clear();
  },
  render:function() {
    return <canvas width={1200} height={600} onMouseDown={this.mousedown} onMouseMove={this.mousemove}
      onMouseUp={this.mouseup} ref='canvas' style={{border:'solid black 2px'}}/>;
  }
});