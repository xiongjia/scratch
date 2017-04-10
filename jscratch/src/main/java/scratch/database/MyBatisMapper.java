package scratch.database;

import java.util.List;

public interface MyBatisMapper {
  public List<UserEntity> selectAll();
}
