'use strict';

import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as axios from 'axios';

import { getTestData, getTestDataErr } from '../actions/data-action.js';

import * as misc from '../misc.js';
const logger = misc.getLogger('layout');

@connect((store) => store.data)
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
    logger.debug('will mount');
    this.props.dispatch((dispatch) => {
      axios.get('/api/v1/test-data')
        .then(res => dispatch(getTestData(res.data)))
        .catch(err => dispatch(getTestDataErr(err)));
    });
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
