package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Builder tests. */
  public static void main(String[] args) {
    log.debug("Builder tests");

    final Builder textBuilder = new TextBuilder();
    final Director textDirector = new Director(textBuilder);
    textDirector.construct();
    log.debug("Text Builder: {}", textBuilder.getResult());

    final Builder htmlBuilder = new HtmlBuilder();
    final Director htmlDirector = new Director(htmlBuilder);
    htmlDirector.construct();
    log.debug("Html Builder: {}", htmlBuilder.getResult());
  }
}
