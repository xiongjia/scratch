package snow.data;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;

@ConditionalOnProperty(prefix = "snow", name = "data.enableCache")
@Service
public class CacheService {
  private static final Logger log = LoggerFactory.getLogger(CacheService.class);

  @Autowired
  private DataItemRepository dataItemRepository;

  @PostConstruct
  public void init() {
    log.debug("cache service init");
  }

  public void test() {
    log.debug("service test");

    final DataItem item = new DataItem();
    item.setId("1");
    item.setName("first");
    item.setValue("first value");
    dataItemRepository.put(item);

    final DataItem result = dataItemRepository.get("1");
    log.debug("result: {}", result.toString());
  }
}
