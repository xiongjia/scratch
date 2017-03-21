package scratch.distributed;

import akka.actor.ActorRef;
import akka.actor.ActorSystem;
import akka.actor.Props;
import akka.actor.UntypedActor;

import lombok.Builder;
import lombok.Data;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import scala.concurrent.duration.Duration;

import java.util.concurrent.TimeUnit;

public class AkkaScratch {
  private static final Logger log = LoggerFactory.getLogger(AkkaScratch.class);

  @Builder
  @Data
  public static class Message {
    private String message = "defaultMessage";
  }

  public static class MessagePrinter extends UntypedActor {
    @Override
    public void onReceive(Object msg) {
      if (msg instanceof Message) {
        log.debug("msg: {}", ((Message) msg).getMessage());
      }
    }
  }

  public static class MessageHandler extends UntypedActor {
    String msg = "";

    @Override
    public void onReceive(Object msg) {
      if (msg instanceof Message) {
        this.msg = ((Message) msg).getMessage();
        log.debug("msg handler: {}", this.msg);
      } else {
        unhandled(msg);
      }
    }
  }

  /** Akka tests. */
  public static void akkaTest() {
    final ActorSystem system = ActorSystem.create("scratch");

    final ActorRef msgPrinter = system.actorOf(Props.create(MessagePrinter.class));
    final ActorRef msgHandler = system.actorOf(Props.create(MessageHandler.class), "greeter");

    /* testing schedule once */
    system.scheduler().scheduleOnce(Duration.Zero(),
        msgHandler,
        new Message.MessageBuilder().message("once").build(),
        system.dispatcher(), msgPrinter);

    /* testing schedule repeat */
    system.scheduler().schedule(Duration.Zero(), Duration.create(1, TimeUnit.SECONDS),
        msgHandler,
        new Message.MessageBuilder().message("repeat").build(),
        system.dispatcher(), msgPrinter);
  }
}
