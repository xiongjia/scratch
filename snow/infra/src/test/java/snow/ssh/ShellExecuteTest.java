package snow.ssh;

import org.junit.Test;

public class ShellExecuteTest {
  @Test
  public void testSshExec() {
    final HostInfo testHost = new HostInfo("192.168.1.37", "lexj", "asdfgh");
    final ShellExecute shellExecute = new ShellExecute(testHost);



  }
}
