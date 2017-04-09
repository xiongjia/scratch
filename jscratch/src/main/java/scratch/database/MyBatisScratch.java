package scratch.database;

import org.apache.ibatis.io.Resources;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.io.Reader;

public class MyBatisScratch {
  private static final Logger log = LoggerFactory.getLogger(MyBatisScratch.class);

  /** MyBatis tests. */
  public static void test() {
    final String conf = "mybatis-conf.xml";
    try {
      final Reader reader = Resources.getResourceAsReader(conf);
      final SqlSessionFactory sessionFactory = new SqlSessionFactoryBuilder().build(reader);
      final SqlSession session = sessionFactory.openSession();

      final MyBatisMapper mapper = session.getMapper(MyBatisMapper.class);
      final UserEntity user = mapper.select();
      log.debug("usrer {}: {}", user.getId(), user.getName());
      session.close();
    } catch (IOException err) {
      log.error("MyBatis error: ", err);
    }
  }
}
