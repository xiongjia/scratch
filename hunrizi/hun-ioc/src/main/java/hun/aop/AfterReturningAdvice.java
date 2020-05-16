package hun.aop;

import java.lang.reflect.Method;
import net.sf.cglib.proxy.MethodInterceptor;
import net.sf.cglib.proxy.MethodProxy;

public abstract class AfterReturningAdvice implements MethodInterceptor {

  @Override
  public Object intercept(Object obj, Method method,
                          Object[] args, MethodProxy proxy) throws Throwable {
    Object result = proxy == null ? method.invoke(obj, args) : proxy.invokeSuper(obj, args);
    afterReturning(result, method, args, obj);
    return result;
  }

  public abstract void afterReturning(Object returnValue,
                                      Method method, Object[] args, Object target) throws Throwable;
}
