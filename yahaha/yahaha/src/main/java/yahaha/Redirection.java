package yahaha;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import javax.servlet.http.HttpServletRequest;
import java.net.URI;
import java.net.URISyntaxException;

@RestController
@RequestMapping("/redirection")
public class Redirection {
  private static final Logger log = LoggerFactory.getLogger(Redirection.class);

  @RequestMapping("/**")
  public String mirror(@RequestBody(required = false) String body,
                       HttpMethod method, HttpServletRequest request) throws URISyntaxException {
    log.debug("redirection: {} [{}] - {}", request.getRequestURI(), method.name(), request.getQueryString());
    final URI uri = new URI("http", null, "localhost", 8891,
        request.getRequestURI(), request.getQueryString(), null);

    final RestTemplate restTemplate = new RestTemplate();
    final ResponseEntity<String> responseEntity = restTemplate.exchange(uri,
      method, new HttpEntity<String>(body), String.class);
    return responseEntity.getBody();
  }
}
