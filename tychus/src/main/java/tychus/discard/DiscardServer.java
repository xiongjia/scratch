package tychus.discard;

import io.netty.bootstrap.ServerBootstrap;

import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;

public class DiscardServer {
  private int port;

  public DiscardServer(int port) {
    this.port = port;
  }

  /** Start server. */
  public void run() throws Exception {
    final EventLoopGroup srvGroup = new NioEventLoopGroup();
    final EventLoopGroup wrkGroup = new NioEventLoopGroup();

    try {
      final ServerBootstrap bootstrap = new ServerBootstrap();
      bootstrap.group(srvGroup, wrkGroup)
        .channel(NioServerSocketChannel.class)
        .childHandler(new ChannelInitializer<SocketChannel>() {
          @Override
          public void initChannel(SocketChannel channel) throws Exception {
            channel.pipeline().addLast(new DiscardServerHandler());
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
