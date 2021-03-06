package scratch.misc;

import org.joor.Reflect;
import org.joor.ReflectException;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class ReflectionScratch {
  private static final Logger log = LoggerFactory.getLogger(ReflectionScratch.class);

  public static class TestObj {
    public String testMethod(String src1, String src2) {
      return String.format("Src = %s; %s", src1, src2);
    }
  }

  /** jOOR tests. */
  public static void testjoor() {
    try {
      final Reflect objReflect = Reflect.on("scratch.misc.ReflectionScratch$TestObj");
      final String ret = objReflect.create().call("testMethod", "str1", "str2").get();
      log.debug("ret: {}", ret);
    } catch (ReflectException reflectErr) {
      log.error("err: ", reflectErr);
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
