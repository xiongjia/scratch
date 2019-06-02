package snow.ssh;

import com.google.common.base.Strings;
import com.jcraft.jsch.ChannelExec;
import com.jcraft.jsch.JSchException;
import com.jcraft.jsch.Session;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class ShellExecute {
  private static final Logger logger = LoggerFactory.getLogger(ShellSession.class);
  private static final int DEFAULT_SHELL_EXEC_CONNECT_TIMEOUT = 2000;
  private static final int DEFAULT_STREAM_PAUSE_MS = 1000;
  private ExecuteEvent executeEvent;
  private final ShellSession session;

  public ShellExecute(ShellSession session) {
    this.session = new ShellSession(session);
  }

  public ShellExecute(HostInfo hostInfo) {
    this.session = new ShellSession(hostInfo);
  }

  public void setExecuteEvent(ExecuteEvent executeEvent) {
    this.executeEvent = executeEvent;
  }

  private void onShellOutput(String message) {
    if (!Strings.isNullOrEmpty(message) && this.executeEvent != null) {
      this.executeEvent.appendOutput(message);
    }
  }



  private void readOutput(ChannelExec channelExec, InputStream inputStream) throws IOException {
    final BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream));
    while (true) {
      while (reader.ready()) {
        final String line = reader.readLine();
        onShellOutput(line + "\n");
      }

      if (channelExec.isClosed()) {
        if (inputStream.available() > 0) {
          continue;
        }
        break;
      }
      try {
        Thread.sleep(DEFAULT_STREAM_PAUSE_MS);
      } catch (InterruptedException err) {
        logger.warn("SSH Stream pause err", err);
      }
    }
  }
  
  public void run(String command) throws JSchException, IOException {
    this.run(command, DEFAULT_SHELL_EXEC_CONNECT_TIMEOUT);
  }

  /** Run ssh command at remote host via ssh protocol. */
  public void run(String command, int timeout) throws JSchException, IOException {
    if (Strings.isNullOrEmpty(command)) {
      throw new JSchException("Invalid shell command");
    }
    logger.debug("shell exec: {}", command);
    ChannelExec channelExec = null;
    try {
      this.session.connect(timeout);
      final Session jschSession = this.session.getSession();
      channelExec = (ChannelExec)jschSession.openChannel("exec");
      channelExec.setCommand(String.format("%s \n", command));
      channelExec.setInputStream(null);

      final InputStream inputStream = channelExec.getInputStream();
      channelExec.connect(timeout);
      this.readOutput(channelExec, inputStream);
    } catch (JSchException | IOException err) {
      logger.error("cannot exec command {}", command, err);
      throw err;
    } finally {
      if (channelExec != null) {
        channelExec.disconnect();
      }
      this.session.disconnect();
    }
  }
}
