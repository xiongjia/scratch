package transmuter;

public interface Element {
  void accept(Visitor visitor);
}
