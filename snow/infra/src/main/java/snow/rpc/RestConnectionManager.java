package snow.rpc;

import org.apache.http.impl.conn.PoolingHttpClientConnectionManager;

public class RestConnectionManager {
  private static final int DEFAULT_POOL_MAX_PRE_ROUTE = 15;
  private static final int DEFAULT_POOL_MAX_TOTAL = 50;

  private final PoolingHttpClientConnectionManager pool = new PoolingHttpClientConnectionManager();

  private static class Holder {
    public static final RestConnectionManager INSTANCE = new RestConnectionManager();
  }

  public RestConnectionManager() {
    this(DEFAULT_POOL_MAX_PRE_ROUTE, DEFAULT_POOL_MAX_TOTAL);
  }

  public RestConnectionManager(int maxPreRoute, int maxTotal) {
    pool.setDefaultMaxPerRoute(maxPreRoute);
    pool.setMaxTotal(maxTotal);
  }

  public static RestConnectionManager getInstance() {
    return Holder.INSTANCE;
  }

  public PoolingHttpClientConnectionManager getPool() {
    return pool;
  }
}
