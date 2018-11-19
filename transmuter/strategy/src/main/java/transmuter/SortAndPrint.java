package transmuter;

public class SortAndPrint {
  private final Comparable[] data;
  private final Sorter sorter;

  public SortAndPrint(Comparable[] data, Sorter sorter) {
    this.data = data;
    this.sorter = sorter;
  }

  /** Print data. */
  public void print() {
    for (int i = 0; i < data.length; i++) {
      System.out.print(i == 0 ?  data[i] : (", " + data[i]));
    }
    System.out.print("\n");
  }

  /** Print & sort data. */
  public void execute() {
    print();
    sorter.sort(data);
    print();
  }
}
