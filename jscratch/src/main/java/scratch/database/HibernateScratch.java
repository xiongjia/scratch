package scratch.database;

import java.util.List;

import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.Transaction;
import org.hibernate.cfg.Configuration;
import org.hibernate.query.Query;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class HibernateScratch {
  private static final Logger log = LoggerFactory.getLogger(HibernateScratch.class);

  private static void insert(Session session) {
    log.debug("begin transaction");
    final Transaction tx = session.beginTransaction();

    /* insert data */
    final UserEntity usr1 = new UserEntity("user_1");
    session.save(usr1);
    log.debug("usr1[{}]: {}", usr1.getId(), usr1.getName());

    final UserEntity usr2 = new UserEntity("user_2");
    session.save(usr2);
    log.debug("usr2[{}]: {}", usr2.getId(), usr2.getName());

    tx.commit();
    log.debug("commit transaction");
  }

  private static void query(Session session) {
    final String hql = "FROM UserEntity WHERE name = :name";
    final Query<UserEntity> query = session.createQuery(hql, UserEntity.class);
    query.setParameter("name", "user_1");

    final List<UserEntity> results = query.list();
    results.stream().forEach(item -> {
      log.debug("user[{}] = {}", item.getId(), item.getName());
    });
  }

  /** Hibernate tests. */
  public static void testHibernate() {
    /* default sqlite db folder is ${db_store_sqlite}
     * Add this folder to JVM args (e.g. -Ddb_store_sqlite=f:/db
     */
    final SessionFactory sessionFactory = new Configuration()
        .configure()
        .buildSessionFactory();

    final Session session = sessionFactory.openSession();

    insert(session);
    query(session);

    session.close();
    sessionFactory.close();
  }
}
