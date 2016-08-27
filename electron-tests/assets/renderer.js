'use strict';

(() => {
  console.log('electron test!');
  document.addEventListener('dragover', (evt) => {
    evt.preventDefault();
    return false;
  }, false);
  document.addEventListener('drop', (evt) => {
    evt.preventDefault();
    return false;
  }, false);
})();
