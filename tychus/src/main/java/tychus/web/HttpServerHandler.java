package tychus.web;

import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelFutureListener;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import io.netty.handler.codec.http.DefaultFullHttpResponse;
import io.netty.handler.codec.http.FullHttpResponse;
import io.netty.handler.codec.http.HttpRequest;
import io.netty.handler.codec.http.HttpUtil;
import io.netty.util.AsciiString;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import static io.netty.handler.codec.http.HttpResponseStatus.OK;
import static io.netty.handler.codec.http.HttpVersion.HTTP_1_1;

public class HttpServerHandler extends ChannelInboundHandlerAdapter {
  private static final Logger log = LoggerFactory.getLogger(HttpServerHandler.class);

  private static final AsciiString CONTENT_TYPE = AsciiString.cached("Content-Type");
  private static final AsciiString CONTENT_LENGTH = AsciiString.cached("Content-Length");
  private static final AsciiString CONNECTION = AsciiString.cached("Connection");
  private static final AsciiString KEEP_ALIVE = AsciiString.cached("keep-alive");
  private static final byte[] CONTENT = { 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd' };

  @Override
  public void channelReadComplete(ChannelHandlerContext ctx) {
    ctx.flush();
  }

  @Override
  public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
    log.warn("channel error, ", cause);
    ctx.close();
  }

  @Override
  public void channelRead(ChannelHandlerContext ctx, Object msg) {
    if (!(msg instanceof  HttpRequest)) {
      return;
    }

    final HttpRequest req = (HttpRequest) msg;
    log.debug("new http request: m={},u={}", req.method(), req.uri());

    final boolean keepAlive = HttpUtil.isKeepAlive(req);
    final FullHttpResponse response = new DefaultFullHttpResponse(HTTP_1_1, OK, Unpooled.wrappedBuffer(CONTENT));
    response.headers().set(CONTENT_TYPE, "text/plain");
    response.headers().setInt(CONTENT_LENGTH, response.content().readableBytes());

    if (!keepAlive) {
      ctx.write(response).addListener(ChannelFutureListener.CLOSE);
    } else {
      response.headers().set(CONNECTION, KEEP_ALIVE);
      ctx.write(response);
    }
  }
}
