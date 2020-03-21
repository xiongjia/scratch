package hun;

import io.reactivex.Observable;
import io.reactivex.disposables.Disposable;
import java.io.IOException;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
import org.zeroturnaround.exec.ProcessExecutor;
import org.zeroturnaround.exec.ProcessResult;
import org.zeroturnaround.exec.stream.LogOutputStream;

public class Process {
  public static void main(String[] args) throws IOException {
    final Future<ProcessResult> future =  new ProcessExecutor()
        .command("java", "-version")
        .timeout(300, TimeUnit.SECONDS)
        .readOutput(true).redirectOutput(new LogOutputStream() {
          @Override
          protected void processLine(String line) {
            System.out.println("line: " + line);
          }
        })
        .start().getFuture();

    final Disposable disposable = Observable.fromFuture(future).subscribe(
        result -> System.out.println(String.format("Exit %d", result.getExitValue())),
        Throwable::printStackTrace, () -> System.out.println("Done"));
    System.out.println("exit");
  }
}
