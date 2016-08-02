'use strict';

function funcA(delay, callback) {
  callback = callback || function () {};
  console.log('Begin function A');
  setTimeout(() => {
    console.log('End function A');
    callback();
  }, delay);
}

function funcB(delay, callback) {
  callback = callback || function () {};
  console.log('Begin function B');
  setTimeout(() => {
    console.log('End function B');
    callback();
  }, delay);
}

function funcC(delay, callback) {
  callback = callback || function () {};
  console.log('Begin function C');
  setTimeout(() => {
    console.log('End function C');
    callback();
  }, delay);
}

function funcFinal(err) {
  if (err) {
    console.log('Err: %s', err.toString());
  }
  console.log('Final');
}

(() => {
  new Promise((resolve, reject) => {
    funcA(1000, (err) => {
      if (err) {
        reject(err);
      }
      else {
        resolve();
      }
    });
  }).then(() => {
    return new Promise((resolve, reject) => {
      funcB(1000, (err) => {
        if (err) {
          reject(err);
        }
        else {
          resolve();
        }
      });
    });
  }).then(() => {
    return new Promise((resolve, reject) => {
      funcC(1000, (err) => {
        if (err) {
          reject(err);
        }
        else {
          resolve();
        }
      });
    });
  }).then(() => funcFinal(), (err) => funcFinal(err));
})();
