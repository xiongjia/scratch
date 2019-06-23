package snow.rpc;

import java.rmi.RemoteException;
import java.rmi.registry.LocateRegistry;
import java.rmi.registry.Registry;
import java.rmi.server.UnicastRemoteObject;

public class RemoteMethodServer extends UnicastRemoteObject implements RemoteMethod {
  private static final long serialVersionUID = 5473723816692450699L;

  protected RemoteMethodServer() throws RemoteException {
    super();
  }

  @Override
  public String echo(String message) throws RemoteException {
    System.out.println("client call message: " + message);
    return "rep: " + message;
  }

  static public void bind(int port) throws RemoteException {
    final Registry registry = LocateRegistry.createRegistry(port);
    registry.rebind("echo", new RemoteMethodServer());
  }

  public static void main(String args[]) throws RemoteException {
    bind(1807);
    System.out.println("server ready...");
  }
}
