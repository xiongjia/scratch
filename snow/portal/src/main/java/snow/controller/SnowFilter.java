package snow.controller;

import java.io.IOException;
import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.annotation.WebFilter;
import javax.servlet.http.HttpServletRequest;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@WebFilter(urlPatterns = { "/filter/*" })
public class SnowFilter implements Filter {
  private static final Logger log = LoggerFactory.getLogger(SnowFilter.class);

  @Override
  public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
      throws IOException, ServletException {
    final String url = request instanceof HttpServletRequest
        ? ((HttpServletRequest) request).getRequestURL().toString() : "N/A";

    log.debug("filter: {}", url);
    chain.doFilter(request, response);
  }
}
