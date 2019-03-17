package wserver;

import java.io.*;
import javax.servlet.*;
import javax.servlet.http.*;

public class WServer extends HttpServlet {
  private static final long serialVersionUID = 1L;

  @Override
	public void doGet(HttpServletRequest request,
      HttpServletResponse response) throws ServletException, IOException {
    response.setContentType("text/html;charset=UTF-8");
    final PrintWriter out = response.getWriter();
    out.println("<!DOCTYPE html>");
    out.println("<html><head>");
    out.println("<meta http-equiv='Content-Type' content='text/html; charset=UTF-8'>");
    out.println("<title>test</title></head>");
    out.println("<body>");
    out.println("<h1>a simple servlet test</h1>");
    out.println("</body>");
    out.println("</html>");
    out.close();
	}
}
