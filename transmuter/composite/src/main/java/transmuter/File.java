package transmuter;

public class File extends Entry {
  private final String name;
  private final int size;

  public File(String name, int size) {
    this.name = name;
    this.size = size;
  }

  @Override
  public String getName() {
    return this.name;
  }

  @Override
  public int getSize() {
    return this.size;
  }

  @Override
  protected void printList(String prefix) {
    System.out.println(prefix + "/" + this);
  }
}
