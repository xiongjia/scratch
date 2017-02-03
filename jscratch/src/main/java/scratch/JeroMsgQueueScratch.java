package scratch;

import lombok.Builder;
import lombok.Getter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.zeromq.ZContext;
import org.zeromq.ZMQ;
import org.zeromq.ZMQ.Poller;
import org.zeromq.ZMQ.Socket;

import java.util.HashMap;
import java.util.Map;

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
  private ZContext zctx = new ZContext();

  public interface ZChannel {
    void send(String topic, String data);

    Map<String, String> recvStr(long timeoutMilliseconds);

    String getBindEndPoint();

    void close();
  }

  /** Create JeroMQ publisher. */
  public ZChannel buildPublisher() {
    if (zctx == null) {
      return null;
    }
    return ZChannelImpl.builder(zctx, ZMQ.PUB)
      .bindRandomPort(true)
      .bindEndPoint("tcp://127.0.0.1")
      .linger(0)
      .build();
  }

  /** Create JeroMQ subscriber. */
  public ZChannel buildSubscriber(String endPoint, String topic) {
    if (zctx == null) {
      return null;
    }
    return ZChannelImpl.builder(zctx, ZMQ.SUB)
      .connEndPoint(endPoint)
      .topic(topic)
      .build();
  }

  /** Destroy JeroMQ context. */
  public void destoryContext() {
    zctx.close();
    zctx.destroy();
    zctx = null;
  }

  @Builder(toBuilder = true,
           builderClassName = "ZChannelInternalBuilder",
           builderMethodName = "internalBuilder")
  private static class ZChannelImpl implements ZChannel {
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

    private ZContext zctx;
    private Socket zsock;

    @SuppressWarnings("serial")
    @Override
    public Map<String, String> recvStr(long timeoutMilliseconds) {
      log.debug("start recvStr");
      final Poller poller = new Poller(1);
      poller.register(zsock, Poller.POLLIN);
      final int pollerRet = poller.poll(timeoutMilliseconds);
      if (pollerRet == -1) {
        log.debug("invalid poll=(-1)");
        return null;
      }

      log.debug("poll return num = {}", pollerRet);
      if (poller.pollin(0)) {
        final String topic = zsock.recvStr();
        final String data = zsock.recvStr();
        return new HashMap<String, String>() {{
            put("topic", topic);
            put("data", data);
          }
        };
      } else {
        log.debug("invalid !poll(0)");
        return null;
      }
    }

    @Override
    public void close() {
      zsock.close();
    }

    @Override
    public void send(String topic, String data) {
      log.debug("sending: {}, {}", topic, data);
      zsock.sendMore(topic.getBytes());
      zsock.send(data);
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

    public static Builder builder(ZContext zctx, int zsocketType) {
      return new Builder(zctx, zsocketType);
    }

    public static class Builder extends ZChannelInternalBuilder {
      Builder(ZContext zctx, int zsocketType) {
        super();
        this.zsocketType(zsocketType);
        this.zctx(zctx);
      }

      @Override
      public ZChannelImpl build() {
        ZChannelImpl instance = super.build();
        instance.init();
        return instance;
      }
    }
  }
}
