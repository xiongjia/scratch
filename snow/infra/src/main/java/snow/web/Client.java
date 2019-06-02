package snow.web;

import com.google.common.base.Charsets;
import org.apache.http.HttpEntity;
import org.apache.http.client.config.RequestConfig;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.conn.routing.HttpRoute;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.impl.conn.PoolingHttpClientConnectionManager;
import org.apache.http.pool.PoolStats;
import org.apache.http.util.EntityUtils;

import java.io.IOException;
import java.util.Set;
import java.util.concurrent.TimeUnit;

public class Client {
  private final PoolingHttpClientConnectionManager pool = new PoolingHttpClientConnectionManager();
  private final int TIMEOUT_CONNECTION_MS = (int)TimeUnit.SECONDS.toMillis(10);
  private final int TIMEOUT_CONNECTION_REQ_MS = (int)TimeUnit.SECONDS.toMillis(10);
  private final int TIMEOUT_SOCK_MS = (int)TimeUnit.SECONDS.toMillis(10);


  public String request(String url) {
    final RequestConfig requestConfig = RequestConfig.custom()
        .setConnectTimeout(TIMEOUT_CONNECTION_MS)
        .setConnectionRequestTimeout(TIMEOUT_CONNECTION_REQ_MS)
        .setSocketTimeout(TIMEOUT_SOCK_MS).build();
    final CloseableHttpClient httpClient = HttpClients.custom()
        .setConnectionManager(pool)
        .setDefaultRequestConfig(requestConfig)
        .build();
    final HttpGet httpGet = new HttpGet(url);
    final CloseableHttpResponse response;
    try {
      response = httpClient.execute(httpGet);
      final HttpEntity entity = response.getEntity();
      String responseContent = EntityUtils.toString(entity, Charsets.UTF_8);
      EntityUtils.consume(entity);
      return responseContent;
    } catch (IOException e) {
      e.printStackTrace();
    } finally {
      final Set<HttpRoute> routes = pool.getRoutes();
      for (final HttpRoute route : routes) {
        final PoolStats poolStats = pool.getStats(route);
        System.out.println(poolStats.toString());
      }
    }
    return "";
  }
}
