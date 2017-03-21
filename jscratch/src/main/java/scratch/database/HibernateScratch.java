package scratch.database;

import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.Transaction;
import org.hibernate.cfg.Configuration;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class HibernateScratch {
  private static final Logger log = LoggerFactory.getLogger(HibernateScratch.class);

  /** Hibernate tests. */
  public static void testHibernate() {
    /* default sqlite db folder is ${db_store_sqlite}
     * Add this folder to JVM args (e.g. -Ddb_store_sqlite=f:/db
     */
    final SessionFactory sessionFactory = new Configuration()
        .configure()
        .buildSessionFactory();

    final Session session = sessionFactory.openSession();

    log.debug("begin transaction");
    final Transaction tx = session.beginTransaction();

    /* insert data */
    final UserEntity usr1 = new UserEntity(new Long(1), "user_1");
    session.save(usr1);
    final UserEntity usr2 = new UserEntity(new Long(2), "user_2");
    session.save(usr2);

    tx.commit();
    log.debug("commit transaction");
    session.close();
  }
}
