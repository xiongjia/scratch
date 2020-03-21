package hun;

import io.reactivex.Observable;
import io.reactivex.disposables.Disposable;
import java.util.concurrent.TimeUnit;

public class Reactive {
  public static void main(String[] args) throws InterruptedException {
    final Observable<String> observable = Observable.just("data1", "data10");
    System.out.println(String.format("Begin: [%s]", Thread.currentThread().getName()));
    Disposable disposable = observable.map(i -> "hdr" + i).subscribe(item -> {
      System.out.println(String.format("Data[%s]: %s", Thread.currentThread().getName(), item));
    });
    System.out.println(String.format("End: [%s]", Thread.currentThread().getName()));

    System.out.println(String.format("Begin Interval: [%s]", Thread.currentThread().getName()));
    final Observable<Long> observableInterval = Observable.interval(2, TimeUnit.SECONDS);
    disposable = observableInterval.subscribe(iterval -> {
      System.out.println(String.format("Interval [%s]: %s",
          Thread.currentThread().getName(), iterval));
    });
    System.out.println(String.format("End Interval: [%s]", Thread.currentThread().getName()));
    Thread.sleep(1000 * 10);
  }
}
