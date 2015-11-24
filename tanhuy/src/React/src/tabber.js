import React, { Component } from 'react';

export default React.createClass({
  getInitialState() {
    return {
      selected: this.props.defaultSelected
    };
  },
  selectTab(tabkey) {
    this.setState({ selected: tabkey });
  },
  render() {
    const { children } = this.props;
    const { selected } = this.state;
    console.log('children', children);

    return (
      <div>
        {
          React.Children.map(children, (child, index) => (
            <span>
              { index ? ' | ' : '' }
              <a className={ selected === child.props.tabkey ? 'selected' : '' }
                onClick={() => this.selectTab(child.props.tabkey)}
                >{child.props.label}</a>
            </span>
          ))
        }
        { 
          React.Children.map(children, (child) => 
            React.cloneElement(child, { __selected: selected }))
        }
      </div>
    );
  }
})

export const Tab = React.createClass({
  render() {
    const { children, tabkey, __selected } = this.props;

    return (
      <div className={ __selected === tabkey ? '' : 'hidden' }>
        { children }
      </div>
    );
  }
});