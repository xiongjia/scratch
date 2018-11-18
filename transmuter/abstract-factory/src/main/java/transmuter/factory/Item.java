package transmuter.factory;

public abstract class Item {
  protected final String caption;

  public Item(String caption) {
    this.caption = caption;
  }

  public abstract String makeHtml();
}
