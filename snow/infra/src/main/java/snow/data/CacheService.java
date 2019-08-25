package snow.data;

import java.util.Optional;
import javax.annotation.PostConstruct;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.stereotype.Service;

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
    item.setName("first");
    item.setValue("first value + 1");
    dataItemRepository.save(item);

    final Optional<DataItem> result = dataItemRepository.findById(item.getName());
    if (result.isPresent()) {
      log.debug("result: {}", result.get().toString());
    }
  }
}