package scratch;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class ServletScratch extends HttpServlet {
  private static final Logger log = LoggerFactory.getLogger(ServletScratch.class);
  private static final long serialVersionUID = 1L;

  @Override
  protected void doGet(HttpServletRequest request, HttpServletResponse response)
      throws IOException, ServletException {
    log.debug("new request");
    request.getRequestDispatcher("scratch.jsp").forward(request, response);
  }
}
