package transmuter;

public abstract class Entry {
  public abstract String getName();

  public abstract int getSize();

  protected abstract void printList(String prefix);

  public void printList() {
    printList("");
  }

  public Entry add(Entry entry) throws FileTreatmentException {
    throw new FileTreatmentException();
  }

  @Override
  public String toString() {
    return String.format("%s[%d]", getName(), getSize());
  }
}
