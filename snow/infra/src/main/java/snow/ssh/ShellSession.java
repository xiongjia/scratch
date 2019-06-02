package snow.ssh;

import com.jcraft.jsch.JSch;
import com.jcraft.jsch.JSchException;
import com.jcraft.jsch.Session;
import com.jcraft.jsch.UserInfo;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;


public class ShellSession {
  private static final Logger logger = LoggerFactory.getLogger(ShellSession.class);
  private static final UserInfo jscUserInfo = new UserInfo() {
    @Override
    public String getPassphrase() {
      return null;
    }

    @Override
    public String getPassword() {
      return null;
    }

    @Override
    public boolean promptPassword(String s) {
      return false;
    }

    @Override
    public boolean promptPassphrase(String s) {
      return false;
    }

    @Override
    public boolean promptYesNo(String s) {
      return true;
    }

    @Override
    public void showMessage(String s) {
    }
  };

  protected final JSch jsch = new JSch();
  private final HostInfo hostInfo;
  private Session session;

  ShellSession(ShellSession session) {
    this.hostInfo = session.getHostInfo();
  }

  ShellSession(HostInfo hostInfo) {
    this.hostInfo = hostInfo;
  }

  private Session makeSession(int timeout) throws JSchException {
    logger.debug("making jsc session: {}", this.hostInfo.toString());
    final Session session = jsch.getSession(this.hostInfo.getSshUsername(),
        this.hostInfo.getAddress(), this.hostInfo.getSshPort());
    session.setConfig("StrictHostKeyChecking", "no");
    session.setConfig("PreferredAuthentications", "password");
    session.setPassword(this.hostInfo.getSshPassword());
    session.setUserInfo(jscUserInfo);
    session.connect(timeout);
    return session;
  }

  public boolean isConnected() {
    return null != session;
  }

  public void connect() throws JSchException {
    this.connect(0);
  }

  /** Open the SSH session. */
  public void connect(int timeout) throws JSchException {
    if (isConnected()) {
      this.disconnect();
    }
    this.session = makeSession(timeout);
  }

  /** Close the SSH session. */
  public void disconnect() {
    if (this.session == null) {
      return;
    }
    this.session.disconnect();
    this.session = null;
  }

  public Session getSession() {
    return session;
  }

  public HostInfo getHostInfo() {
    return hostInfo;
  }
}
