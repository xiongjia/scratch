package hun.beans;

public class BeanOne {
  private String name;
  private BeanTwo beanTwo;

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public BeanTwo getBeanTwo() {
    return beanTwo;
  }

  public void setBeanTwo(BeanTwo beanTwo) {
    this.beanTwo = beanTwo;
  }
}
