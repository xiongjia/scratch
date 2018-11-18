package transmuter.factory;

public abstract class Link extends Item {
  protected final String url;

  public Link(String caption, String url) {
    super(caption);
    this.url = url;
  }
}
