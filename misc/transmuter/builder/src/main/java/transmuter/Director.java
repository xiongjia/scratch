package transmuter;

public class Director {
  private final Builder builder;

  public Director(Builder builder) {
    this.builder = builder;
  }

  /** Construct. */
  public void construct() {
    builder.makeTitle("make title");
    builder.makeString("make string");
    builder.makeItems(new String[] { "item 1", "item 2", "item 3" });
    builder.close();
  }
}
