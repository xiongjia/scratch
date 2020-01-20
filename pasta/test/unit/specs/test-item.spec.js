
describe('pool ', () => {
  test('create pool', async () => {
    await browser.get('https://cn.bing.com/');

    // inputting
    const questionInput = await browser.findElement(by.id('sb_form_q'));
    await questionInput.clear();
    await questionInput.sendKeys('webdriver');
    const value = await questionInput.getAttribute('value');
    console.log('value = %s', value);

    // submit
    const searchSubmit = await browser.findElement(by.id('sb_form_go'));
    await searchSubmit.click();

    await browser.findElement(by.id('b_results'));

    // questionInput.setText('webdriver');
    await new Promise((resolve, reject) => setTimeout(resolve, 3000));
  });
});
