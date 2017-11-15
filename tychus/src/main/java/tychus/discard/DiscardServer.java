package tychus.discard;

import io.netty.channel.ChannelHandler;

import tychus.misc.Server;

public class DiscardServer extends Server {
  public DiscardServer(int port) {
    super("discard", port);
  }

  @Override
  protected ChannelHandler createChannelHandler() {
    return new DiscardServerHandler();
  }
}
