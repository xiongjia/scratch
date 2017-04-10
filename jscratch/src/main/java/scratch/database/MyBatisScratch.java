package scratch.database;

import org.apache.ibatis.io.Resources;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.io.Reader;
import java.util.List;

public class MyBatisScratch {
  private static final Logger log = LoggerFactory.getLogger(MyBatisScratch.class);

  private static void selectAll(MyBatisMapper mapper) {
    try {
      final List<UserEntity> users = mapper.selectAll();
      users.stream().forEach((usr) -> {
        log.debug("usr {}: {}", usr.getId(), usr.getName());
      });
    } catch (Throwable err) {
      log.error("SELECT Error", err);
    }
  }

  /** MyBatis tests. */
  public static void test() {
    final String conf = "mybatis-conf.xml";
    try {
      final Reader reader = Resources.getResourceAsReader(conf);
      final SqlSessionFactory sessionFactory = new SqlSessionFactoryBuilder().build(reader);
      final SqlSession session = sessionFactory.openSession();

      final MyBatisMapper mapper = session.getMapper(MyBatisMapper.class);
      selectAll(mapper);
      session.close();
    } catch (IOException err) {
      log.error("MyBatis error: ", err);
    }
  }
}
