package hun;

import io.reactivex.Observable;
import io.reactivex.disposables.Disposable;

import java.io.IOException;

public class Reactive {
  public static void main(String[] args) throws IOException {
    final Observable<String> observable = Observable.just("data1", "data10");
    System.out.println(String.format("Begin: [%s]", Thread.currentThread().getName()));
    final Disposable disposable = observable.map(i -> "hdr" + i).subscribe(item -> {
      System.out.println(String.format("Data[%s]: %s", Thread.currentThread().getName(), item));
    });
    System.out.println(String.format("End: [%s]", Thread.currentThread().getName()));
  }
}
