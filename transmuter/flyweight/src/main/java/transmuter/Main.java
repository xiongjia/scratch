package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);
  private static final String[] colors = { "Red", "Green", "Blue", "White", "Black" };

  /** Flyweight tests. */
  public static void main(String[] args) {
    log.debug("Flyweight tests");

    for (int i = 0; i < 20; i++) {
      final String color = colors[(int)(Math.random() * colors.length)];
      Circle circle = ShapeFactory.getInstance().getCircle(color);
      circle.setNumX(0 + i);
      circle.setNumY(0 + i);
      circle.setRadius(100);
      circle.draw();
    }
  }
}
