package scratch.unittest;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.testng.Assert;
import org.testng.annotations.AfterClass;
import org.testng.annotations.AfterGroups;
import org.testng.annotations.AfterMethod;
import org.testng.annotations.BeforeClass;
import org.testng.annotations.BeforeGroups;
import org.testng.annotations.BeforeMethod;
import org.testng.annotations.Test;

public class TestNgSample {
  private static final Logger log = LoggerFactory.getLogger(TestNgSample.class);

  @BeforeGroups("sampleGroup")
  public void beforeSampleGroups() {
    log.debug("TestNG beforeSampleGroups");
  }

  @AfterGroups("sampleGroup")
  public void afterSampleGroups() {
    log.debug("TestNG afterSampleGroups");
  }

  @BeforeClass
  public void beforeClass() {
    log.debug("TestNG beforeClass");
  }

  @AfterClass
  public void afterClass() {
    log.debug("TestNG afterClass");
  }

  @BeforeMethod
  public void beforeMethod() {
    log.debug("TestNG before method");
  }

  @AfterMethod
  public void afterMethod() {
    log.debug("TestNG after method");;
  }

  @Test(invocationCount = 100,
        threadPoolSize = 5,
        groups = { "sampleGroup", "baseTest", "group3" })
  public void testSample3() {
    log.debug("TestNG testSample3() method");
  }

  @Test(groups = { "sampleGroup", "baseTest", "group2" })
  public void testSample2() {
    log.debug("TestNG testSample2() method");
  }

  @Test(groups = { "sampleGroup", "baseTest", "group1" },
        expectedExceptions = ArithmeticException.class,
        enabled = true,
        timeOut = 5000,
        dependsOnMethods = "testSample2")
  public void testSample1() {
    log.debug("TestNG testSample1() method");

    Assert.assertNotNull("test");
    Assert.assertEquals("123", "123");
    throw new ArithmeticException("test");
  }
}
