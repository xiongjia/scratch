package yahaha;

import com.google.common.base.Strings;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import sheikah.ExecService;
import sheikah.ShellService;


@RestController
public class Controller {
  private static final Logger log = LoggerFactory.getLogger(Controller.class);

  @Autowired
  private ExecService execService;
  @Autowired
  private ShellService shellService;


  @GetMapping("/exec")
  public String exec() {
    final String message = Strings.nullToEmpty(execService.getMessage());
    final String shellSvc = shellService.getMessage();
    log.debug("a testing api: {}, {}", message, shellSvc);

    execService.run();
    return message + ", " + shellSvc;
  }
}
