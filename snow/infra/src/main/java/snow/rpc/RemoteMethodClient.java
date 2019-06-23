package snow.rpc;

import java.rmi.NotBoundException;
import java.rmi.RemoteException;
import java.rmi.registry.LocateRegistry;
import java.rmi.registry.Registry;

public class RemoteMethodClient {
  static public String client(int serverPort) throws RemoteException, NotBoundException {
    final Registry registry = LocateRegistry.getRegistry(serverPort);
    final RemoteMethod method = (RemoteMethod)registry.lookup("echo");
    return method.echo("test");
  }

  public static void main(String args[]) throws RemoteException, NotBoundException {
    final String result = client(1807);
    System.out.println("result: " + result);
  }
}
