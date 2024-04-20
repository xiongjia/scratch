package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import transmuter.factory.Factory;
import transmuter.factory.Page;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Abstract factory tests. */
  public static void main(String[] args) {
    log.debug("Abstract factory tests");
    final Factory listFactory = Factory.getFactory("transmuter.listfactory.ListFactory");
    final Page testListPage = listFactory.createTestPage();
    log.debug("ListPage: {}", testListPage.makeHtml());
  }
}
