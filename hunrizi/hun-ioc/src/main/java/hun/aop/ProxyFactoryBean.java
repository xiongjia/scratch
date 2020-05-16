package hun.aop;

import com.google.common.base.Strings;
import java.lang.reflect.Proxy;
import net.sf.cglib.proxy.Enhancer;
import net.sf.cglib.proxy.MethodInterceptor;

public class ProxyFactoryBean {
  private Object target;
  private MethodInterceptor interceptor;
  private String proxyInterface;

  public Object getTarget() {
    return target;
  }

  public void setTarget(Object target) {
    this.target = target;
  }

  public MethodInterceptor getInterceptor() {
    return interceptor;
  }

  public void setInterceptor(MethodInterceptor interceptor) {
    this.interceptor = interceptor;
  }

  public String getProxyInterface() {
    return proxyInterface;
  }

  public void setProxyInterface(String proxyInterface) {
    this.proxyInterface = proxyInterface;
  }

  public Object createProxy() {
    return Strings.isNullOrEmpty(proxyInterface) ? createCgLibProxy() : createJdkProxy();
  }

  private Object createCgLibProxy() {
    final Enhancer enhancer = new Enhancer();
    enhancer.setSuperclass(target.getClass());
    enhancer.setCallback(interceptor);
    return enhancer.create();
  }

  private Object createJdkProxy() {
    Class<?> clazz;
    try {
      clazz = Class.forName(proxyInterface);
    } catch (ClassNotFoundException e) {
      e.printStackTrace();
      throw new RuntimeException(proxyInterface + "cannot found");
    }
    return Proxy.newProxyInstance(clazz.getClassLoader(), new Class[] { clazz },
      (proxy, method, args) -> interceptor.intercept(target, method, args, null));
  }
}
