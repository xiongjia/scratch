package scratch;

import lombok.Setter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class SpringScratch {
  private static final Logger log = LoggerFactory.getLogger(SpringScratch.class);

  public static class PrintMessage {
    @Setter private String message;

    public void printMessage() {
      log.info("Message: {}", message);
    }
  }

  /** Spring tests. */
  @SuppressWarnings("resource")
  public static void testSpring() {
    final ApplicationContext context = new ClassPathXmlApplicationContext("spring-scratch.xml");
    final PrintMessage obj = (PrintMessage)context.getBean("PrintMessage");
    obj.printMessage();
  }
}
