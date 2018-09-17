package jinstrument;

public class TestApp {
  private void test() {
    System.out.println("Test Method");
  }

  public static void main(String[] args) {
    System.out.println("Test App");
    final TestApp app = new TestApp();
    app.test();
  }
}
