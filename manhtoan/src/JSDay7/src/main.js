import React from 'react';
import ReactDOM from 'react-dom';
import state, {deleteItem, addItem, subscribe} from './state'

var TodoList = React.createClass({
  deleteItem: function(index){
    // this.props.deleteItem(index);
    deleteItem(index)
  },
  render: function() {
    var createItem = (itemText, index) => {
      return (
            <li key={index + itemText}>
                {itemText}
                &nbsp;
                (<a style={{color: '#600', cursor:'pointer'}} onClick={() => this.deleteItem(index)}>x</a>)
            </li>
        );
    };
    return <ul>{this.props.items.map(createItem)}</ul>;
  }
});
var TodoApp = React.createClass({
  getInitialState: function() {
    subscribe(() => {
        this.setState({items: state.items});
    });
    return {items: ['hello'], text: ''};
  },
  onChange: function(e) {
    this.setState({text: e.target.value});
  },
  handleSubmit: function(e) {
    e.preventDefault();
    // var nextItems = this.state.items.concat([this.state.text]);
    // var nextText = '';
    addItem(this.state.text);
    // this.setState({items: nextItems, text: nextText});
    this.setState({text: ''});
  },
  // deleteItem: function(index) {
  //   this.state.items.splice(index, 1);
  //   this.setState({ items: this.state.items });
  // },
  render: function() {
    return (
      <div>
        <h3>TODO (x)</h3>
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