package snow.data;

import org.springframework.data.annotation.Id;
import org.springframework.data.redis.core.RedisHash;

import java.io.Serializable;

@RedisHash(value = "dataitem", timeToLive = 180)
public class DataItem implements Serializable {
  private static final long serialVersionUID = 946670577457860370L;

  @Id
  private String name;
  private String value;

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getValue() {
    return value;
  }

  public void setValue(String value) {
    this.value = value;
  }

  @Override
  public String toString() {
    return "DataItem{"
        + "name='" + name + '\''
        + ", value='" + value + '\''
        + '}';
  }
}
