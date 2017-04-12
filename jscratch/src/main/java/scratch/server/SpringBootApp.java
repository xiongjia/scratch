package scratch.server;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class SpringBootApp {
  /** Spring boot application. */
  public static void main(String[] args) {
    SpringApplication.run(SpringBootApp.class, args);
    // final ApplicationContext context = SpringApplication.run(SpringBootApp.class, args);
    //    // System.out.println("Let's inspect the beans provided by Spring Boot:");
    //    final String[] beanNames = context.getBeanDefinitionNames();
    //    Arrays.sort(beanNames);
    //    for (String beanName : beanNames) {
    //      // System.out.println(beanName);
    //    }
  }
}
