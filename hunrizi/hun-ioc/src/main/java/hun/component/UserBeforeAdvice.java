package hun.component;

import hun.aop.MethodBeforeAdvice;
import java.lang.reflect.Method;

public class UserBeforeAdvice extends MethodBeforeAdvice {
  @Override
  public void before(Method method, Object[] args, Object target) {
    System.out.println("before user dao");
  }
}
