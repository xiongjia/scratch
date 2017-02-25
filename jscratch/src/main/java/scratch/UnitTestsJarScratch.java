package scratch;

import org.junit.runner.Result;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import scratch.UnitTestsScratch.TestItem;

import java.io.File;
import java.net.MalformedURLException;
import java.net.URL;
import java.net.URLClassLoader;
import java.util.ArrayList;
import java.util.List;

class UnitTestsJarScratch {
  private static final Logger log = LoggerFactory.getLogger(UnitTestsJarScratch.class);


  @SuppressWarnings("resource")
  public static ArrayList<TestItem> loadTestFromJar(String jarFile, String className)
      throws ClassNotFoundException, MalformedURLException {
    log.debug("Load UnitTest from {}", jarFile);

    final File file = new File(jarFile);
    final List<URL> urls = new ArrayList<URL>();
    urls.add(file.toURI().toURL());

    final ClassLoader srcClassLoader = new URLClassLoader(urls.toArray(new URL[0]));
    final Class<?> srcClass = srcClassLoader.loadClass(className);
    final UnitTestsScratch unitTestsScratch = new UnitTestsScratch();
    return unitTestsScratch.getTestItems(srcClass);
  }

  @SuppressWarnings("resource")
  public static Result invokeTestFromJar(String jarFile, String className)
      throws ClassNotFoundException, MalformedURLException {
    log.debug("Load UnitTest from {}", jarFile);

    final File file = new File(jarFile);
    final List<URL> urls = new ArrayList<URL>();
    urls.add(file.toURI().toURL());

    final ClassLoader srcClassLoader = new URLClassLoader(urls.toArray(new URL[0]));
    final Class<?> srcClass = srcClassLoader.loadClass(className);
    final UnitTestsScratch unitTestsScratch = new UnitTestsScratch();
    return unitTestsScratch.runTest(srcClass, null);
  }
}
