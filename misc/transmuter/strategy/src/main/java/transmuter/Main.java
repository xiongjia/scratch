package transmuter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Strategy tests. */
  public static void main(String[] args) {
    log.debug("Strategy tests");

    final String[] sourceData1 = { "Dumpty", "Bowman", "Carroll", "Elfland", "Alice" };
    final SortAndPrint quickSorter = new SortAndPrint(sourceData1, new QuickSorter());
    quickSorter.execute();

    final String[] sourceData2 = { "Dumpty", "Bowman", "Carroll", "Elfland", "Alice" };
    final SortAndPrint selectionSorter = new SortAndPrint(sourceData2, new SelectionSorter());
    selectionSorter.execute();
  }
}
