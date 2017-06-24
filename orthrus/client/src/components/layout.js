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
  }

  shouldComponentUpdate () {
    return false;
  }

  render() {
    logger.debug('props: ', this.props);
    return (<div><h1>{'state data: '}</h1>{this.props.data}</div>);
  }
}
