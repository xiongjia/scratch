package hun.aop;

import java.lang.reflect.Method;
import net.sf.cglib.proxy.MethodInterceptor;
import net.sf.cglib.proxy.MethodProxy;

public abstract class MethodBeforeAdvice implements MethodInterceptor {
  @Override
  public Object intercept(Object obj, Method method, Object[] args, MethodProxy proxy) throws Throwable {
    before(method, args, obj);
    return proxy == null ? method.invoke(obj, args) : proxy.invokeSuper(obj, args);
  }

  public abstract void before(Method method, Object[] args, Object target);
}
