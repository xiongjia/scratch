package transmuter;

public interface Printable {
  void setPrinterName(String name);

  String getPrinterName();

  void print(String string);
}
