package transmuter;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.util.Base64;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
  private static final Logger log = LoggerFactory.getLogger(Main.class);

  /** Singleton tests. */
  public static void main(String[] args) {
    log.debug("Singleton tests");

    log.debug("SingletonEager tests");
    final SingletonEager eager = SingletonEager.getInstance();
    eager.print();
    log.debug("SingletonEager {}", eager == SingletonEager.getInstance());

    log.debug("SingletonInnerHolder tests");
    final SingletonInnerHolder innerHolder = SingletonInnerHolder.getInstance();
    innerHolder.print();
    log.debug("SingletonInnerHolder {}", innerHolder == SingletonInnerHolder.getInstance());

    log.debug("SingletonEnum tests");
    final SingletonEnum enumInstance = SingletonEnum.INSTANCE;
    enumInstance.print();

    log.debug("SingletonStaticBlock tests");
    final SingletonStaticBlock staticBlock = SingletonStaticBlock.getInstance();
    staticBlock.print();
    log.debug("SingletonStaticBlock {}", staticBlock == SingletonStaticBlock.getInstance());

    log.debug("SingletonLazy tests");
    final SingletonLazy lazy = SingletonLazy.getInstance();
    lazy.print();
    log.debug("SingletonLazy {}", lazy == SingletonLazy.getInstance());

    log.debug("SingletonSerializable tests");
    final SingletonSerializable serializable = SingletonSerializable.getInstance();
    serializable.print();
    serializable.setTestValue(1000);
    log.debug("value = {}", serializable.getTestValue());
    try {
      final ByteArrayOutputStream byteArray = new ByteArrayOutputStream();
      final ObjectOutputStream objectOutput = new ObjectOutputStream(byteArray);
      objectOutput.writeObject(serializable);
      objectOutput.close();
      final String serializableString = Base64.getEncoder().encodeToString(byteArray.toByteArray());
      log.debug("serializable = {}", serializableString);

      final byte[] data = Base64.getDecoder().decode(serializableString);
      final ObjectInputStream ois = new ObjectInputStream(new ByteArrayInputStream(data));
      final SingletonSerializable obj  = (SingletonSerializable)ois.readObject();
      obj.print();
      log.debug("value = {} (after decode)", serializable.getTestValue());
      log.debug("decode vs encode {}", obj == SingletonSerializable.getInstance());
    } catch (IOException e) {
      e.printStackTrace();
    } catch (ClassNotFoundException e) {
      e.printStackTrace();
    }
  }
}
