package portal;

import org.springframework.context.annotation.Configuration;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.annotation.Scheduled;

@Configuration
@EnableScheduling
public class Scheduling {
  @Scheduled(cron = "*/35 * * * * *")
  public void task() {
    System.out.println("job task 1");
  }
}
