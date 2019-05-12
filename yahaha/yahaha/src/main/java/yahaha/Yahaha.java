package yahaha;

import com.jcraft.jsch.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import sheikah.Sheikah;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

// @SpringBootApplication
// @Sheikah
public class Yahaha {
  private static final Logger log = LoggerFactory.getLogger(Yahaha.class);

  /** Yahaha tests. */
  public static void main(String[] args) {
    // log.debug("Yahaha tests");
    // SpringApplication.run(Yahaha.class, args);
    Yahaha.run();
  }


  public static void run() {
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
      final ChannelExec channel = (ChannelExec) session.openChannel("exec");
      channel.setCommand(" uname -a && sleep 50 && uname -a");
      channel.setInputStream(null);
      // final BufferedReader input = new BufferedReader(new InputStreamReader(channel.getInputStream()));
      final InputStream in = channel.getInputStream();
      channel.connect();

      log.error("print");
      byte[] tmp = new byte[1024];
      while (true) {
        while (in.available() > 0) {
          int i = in.read(tmp, 0, 1024);
          if (i < 0) {
            break;
          }
          log.error(new String(tmp, 0, i));
        }
        if (channel.isClosed()) {
          if (in.available() > 0) {
            continue;
          }
          break;
        }
      }

//      String line = null;
//      while ((line = input.readLine()) != null) {
//        // lines.append(line).append("\n");
//        log.error(line);
//      }
      channel.disconnect();
    } catch (Exception e) {
      log.error("err: ", e);
    }

    log.error("finished");
  }
}

