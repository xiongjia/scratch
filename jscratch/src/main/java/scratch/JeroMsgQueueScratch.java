package scratch;

import lombok.Builder;
import lombok.Getter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.zeromq.ZContext;
import org.zeromq.ZMQ.Socket;

/**
 * Scratch code for the JeroMQ https://github.com/zeromq/jeromq
 * NOTE:
 *   This sample is using lombok.builder.
 *   In Java IDE (e.g. Eclipse), you need install the Lombok first.
 *   (Run java -jar lombok.jar)
 *
 * @author lexiongjia@gmail.com
 */
public class JeroMsgQueueScratch {
  private static final Logger log = LoggerFactory.getLogger(JeroMsgQueueScratch.class);
  private static final ZContext zctx = new ZContext();

  @Builder(toBuilder = true,
           builderClassName = "ZChannelInternalBuilder",
           builderMethodName = "internalBuilder")
  public static class ZChannel {
    private final long defaultLinger = -1;
    private final long defaultHighWaterMark = 20000;

    private int zsocketType;
    private long sndHighWaterMark = defaultHighWaterMark;
    private long rcvHighWaterMark = defaultHighWaterMark;
    private long linger = defaultLinger;

    private Boolean bindRandomPort = false;
    @Getter private String bindEndPoint = null;
    @Getter private String connEndPoint = null;
    @Getter private String topic = null;

    private Socket zsock;

    public String recvStr() {
      return zsock.recvStr();
    }

    public void send(String topic, String data) {
      zsock.send(topic);
      zsock.sendMore(data);
    }

    private void init() {
      log.debug("Creating zsocket {}", this.zsocketType);
      zsock = zctx.createSocket(this.zsocketType);
      zsock.setLinger(linger);
      zsock.setSndHWM(sndHighWaterMark);
      zsock.setRcvHWM(rcvHighWaterMark);
      subscribe(topic);
      bindEndPoint = bind(bindEndPoint, bindRandomPort);
      connect(connEndPoint);
    }

    private void subscribe(String topic) {
      if (topic == null || topic.isEmpty()) {
        return;
      }
      zsock.subscribe(topic.getBytes());
    }

    private String bind(String endpoint, Boolean randomPort) {
      if (endpoint == null || endpoint.isEmpty()) {
        return null;
      }

      if (!randomPort) {
        log.debug("binding ZSocket: {}", endpoint);
        zsock.bind(endpoint);
        return endpoint;
      }

      log.debug("binding ZSocket to randomport: {}", endpoint);
      final int port = zsock.bindToRandomPort(endpoint);
      final String boundEndPoint = endpoint + ":" + port;
      log.debug("ZSocket bound to prot: {}", boundEndPoint);
      return boundEndPoint;
    }

    private void connect(String endpoint) {
      if (endpoint == null || endpoint.isEmpty()) {
        return;
      }
      log.debug("Connecting: {}", endpoint);
      zsock.connect(endpoint);
    }

    public static Builder builder(int zsocketType) {
      return new Builder(zsocketType);
    }

    public static class Builder extends ZChannelInternalBuilder {
      Builder(int zsocketType) {
        super();
        this.zsocketType(zsocketType);
      }

      @Override
      public ZChannel build() {
        ZChannel instance = super.build();
        instance.init();
        return instance;
      }
    }
  }
}
