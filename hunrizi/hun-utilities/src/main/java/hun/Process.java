package hun;

import java.io.IOException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;
import org.zeroturnaround.exec.ProcessExecutor;
import org.zeroturnaround.exec.stream.LogOutputStream;

public class Process {
  public static void main(String[] args)
      throws InterruptedException, TimeoutException, IOException {
    int exit = new ProcessExecutor()
        .command("java", "-version")
        .timeout(300, TimeUnit.SECONDS)
        .readOutput(true).redirectOutput(new LogOutputStream() {
          @Override
          protected void processLine(String line) {
            System.out.println("line: " + line);
          }
        })
        .execute().getExitValue();
    System.out.println(String.format("Exit %d", exit));
  }
}
