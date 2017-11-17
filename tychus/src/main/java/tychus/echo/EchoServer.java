package tychus.echo;

import io.netty.channel.ChannelHandler;

import tychus.misc.Server;

public class EchoServer extends Server {
  public EchoServer(int port) {
    super("echo", port);
  }

  @Override
  protected ChannelHandler createChannelHandler() {
    return new EchoServerHandler();
  }
}
