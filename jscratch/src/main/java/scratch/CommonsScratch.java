package scratch;

import com.google.common.base.Optional;
import com.google.common.base.Strings;
import com.google.common.eventbus.EventBus;
import com.google.common.eventbus.Subscribe;

import org.apache.commons.lang3.StringUtils;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class CommonsScratch {
  private static final Logger log = LoggerFactory.getLogger(CommonsScratch.class);

  /** test strings. */
  public static void testStrings() {
    log.debug("string blank(''): {}", Strings.isNullOrEmpty(""));
    log.debug("string blank(' '): {}", Strings.isNullOrEmpty(" "));
    log.debug("string blank(<null>): {}", Strings.isNullOrEmpty(null));
    log.debug("string blank('test'): {}", Strings.isNullOrEmpty("test"));

    log.debug("string blank('')[apache]: {}", StringUtils.isBlank(""));
    log.debug("string blank(' ')[apache]: {}", StringUtils.isBlank(" "));
    log.debug("string blank(<null>)[apache]: {}", StringUtils.isBlank(null));
    log.debug("string blank('test')[apache]: {}", StringUtils.isBlank("test"));
  }

  /** optional tests. */
  public static void testOptional() {
    final Optional<String> possible = Optional.of("test");
    log.debug("optional isPresent() = {}", possible.isPresent());
    log.debug("optional get() = {}", possible.get());

    final Optional<String> possibleNull = Optional.fromNullable(null);
    log.debug("optional null isPresent() = {}", possibleNull.isPresent());
    log.debug("optional null get() = {}", possibleNull.or("defaultValue"));
  }

  /** event bus. */
  public static void testEventBus() {
    final EventBus bus = new EventBus("TestEventBus");
    bus.register(new Object() {
      @Subscribe
      public void onEvent(String data) {
        log.debug("new event(reg1): {}", data);
      }
    });
    bus.register(new Object() {
      @Subscribe
      public void onEvent(String data) {
        log.debug("new event(reg2): {}", data);
      }
    });
    bus.post("hello");
  }
}
