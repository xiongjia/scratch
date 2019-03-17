package slug;

import org.apache.commons.exec.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.LinkedList;
import java.util.List;

public class CommandExecution {
  private static final Logger log = LoggerFactory.getLogger(CommandExecution.class);

  public static class ExecResult extends LogOutputStream {
    private int exitCode;
    public int getExitCode() {
      return exitCode;
    }

    public void setExitCode(int exitCode) {
      this.exitCode = exitCode;
    }

    private final StringBuilder sb = new StringBuilder();

    @Override
    protected void processLine(String line, int level) {
      sb.append(line);
    }

    public String getOutput() {
      return sb.toString();
    }
  }

  public String runTest() {
    final CommandLine cmdl = new CommandLine("node");
    cmdl.addArgument("-pe");
    cmdl.addArgument("process.stderr.write('stdout')");

    final DefaultExecutor executor = new DefaultExecutor();
    final ExecuteWatchdog watchdog = new ExecuteWatchdog(60000);
    final ExecResult result = new ExecResult();
    executor.setWatchdog(watchdog);
    // executor.setStreamHandler(new PumpStreamHandler(result));

    final ByteArrayOutputStream stdout = new ByteArrayOutputStream();
    final ByteArrayOutputStream stderr = new ByteArrayOutputStream();
    final PumpStreamHandler psh = new PumpStreamHandler(stdout, stderr);
    executor.setStreamHandler(psh);
    try {
      executor.execute(cmdl);
      final String test = psh.toString();
      final String out = stdout.toString();
      final String errStr = stderr.toString();
      return stdout.toString();
    } catch (IOException e) {
      e.printStackTrace();
      return "";
    }

  }
}
