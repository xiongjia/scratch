package transmuter;

public class BuyStock implements Order {
  private final Stock stock;

  public BuyStock(Stock stock) {
    this.stock = stock;
  }

  @Override
  public void execute() {
    this.stock.buy();
  }
}
