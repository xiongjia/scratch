package scratch.misc;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import scratch.GroovyExtends;
import scratch.GroovyScratch;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class ReflectionScratch {
  private static final Logger log = LoggerFactory.getLogger(ReflectionScratch.class);

  public static class TestObj {
    public String testMethod(String src1, String src2) {
      return String.format("Src = %s; %s", src1, src2);
    }
  }

  /** reflection tests (groovy). */
  public static void testGroovy() {
    final GroovyScratch obj = new GroovyScratch();
    final GroovyExtends errObj = new GroovyExtends();

    try {
      final Method method = obj.getClass().getMethod("get", String.class);
      final String ret = (String)method.invoke(obj, "test string1");
      log.debug("retVal = {}", ret);
      final Method mtGetErrCode = errObj.getClass().getMethod("getErrorCode");
      log.debug("retVal = {}", (int)mtGetErrCode.invoke(errObj));
    } catch (NoSuchMethodException
           | IllegalAccessException
           | IllegalArgumentException
           | InvocationTargetException  err) {
      log.debug("Error: ", err);
    }
  }

  /** reflection tests. */
  public static void test() {
    final Object obj = new TestObj();

    try {
      final Method method = obj.getClass().getMethod("testMethod", String.class, String.class);
      final String ret = (String)method.invoke(obj, "test string1", "test string2");
      log.debug("retVal = {}", ret);
    } catch (NoSuchMethodException
           | IllegalAccessException
           | IllegalArgumentException
           | InvocationTargetException  err) {
      log.debug("Error: ", err);
    }
  }
}
