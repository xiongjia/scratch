package snow.rpc;

import java.rmi.Remote;
import java.rmi.RemoteException;

public interface RemoteMethod extends Remote {
  String echo(String message) throws RemoteException;
}
