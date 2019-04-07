package portal;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.context.annotation.Configuration;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.annotation.Scheduled;

@Configuration
@ConditionalOnProperty(prefix = "portal.scheduling", value = {"enable"}, havingValue = "true")
@EnableScheduling
public class Scheduling {
  private static final Logger logger = LoggerFactory.getLogger(Scheduling.class);
  private int taskCnt = 0;

  @Scheduled(cron = "${portal.scheduling.taskPing}")
  public void task() {
    logger.debug("Task # {}", ++taskCnt);
  }
}
