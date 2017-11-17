package tychus.echo;

import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class EchoServerHandler extends ChannelInboundHandlerAdapter {
  private static final Logger log = LoggerFactory.getLogger(EchoServerHandler.class);

  @Override
  public void channelRead(ChannelHandlerContext ctx, Object msg) {
    log.debug("channel read");

    log.debug("channel write");
    ctx.write(msg);
    ctx.flush();
  }

  @Override
  public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
    log.warn("channel error, ", cause);
    ctx.close();
  }
}
