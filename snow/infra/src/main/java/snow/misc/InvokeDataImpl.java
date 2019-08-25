package snow.misc;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.util.HashMap;
import java.util.Map;

public class InvokeDataImpl implements InvocationHandler {
  private final Data target = new Data();
  private final Map<String, Method> methods;

  public InvokeDataImpl() {
    methods = new HashMap<>();
    for (Method method: target.getClass().getDeclaredMethods()) {
      this.methods.put(method.getName(), method);
    }
  }

  private final static class Data implements InvokeData {
    @Override
    public String echo(String src) {
      return "SRC: " + src;
    }
  }

  @Override
  public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
    return methods.get(method.getName()).invoke(target, args);
  }
}
