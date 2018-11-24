package transmuter.pagemaker;

public class PageMaker {
  private PageMaker() {
    // NOP
  }

  public static String makeWelcomePage(String mailAddr) {
    final String user = Database.getUsername(mailAddr);
    return "Welcome " + user + " !";
  }
}
