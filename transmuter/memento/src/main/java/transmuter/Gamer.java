package transmuter;

public class Gamer {
  int money;

  public Gamer(int money) {
    this.money = money;
  }

  public int getMoney() {
    return money;
  }

  public void addMoney(int money) {
    this.money += money;
  }

  public Memento createMemento() {
    final Memento memenot = new Memento(money);
    return memenot;
  }
}
