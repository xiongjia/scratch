package scratch.webdriver;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import org.openqa.selenium.WebDriver;
import org.openqa.selenium.chrome.ChromeDriver;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.concurrent.TimeUnit;

public class WebDriverUnitSample {
  private static final Logger log = LoggerFactory.getLogger(WebDriverUnitSample.class);
  private WebDriver driver;

  /** Creating webdriver instance. */
  @Before
  public final void setUp() {
    log.debug("creating webdriver instance");
    driver = new ChromeDriver();
    driver.manage().timeouts().implicitlyWait(30, TimeUnit.SECONDS);
  }

  /** Closing webdriver instance. */
  @After
  public final void tearDown() {
    log.info("closing webdriver instance");
    driver.close();
    driver.quit();
  }

  @Test
  public void webDriverTest() {
    final String webPage = "http://localhost:8083/index.html";
    log.info("opening web page: {}", webPage);
    driver.get(webPage);

    final String pageSrc = driver.getPageSource();
    log.debug("page source: {}", pageSrc);
  }
}
