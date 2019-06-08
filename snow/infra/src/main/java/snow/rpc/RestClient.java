package snow.rpc;

import com.google.common.base.Strings;
import org.glassfish.jersey.apache.connector.ApacheClientProperties;
import org.glassfish.jersey.apache.connector.ApacheConnectorProvider;
import org.glassfish.jersey.client.ClientConfig;
import org.glassfish.jersey.client.ClientProperties;

import javax.ws.rs.client.Client;
import javax.ws.rs.client.ClientBuilder;
import javax.ws.rs.client.Invocation;
import javax.ws.rs.client.WebTarget;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.concurrent.TimeUnit;

public class RestClient {
  private static final long TIMEOUT_MS_CONNNECTION = TimeUnit.SECONDS.toMillis(10);
  private static final long TIMEOUT_MS_READ = TimeUnit.SECONDS.toMillis(30);
  private static final String acceptedResponseType = MediaType.APPLICATION_JSON;

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


  public RestClient(RestConnectionManager connectionManager,
                    long timeoutMsConnection, long timeoutMsRead) {
    this.connectionManager = connectionManager;
    this.timeoutMsConnection = timeoutMsConnection;
    this.timeoutMsRead = timeoutMsRead;

    final ApacheConnectorProvider connectorProvider = new ApacheConnectorProvider();
    final ClientConfig clientConfig = new ClientConfig()
      .connectorProvider(connectorProvider)
      .property(ApacheClientProperties.CONNECTION_MANAGER, connectionManager)
      .property(ClientProperties.CONNECT_TIMEOUT, timeoutMsConnection)
      .property(ClientProperties.READ_TIMEOUT, timeoutMsRead);
    this.client = ClientBuilder.newClient();
  }

  public WebTarget makeTarget(String uri, String path) {
    final WebTarget webTarget = client.target(uri);
    if (!Strings.isNullOrEmpty(path)) {
      return webTarget.path(path);
    } else {
      return webTarget;
    }
  }

  public Response invokeGet(WebTarget webTarget) {
    final Invocation.Builder invocationBuilder = webTarget.request(acceptedResponseType);
    return invocationBuilder.get();
  }
}
