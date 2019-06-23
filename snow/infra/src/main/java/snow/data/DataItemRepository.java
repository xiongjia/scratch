package snow.data;

import javax.annotation.PostConstruct;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.HashOperations;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Repository;

@Repository
public class DataItemRepository {
  private static final Logger log = LoggerFactory.getLogger(DataItemRepository.class);
  private static final String KEY = "data-item";

  private HashOperations<String, String, DataItem> hashOperations;

  @Autowired
  private RedisTemplate<String, Object> redisTemplate;

  @PostConstruct
  public void init() {
    log.debug("DataItem Repository init");
    this.hashOperations = this.redisTemplate.opsForHash();
  }


  public void put(DataItem item) {
    hashOperations.put(KEY, item.getId(), item);
  }

  public DataItem get(String id) {
    return hashOperations.get(KEY, id);
  }
}
