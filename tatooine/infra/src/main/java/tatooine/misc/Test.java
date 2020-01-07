package tatooine.misc;

import org.apache.shiro.SecurityUtils;
import org.apache.shiro.authc.UsernamePasswordToken;
import org.apache.shiro.config.Ini;
import org.apache.shiro.mgt.DefaultSecurityManager;
import org.apache.shiro.mgt.SecurityManager;
import org.apache.shiro.realm.text.IniRealm;
import org.apache.shiro.subject.Subject;

public class Test {
  public static void main(String[] args) {
    final Ini config = Ini.fromResourcePath("classpath:shiro.ini");
    final IniRealm iniRealm = new IniRealm(config);
    final SecurityManager securityManager = new DefaultSecurityManager(iniRealm);
    SecurityUtils.setSecurityManager(securityManager);
    final Subject subject = SecurityUtils.getSubject();
    final UsernamePasswordToken token = new UsernamePasswordToken("root", "secret");

    subject.login(token);
    System.out.println("test");
  }
}
