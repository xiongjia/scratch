package transmuter;

public class DigitObserver implements Observer {
  @Override
  public void update(NumberGenerator generator) {
    System.out.println("DigitObserver: " + generator.getNumber());
  }
}
