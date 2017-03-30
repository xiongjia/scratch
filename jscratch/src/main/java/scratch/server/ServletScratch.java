package scratch.server;

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
    request.setAttribute("scratch", "test attr");
    request.getRequestDispatcher("scratch.jsp").forward(request, response);
  }

  @Override
  protected void doPost(HttpServletRequest request, HttpServletResponse response)
      throws ServletException, IOException {
    /* The example of the post request:
     * curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d 'data=your data' \
     *    "http://localhost:9000/jscratch/scratch"
     */
    final String reqData = request.getParameter("data");
    log.debug("post parameter: {}", reqData);
    request.setAttribute("scratch", String.format("request parameter: %s",  reqData));
    request.getRequestDispatcher("scratch.jsp").forward(request, response);
  }
}
