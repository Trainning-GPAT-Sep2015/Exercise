var React = require('react');
var ReactDOM= require('react-dom');
import state, {deleteItem, addItem,subscribe} from './state';
var TodoList = React.createClass({
  // deleteItem(index){
  // 	deleteItem(index);
  // },
  render: function() {
    var createItem = (itemText, index)=> {
      return (
      	<li key={index + itemText}>
      		{itemText}
      		(<a style={{ color:'#600',cursor:'pointer'}}, onClick={()=>deleteItem(index)}>x</a>)
      	</li>
      	);
    };
    return <ul>{this.props.items.map(createItem)}</ul>;
  }
});
var TodoApp = React.createClass({
  getInitialState: function() {
  	subscribe(()=>{
  		this.setState({items:state.items})
  	})
    return {items: [], text: ''};
  },
  onChange: function(e) {
    this.setState({text: e.target.value});
  },
  handleSubmit: function(e) {
    e.preventDefault();
    addItem(this.state.text);
    this.setState({text:''})
  },
  render: function() {
    return (
      <div>
        <h3>TODO</h3>
        <TodoList items={this.state.items}/>
        <form onSubmit={this.handleSubmit}>
          <input onChange={this.onChange} value={this.state.text} />
          <button>{'Add #' + (this.state.items.length + 1)}</button>
        </form>
      </div>
    );
  }
});

ReactDOM.render(<TodoApp />, document.getElementById('app'));
