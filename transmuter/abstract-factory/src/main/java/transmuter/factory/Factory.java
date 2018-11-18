package transmuter.factory;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public abstract class Factory {
  private static final Logger log = LoggerFactory.getLogger(Factory.class);

  /** Get factory. */
  public static Factory getFactory(String className) {
    try {
      return (Factory)Class.forName(className).newInstance();
    } catch (ClassNotFoundException e) {
      log.error("Cannot find factory class {}", className, e);
    } catch (Exception e) {
      log.error("Cannot creating factory class {}", className, e);
    }
    return null;
  }

  public abstract Link createLink(String caption, String url);

  public abstract Page createPage(String title, String author);

  /** Create test page. */
  public Page createTestPage() {
    final Page page = createPage("test page", "tester");
    page.add(createLink("test", "http://test.com"));
    return page;
  }
}
