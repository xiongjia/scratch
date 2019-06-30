package snow.daemon;

import java.util.Date;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import javax.annotation.PostConstruct;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.stereotype.Service;

@Service
@ConditionalOnProperty("snow.daemon.enable")
public class DaemonService {
  private static final Logger logger = LoggerFactory.getLogger(DaemonService.class);
  private ScheduledExecutorService scheduledExecutor;


  public class Task implements Runnable {
    private Date lastRunTime = new Date();

    @Override
    public void run() {
      final Date now = new Date();
      logger.debug("Task running ... {} -> {}", lastRunTime, now);
      lastRunTime = now;
    }
  }

  @PostConstruct
  void init() {
    scheduledExecutor = Executors.newScheduledThreadPool(10);
    scheduledExecutor.scheduleAtFixedRate(new Task(), 5, 50, TimeUnit.SECONDS);
  }

  public void stop() {
    scheduledExecutor.shutdownNow();
  }
}
