package hun;

import hun.beans.BeanOne;
import hun.beans.BeanTwo;
import hun.component.User;
import hun.core.BeanFactory;
import hun.core.ClassPathXmlApplicationContext;

public class HunIoc {
  public static void main(String[] args) {
    System.out.println("tests");
    final BeanFactory beanFactory = new ClassPathXmlApplicationContext("/app.xml");
    final BeanOne beanOne = (BeanOne)beanFactory.getBean("BeanOne");
    System.out.println("BeanOne: " + beanOne.getName());
    System.out.println("BeanOne (2): " + beanOne.getBeanTwo().getNum());

    final BeanTwo beanTwo = (BeanTwo)beanFactory.getBean("BeanTwo");
    System.out.println("BeanTwo: " + beanTwo.getNum());
    beanTwo.setNum(66);
    System.out.println("BeanTwo: " + beanTwo.getNum());
    System.out.println("BeanOne (2): " + beanOne.getBeanTwo().getNum());

    final User user = (User) beanFactory.getBean("User");
    user.addUser();

    final User userProxyBefore = (User) beanFactory.getBean("UserProxyBefore");
    userProxyBefore.addUser();

    final User userProxyAfter = (User) beanFactory.getBean("UserProxyAfter");
    userProxyAfter.addUser();
  }
}
