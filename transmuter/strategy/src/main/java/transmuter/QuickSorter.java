package transmuter;

public class QuickSorter implements Sorter {
  @Override
  public void sort(Comparable[] data) {
    quickSort(0, data.length - 1, data);
  }

  private void quickSort(int pre, int post, Comparable[] data) {
    final int savedPre = pre;
    final int savedPost = post;

    final Comparable mid = data[(pre + post) / 2];
    do {
      while (data[pre].compareTo(mid) < 0) {
        pre++;
      }
      while (mid.compareTo(data[post]) < 0) {
        post--;
      }
      if (pre <= post) {
        Comparable tmp = data[pre];
        data[pre] = data[post];
        data[post] = tmp;
        pre++;
        post--;
      }
    } while (pre <= post);
    if (savedPre < post) {
      quickSort(savedPre, post, data);
    }
    if (pre < savedPost) {
      quickSort(pre, savedPost, data);
    }
  }
}
