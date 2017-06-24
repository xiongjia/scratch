'use strict';

import React from 'react';
import PropTypes from 'prop-types';

import * as misc from '../misc.js';
const logger = misc.getLogger('layout');

export default class Layout extends React.Component {
  static propTypes = {
    data: PropTypes.string.isRequired
  }

  constructor() {
    super();
    this.state = {
      testState: 'test state',
    };
  }

  shouldComponentUpdate () {
    return true;
  }

  render() {
    logger.debug('props: ', this.props);
    setTimeout(() => {
      logger.debug('updating state');
      this.setState({testState: '123'});
    }, 1000 * 5);
    return (
      <div>
        <h1>{'prop data: '}</h1>{this.props.data}
        <h1>{'state data:'}</h1>{this.state.testState}
      </div>);
  }
}
