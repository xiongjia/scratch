package snow.ssh;

import com.google.common.base.Strings;

import com.jcraft.jsch.JSchException;

import java.io.IOException;

import org.junit.Ignore;
import org.junit.Test;

public class ShellExecuteTest {
  @Ignore
  @Test
  public void testSshExec() {
    final HostInfo testHost = new HostInfo("192.168.1.37", "lexj", "asdfgh");
    final ShellExecute shellExecute = new ShellExecute(testHost);
    shellExecute.setExecuteEvent(new ExecuteEvent() {
      @Override
      public void appendOutput(String message) {
        System.out.println("line: " + Strings.nullToEmpty(message));
      }
    });

    try {
      shellExecute.run(" echo 'abc' && echo '123' && echo '!@#$'");
    } catch (JSchException e) {
      e.printStackTrace();
    } catch (IOException e) {
      e.printStackTrace();
    }
  }
}
