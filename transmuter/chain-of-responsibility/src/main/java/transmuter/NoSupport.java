package transmuter;

public class NoSupport extends Support {
  public NoSupport(String name) {
    super(name);
  }

  @Override
  protected boolean resolve(Trouble trouble) {
    return false;
  }
}
