package sheikah;

import com.jcraft.jsch.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.stereotype.Service;

import java.io.BufferedReader;
import java.io.InputStreamReader;

@Service
@EnableConfigurationProperties(ExecServiceProp.class)
public class ExecService {
  private static final Logger log = LoggerFactory.getLogger(ExecService.class);

  private final ExecServiceProp execServiceProp;

  public ExecService(ExecServiceProp serviceProperties) {
    this.execServiceProp = serviceProperties;
  }

  public String getMessage() {
    return execServiceProp.getMessage();
  }


  public void run() {
    log.debug("run command");

    try {
      final JSch jsch = new JSch();
      final Session session = jsch.getSession("lexj", "192.168.1.37", 22);
      session.setPassword("asdfgh");
      session.setUserInfo(new UserInfo() {
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

        // Accept all server keys
        @Override
        public void showMessage(String s) {
        }
      });

      session.connect();
      final ChannelExec channel = (ChannelExec)session.openChannel("exec");
      channel.setCommand("uname -a");
      channel.setInputStream(null);
      channel.connect(30);
      final BufferedReader input = new BufferedReader(new InputStreamReader(channel.getInputStream()));

      System.out.println("print");
      String line = null;
      while ((line = input.readLine()) != null) {
        // lines.append(line).append("\n");
        System.out.println(line);
      }

      channel.disconnect();
    } catch (Exception e) {

    }
  }
}
