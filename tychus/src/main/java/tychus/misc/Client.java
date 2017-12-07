package tychus.misc;

import io.netty.handler.codec.http.QueryStringDecoder;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.URI;
import java.net.URISyntaxException;

public class Client {
  private static final Logger log = LoggerFactory.getLogger(Client.class);

  static public void parseUrl() throws URISyntaxException {
    final String url = "HTTP://localhost/api/test?data=1&a=2%323";
    final URI uri = new URI(url);

    log.debug("host: {}, port: {}, path = {}, scheme = {}",
      uri.getHost(), uri.getPort(), uri.getPath(), uri.getScheme());
    log.debug("query: {}, rquery: {}", uri.getQuery(), uri.getRawQuery());

    final String qeuryStr = String.format("%s?%s", uri.getPath(), uri.getRawQuery());
    final QueryStringDecoder decoder = new QueryStringDecoder(qeuryStr);
    log.debug("parameters: {}", decoder.parameters());
  }
}
