package scratch.server;

import io.netty.channel.Channel;

import org.glassfish.jersey.netty.httpserver.NettyHttpContainerProvider;
import org.glassfish.jersey.server.ResourceConfig;

import java.net.InetSocketAddress;
import java.net.URI;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;

public class JerseyScratch {
  final ResourceConfig servResConf = new ResourceConfig(RootResource.class);
  private Channel serv;
  private int port;

  @Path("/")
  public static class RootResource {
    @GET
    @Produces("text/plain")
    public String getIndex() {
      return "index";
    }

    @GET
    @Path("/hello")
    @Produces("text/plain")
    public String getHello() {
      return "hello";
    }

    @GET
    @Path("/hello2")
    @Produces("text/plain")
    public String getHello2() {
      return "hello2";
    }
  }

  JerseyScratch(int port) {
    this.port = port;
  }

  public int getServPort() {
    final int port = ((InetSocketAddress)this.serv.localAddress()).getPort();
    return port;
  }

  public void start() {
    final URI baseUri = URI.create(String.format("http://localhost:%d/", this.port));
    this.serv = NettyHttpContainerProvider.createHttp2Server(baseUri, servResConf, null);
  }

  public void stop() {
    this.serv.close();
  }

  public static class Builder {
    private int port = 0;

    public Builder setPort(int port) {
      this.port = port;
      return this;
    }

    public JerseyScratch build() {
      return new JerseyScratch(this.port);
    }
  }
}
