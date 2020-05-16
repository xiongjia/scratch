package hun.core;

import java.util.HashMap;
import java.util.Map;

public class ClassPathXmlApplicationContext implements BeanFactory {
  private Map<String, Object> context = new HashMap<>();

  public ClassPathXmlApplicationContext(String path) {
  }

  @Override
  public Object getBean(String name) {
    return null;
  }
}
