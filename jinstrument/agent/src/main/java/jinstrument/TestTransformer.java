package jinstrument;

import java.io.ByteArrayInputStream;
import java.lang.instrument.ClassFileTransformer;
import java.security.ProtectionDomain;

import javassist.ClassPool;
import javassist.CtClass;
import javassist.CtMethod;

public class TestTransformer implements ClassFileTransformer {
  @Override
  public byte[] transform(ClassLoader loader,
                          String className,
                          Class<?> classBeingRedefined,
                          ProtectionDomain protectionDomain,
                          byte[] classfileBuffer) {
    if (!className.equals("jinstrument/TestApp")) {
      return null;
    }
    System.out.println(String.format("Transform: %s", className));

    try {
      final ClassPool classPool = ClassPool.getDefault();
      final CtClass testAppClass = classPool.get("jinstrument.TestApp");
      final CtMethod testMethod = testAppClass.getDeclaredMethod("test");
      testMethod.addLocalVariable("startTime", CtClass.longType);
      testMethod.insertBefore("startTime = System.nanoTime();");
      testMethod.insertAfter("System.out.println(\"Execution Duration "
          + "(nano sec): \"+ (System.nanoTime() - startTime) );");
      final byte[] newCalssBytecode = testAppClass.toBytecode();
      testAppClass.detach();
      return newCalssBytecode;
    } catch (Throwable ex) {
      ex.printStackTrace();
    }
    return classfileBuffer;
  }
}
