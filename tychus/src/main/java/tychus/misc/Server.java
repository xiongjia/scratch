package tychus.misc;

import io.netty.bootstrap.ServerBootstrap;

import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelHandler;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;


public abstract class Server {
  private static final Logger log = LoggerFactory.getLogger(Server.class);
  private final int port;
  private final String name;

  protected abstract ChannelHandler createChannelHandler();

  public Server(String name, int port) {
    this.name = name;
    this.port = port;
  }

  /** Start server. */
  public void run() throws Exception {
    log.info("{} server (port {}) running", this.name, this.port);

    final EventLoopGroup srvGroup = new NioEventLoopGroup();
    final EventLoopGroup wrkGroup = new NioEventLoopGroup();

    try {
      final ServerBootstrap bootstrap = new ServerBootstrap();
      bootstrap.group(srvGroup, wrkGroup)
        .channel(NioServerSocketChannel.class)
        .childHandler(new ChannelInitializer<SocketChannel>() {
          @Override
          public void initChannel(SocketChannel channel) throws Exception {
            channel.pipeline().addLast(createChannelHandler());
          }
        })
        .option(ChannelOption.SO_BACKLOG, 128)
        .childOption(ChannelOption.SO_KEEPALIVE, true);

      final ChannelFuture future = bootstrap.bind(this.port).sync();
      future.channel().closeFuture().sync();
    } finally {
      srvGroup.shutdownGracefully();
      wrkGroup.shutdownGracefully();
    }
  }
}
