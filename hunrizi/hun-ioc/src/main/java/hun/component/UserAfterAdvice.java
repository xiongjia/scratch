package hun.component;

import hun.aop.AfterReturningAdvice;

import java.lang.reflect.Method;

public class UserAfterAdvice extends AfterReturningAdvice {
  @Override
  public void afterReturning(Object returnValue,
                             Method method, Object[] args, Object target) throws Throwable {
    System.out.println("After method");
  }
}
