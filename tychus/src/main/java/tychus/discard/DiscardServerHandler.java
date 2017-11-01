package tychus.discard;

import io.netty.buffer.ByteBuf;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class DiscardServerHandler extends ChannelInboundHandlerAdapter {
  private static final Logger log = LoggerFactory.getLogger(DiscardServerHandler.class);

  @Override
  public void channelRead(ChannelHandlerContext ctx, Object msg) {
    log.debug("channel read");
    ((ByteBuf) msg).release();
  }

  @Override
  public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
    cause.printStackTrace();
    ctx.close();
  }
}
