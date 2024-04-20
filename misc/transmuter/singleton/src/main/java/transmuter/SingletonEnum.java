package transmuter;

public enum SingletonEnum {
  INSTANCE;

  SingletonEnum() {
    System.out.println("SingletonEnum construct()");
  }

  public void print() {
    System.out.println(SingletonEnum.class);
  }
}
