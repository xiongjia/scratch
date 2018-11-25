package transmuter;

import java.util.HashMap;
import java.util.Map;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ShapeFactory {
  private static final Logger log = LoggerFactory.getLogger(ShapeFactory.class);
  private final Map<String, Circle> circleMap = new HashMap<>();

  public static class ShapeFactoryHolder {
    public static final ShapeFactory INSTANCE = new ShapeFactory();
  }

  public static ShapeFactory getInstance() {
    return ShapeFactoryHolder.INSTANCE;
  }

  /** Get circle. */
  public Circle getCircle(String color) {
    final Circle circle = circleMap.get(color);
    if (circle != null) {
      return circle;
    }

    log.debug("creating new circle {}", color);
    final Circle newCircle = new Circle(color);
    circleMap.put(color, newCircle);
    return newCircle;
  }
}
