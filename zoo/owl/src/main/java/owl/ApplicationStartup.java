package owl;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.ApplicationListener;
import org.springframework.stereotype.Component;

@Component
public class ApplicationStartup implements ApplicationListener<ApplicationReadyEvent> {
  private static final Logger log = LoggerFactory.getLogger(ApplicationStartup.class);

  @Override
  public void onApplicationEvent(ApplicationReadyEvent event) {
    log.debug("app started: {}", event.getTimestamp());
  }
}
