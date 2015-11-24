import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import Tabber, { Tab } from './tabber';

class TabDemo extends Component {
  render() {
    return (
      <div>
        <Tabber defaultSelected="foo">
          <Tab label="foo" tabkey="foo" onClick={() => console.log('foo clicked')}>
            Foo

            Curabitur vitae leo enim. Praesent vitae lacus id turpis volutpat tincidunt vel in nisl. Suspendisse vel tellus nibh. Duis sit amet magna mattis, facilisis magna a, facilisis orci. Sed sit amet ex et est semper efficitur. Ut ullamcorper posuere vestibulum. Maecenas sed nulla auctor est imperdiet dignissim non ut nisl. Donec viverra gravida ante id fringilla. Morbi vitae commodo nunc. Nullam nec libero in felis rhoncus suscipit et tristique magna. Nam in ligula odio. Quisque vestibulum erat neque. Nullam ultrices, massa non fermentum euismod, est ligula interdum est, a rhoncus felis dui vel tellus. Suspendisse ut ultrices ante, at fringilla lacus. Nunc convallis, nisl vel consectetur vulputate, lectus neque dapibus odio, vitae congue sapien lacus quis urna.
          </Tab>
          <Tab label="bar" tabkey="bar" onClick={() => console.log('bar clicked')}>
            Bar

            Nunc maximus lectus nec dui facilisis laoreet. Donec eget magna sed justo luctus malesuada in et leo. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Cras leo metus, eleifend in sapien eu, varius ultrices dolor. Proin lacinia ex id augue pretium cursus. Donec eget tincidunt magna. Aliquam quis sapien eget nulla gravida dictum. Sed faucibus lacinia justo, nec tristique magna sagittis ut. Vivamus fringilla aliquam nibh nec egestas. Mauris nisl nisl, venenatis quis augue a, dapibus mattis massa. Phasellus ultrices nibh et dui cursus, vitae tincidunt dui porta. Mauris dignissim volutpat sem, nec blandit tellus posuere eget. Vestibulum aliquam risus odio, quis elementum
          </Tab>
          <Tab label="baz" tabkey="baz" onClick={() => console.log('baz clicked')}>
            Baz

            Nunc maximus lectus nec dui facilisis laoreet. Donec eget magna sed justo luctus malesuada in et leo. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Cras leo metus, eleifend in sapien eu, varius ultrices dolor. Proin lacinia ex id augue pretium cursus. Donec eget tincidunt magna. Aliquam quis sapien eget nulla gravida dictum. Sed faucibus lacinia justo, nec tristique magna sagittis ut. Vivamus fringilla aliquam nibh nec egestas. Mauris nisl nisl, venenatis quis augue a, dapibus mattis massa. Phasellus ultrices nibh et dui cursus, vitae tincidunt dui porta. Mauris dignissim volutpat sem, nec blandit tellus posuere eget. Vestibulum aliquam risus odio, quis elementum
          </Tab>
        </Tabber>
      </div>
    );
  }
}

ReactDOM.render(<TabDemo/>, document.getElementById('app'));