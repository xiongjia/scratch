package scratch;

import org.openqa.selenium.WebDriver;
import org.openqa.selenium.chrome.ChromeDriver;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class WebDriverScratch {
  private static final Logger log = LoggerFactory.getLogger(WebDriverScratch.class);

  /** webdriver tests. */
  public static void webDriverTest() {
    final WebDriver driver = new ChromeDriver();
    driver.get("http://localhost:8083/index.html");
    final String pageSrc = driver.getPageSource();
    log.debug("page source: {}", pageSrc);

    driver.close();
    driver.quit();
  }
}
