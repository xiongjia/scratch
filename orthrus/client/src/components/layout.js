'use strict';

import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { incData } from '../actions/data-action.js';

import * as misc from '../misc.js';
const logger = misc.getLogger('layout');

@connect((store) => {
  return {
    data: store.data
  };
})

export default class Layout extends React.Component {
  static propTypes = {
    data: PropTypes.number.isRequired,
    desc: PropTypes.string.isRequired,
    dispatch: PropTypes.func.isRequired
  }

  constructor() {
    super();
  }

  componentWillMount() {
    setTimeout(() => {
      this.props.dispatch(incData());
    }, 1000 * 5);
  }

  shouldComponentUpdate () {
    return true;
  }

  render() {
    const { data, desc } = this.props;
    logger.debug('props: ', this.props);
    return (
      <div>
        <h1>{'prop data: '}</h1>{data}
        <h1>{'prop desc:'}</h1>{desc}
      </div>);
  }
}
