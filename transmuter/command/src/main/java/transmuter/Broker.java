package transmuter;

import java.util.ArrayList;
import java.util.List;

public class Broker {
  private final List<Order> orderList = new ArrayList<>();

  /** Take orders. */
  public void takeOrder(Order order) {
    orderList.add(order);
  }

  /** Place orders. */
  public void placeOrders() {
    for (Order order : orderList) {
      order.execute();
    }
    orderList.clear();
  }
}
