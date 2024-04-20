package transmuter;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

public abstract class NumberGenerator {
  private final List<Observer> observers = new ArrayList<>();

  public void addObserver(Observer observer) {
    observers.add(observer);
  }

  /** Notify observers. */
  public void notifyObservers() {
    final Iterator itr = observers.iterator();
    while (itr.hasNext()) {
      final Observer observer = (Observer)itr.next();
      observer.update(this);
    }
  }

  public abstract int getNumber();

  public abstract void execute();
}
