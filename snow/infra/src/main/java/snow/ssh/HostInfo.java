package snow.ssh;

import com.google.common.base.Strings;

public class HostInfo {
  private static int DEFAULT_SSH_PORT = 22;

  private String address;
  private int sshPort = DEFAULT_SSH_PORT;
  private String sshUsername;
  private String sshPassword;

  public HostInfo() {
    // NOP
  }

  /** New HostInfo (address, port, username, password). */
  public HostInfo(String address, int sshPort, String sshUsername, String sshPassword) {
    this.address = address;
    this.sshPort = sshPort;
    this.sshUsername = sshUsername;
    this.sshPassword = sshPassword;
  }

  /** New HostInfo (address, username, password). */
  public HostInfo(String address, String sshUsername, String sshPassword) {
    this.address = address;
    this.sshUsername = sshUsername;
    this.sshPassword = sshPassword;
  }

  /** Check the host info. */
  public boolean validate() {
    return Strings.isNullOrEmpty(address) || Strings.isNullOrEmpty(sshUsername);
  }

  public String getAddress() {
    return address;
  }

  public void setAddress(String address) {
    this.address = address;
  }

  public int getSshPort() {
    return sshPort;
  }

  public void setSshPort(int sshPort) {
    this.sshPort = sshPort;
  }

  public String getSshUsername() {
    return sshUsername;
  }

  public void setSshUsername(String sshUsername) {
    this.sshUsername = sshUsername;
  }

  public String getSshPassword() {
    return sshPassword;
  }

  public void setSshPassword(String sshPassword) {
    this.sshPassword = sshPassword;
  }

  @Override
  public String toString() {
    return "HostInfo{"
      + "address='" + address + '\''
      +  ", sshPort=" + sshPort
      + ", sshUsername='" + sshUsername + '\'' + '}';
  }
}
