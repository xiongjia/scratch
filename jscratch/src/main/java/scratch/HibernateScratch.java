package scratch;

import org.hibernate.Session;
import org.hibernate.SessionFactory;

import org.hibernate.cfg.Configuration;

public class HibernateScratch {
  /** Hibernate tests. */
  public static void testHibernate() {
    final SessionFactory sessionFactory = new Configuration()
        .configure()
        .buildSessionFactory();

    final Session session = sessionFactory.openSession();

    // import org.hibernate.dialect.SQLiteDialect;
    // org.hibernate.dialect.SQLiteDialect, JdbcEnvironmentInitiator.
    /* TODO sql tests */

    session.close();
  }
}
