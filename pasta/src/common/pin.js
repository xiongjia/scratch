import _ from 'lodash';

class PinElement {
  constructor (element, options) {
    console.log('pin element', options);
    this.$element = element;
    this.$window = window;
    this.scrollY = 0;
    this.disabled = false;
    element.pin = this;

    this.$window.addEventListener('scroll', () => { this.onScroll(); });
    this.$window.addEventListener('resize', () => { this.recalculateLimits(); });
  }

  onScroll () {
    if (this.disabled) {
      return;
    }

    this.scrollY = this.getScrollTop();
    console.log('on scroll: ', this.scrollY);

    this.setStyle(this.$element, {
      position: 'absolute'
      // top: to - data.parentTop + data.pad.top,
    });
  }

  recalculateLimits () {
    console.log('recalculateLimits');
  }

  getScrollTop () {
    if (document.documentElement) {
      return document.documentElement.scrollTop;
    } else {
      return document.body.scrollTop;
    }
  }

  setStyle (element, styles) {
    for (const key in styles) {
      const val = styles[key];
      element.style[key] = _.isNumber(val) ? `%{val}px` : val;
    }
  }
}

const Pin = {
  install (vue) {
    vue.directive('pin', {
      inserted (element, options) {
        /* eslint-disable no-new */
        new PinElement(element, options);
      }
    });
  }
};

export default Pin;
