import nunjucks from "nunjucks";

console.log("build template");


nunjucks.configure({ autoescape: true });
console.log(nunjucks.renderString('Hello {{ username }}', { username: 'James' }));



