package scratch.misc;

import com.google.common.base.Optional;
import com.google.common.base.Strings;
import com.google.common.base.Supplier;
import com.google.common.collect.Lists;
import com.google.common.collect.Maps;
import com.google.common.collect.Table;
import com.google.common.collect.Tables;
import com.google.common.eventbus.EventBus;
import com.google.common.eventbus.Subscribe;
import com.google.common.primitives.UnsignedLong;


import org.apache.commons.lang3.StringUtils;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.Map;

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

  /** collections. */
  public static void testCollections() {
    final List<String> listStr = Lists.newArrayList("test1", "test2", "test3");
    log.debug("strings[{}]: {}", listStr.size(), listStr);
    log.debug("strings[{}]: {}", listStr.size(), Lists.reverse(listStr));

    final Map<UnsignedLong, String> mapStr = Maps.newHashMap();
    mapStr.put(UnsignedLong.valueOf(1), "abc");
    log.debug("map val: {}", mapStr.get(UnsignedLong.valueOf(1)));

    final Table<String, String, Integer> tab = Tables.newCustomTable(
        Maps.<String, Map<String, Integer>>newLinkedHashMap(),
        new Supplier<Map<String, Integer>>() {
          @Override
          public Map<String, Integer> get() {
            return Maps.newLinkedHashMap();
          }
        });
    tab.put("r1", "c1", 100);
    log.debug("tab val: {}", tab.get("r1", "c1"));
    log.debug("tab val: {}", tab.row("r1"));
    log.debug("tab val: {}", tab.column("c1"));
  }
}
