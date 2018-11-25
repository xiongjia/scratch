package transmuter;

public class Circle implements Shape {
  private final String color;
  private int numX;
  private int numY;
  private int radius;

  public Circle(String color) {
    this.color = color;
  }

  public void setNumX(int numX) {
    this.numX = numX;
  }

  public void setNumY(int numY) {
    this.numY = numY;
  }

  public void setRadius(int radius) {
    this.radius = radius;
  }

  @Override
  public String toString() {
    return "Circle{" + "color='"
      + color + '\''
      + ", x=" + numX
      + ", y=" + numY
      + ", radius=" + radius + '}';
  }

  @Override
  public void draw() {
    System.out.println(this.toString());
  }
}
