package portal;


import org.springframework.boot.actuate.health.Health;
import org.springframework.boot.actuate.health.HealthIndicator;
import org.springframework.stereotype.Component;

@Component
public class PortalHealthIndicator implements HealthIndicator {
  @Override
  public Health health() {
    return Health.down().withDetail("Memory Usage", "Limit reached.").build();
  }
}
