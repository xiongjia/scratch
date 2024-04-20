package transmuter;

import java.util.Iterator;

public class ListVisitor extends Visitor {
  private String currentDir = "";

  @Override
  public void visit(File file) {
    System.out.println(currentDir + "/" + file);
  }

  @Override
  public void visit(Directory directory) {
    System.out.println(currentDir + "/" + directory);

    final String saveDir = currentDir;
    currentDir = currentDir + "/" + directory.getName();
    final Iterator it = directory.iterator();
    while (it.hasNext()) {
      Entry entry = (Entry)it.next();
      entry.accept(this);
    }
    currentDir = saveDir;
  }
}
