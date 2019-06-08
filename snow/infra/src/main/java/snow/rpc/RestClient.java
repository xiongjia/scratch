package snow.rpc;

import java.util.concurrent.TimeUnit;
import java.util.logging.Level;
import java.util.logging.Logger;

import javax.ws.rs.client.Client;
import javax.ws.rs.client.ClientBuilder;
import javax.ws.rs.client.WebTarget;

import org.glassfish.jersey.apache.connector.ApacheClientProperties;
import org.glassfish.jersey.apache.connector.ApacheConnectorProvider;
import org.glassfish.jersey.client.ClientConfig;
import org.glassfish.jersey.client.ClientProperties;
import org.glassfish.jersey.client.rx.guava.RxListenableFutureInvokerProvider;
import org.glassfish.jersey.logging.LoggingFeature;

public class RestClient {
  private static final long TIMEOUT_MS_CONNNECTION = TimeUnit.SECONDS.toMillis(10);
  private static final long TIMEOUT_MS_READ = TimeUnit.SECONDS.toMillis(30);

  private final RestConnectionManager connectionManager;
  private final long timeoutMsConnection;
  private final long timeoutMsRead;
  private final Client client;

  public RestClient(long timeoutMsConnection, long timeoutMsRead) {
    this(RestConnectionManager.getInstance(), timeoutMsConnection, timeoutMsRead);
  }

  public RestClient() {
    this(new RestConnectionManager(), TIMEOUT_MS_CONNNECTION, TIMEOUT_MS_READ);
  }


  /** RestClient construct. */
  public RestClient(RestConnectionManager connectionManager,
                    long timeoutMsConnection, long timeoutMsRead) {
    this.connectionManager = connectionManager;
    this.timeoutMsConnection = timeoutMsConnection;
    this.timeoutMsRead = timeoutMsRead;

    final ApacheConnectorProvider connectorProvider = new ApacheConnectorProvider();
    final ClientConfig clientConfig = new ClientConfig()
        .connectorProvider(connectorProvider)
        .register(LoggingFeature.class)
        .register(RxListenableFutureInvokerProvider.class)
        .property(ApacheClientProperties.CONNECTION_MANAGER, connectionManager.getPool())
        .property(ApacheClientProperties.CONNECTION_MANAGER_SHARED, true)
        .property(ClientProperties.CONNECT_TIMEOUT, (int)timeoutMsConnection)
        .property(ClientProperties.READ_TIMEOUT, (int)timeoutMsRead);

    final LoggingFeature loggingFeature = new LoggingFeature(
        Logger.getLogger(LoggingFeature.DEFAULT_LOGGER_NAME),
        Level.INFO, LoggingFeature.Verbosity.PAYLOAD_ANY, 1000);
    clientConfig.register(loggingFeature);

    this.client = ClientBuilder.newClient(clientConfig);
  }

  public WebTarget target(String uri) {
    return client.target(uri);
  }
}
