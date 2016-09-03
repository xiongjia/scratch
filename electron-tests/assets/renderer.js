'use strict';

(() => {
  const {remote} = require('electron');
  const {Menu, MenuItem} = remote;

  console.log('electron test!');

  /* button test */
  document.getElementById('btnTest').onclick = () => {
    console.log('btnTest');
  };

  /* menu template */
  const template = [{
    label: 'MainMenu',
    submenu: [{
      label: 'Test',
      accelerator: 'CmdOrCtrl+T',
      click (item, focusedWindow) {
        console.log('Test Menu Click');
      }
    }, {
      type: 'separator'
    }, {
      label: '&Test2',
      accelerator: 'CmdOrCtrl+A',
      click (item, focusedWindow) {
        console.log('Test2 Menu Click');
      }
    }]
  }];
  const mainMenu = Menu.buildFromTemplate(template);
  Menu.setApplicationMenu(mainMenu);

  /* context menu tests */
  const menu = new Menu();
  menu.append(new MenuItem({
    label: 'MenuItem1',
    click() {
      console.log('item 1 clicked');
    }
  }));
  menu.append(new MenuItem({type: 'separator' }));
  menu.append(new MenuItem({
    label: 'MenuItem2',
    type: 'checkbox',
    checked: true
  }));
  window.addEventListener('contextmenu', (e) => {
    console.log('ctx menu');
    e.preventDefault();
    menu.popup(remote.getCurrentWindow());
  }, false);

  document.addEventListener('dragover', (evt) => {
    evt.preventDefault();
    return false;
  }, false);
  document.addEventListener('drop', (evt) => {
    evt.preventDefault();
    return false;
  }, false);
})();
