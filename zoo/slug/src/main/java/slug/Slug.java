package slug;

import java.lang.annotation.Annotation;

public class Slug {

  /** Slug tests. */
  public String getVersion() {
    return "version 1.1";
  }


  @TestTag(name = "name1", age = 11)
  public void test() {

  }

  public void listTag() {
    try {
      final Class aClass = Class.forName("slug.Slug");
      final Annotation[] annotations = aClass.getMethod("test").getAnnotations();
      System.out.println(annotations);
      for (final Annotation annotation : annotations) {
        System.out.println(annotation);
        if (annotation instanceof TestTag) {
          final TestTag tag = (TestTag)annotation;
          System.out.println("tag.name() " + tag.name());
          System.out.println("tag.age() " + tag.age());
        }
      }
    } catch (NoSuchMethodException | ClassNotFoundException e) {
      e.printStackTrace();
    }
  }
}
