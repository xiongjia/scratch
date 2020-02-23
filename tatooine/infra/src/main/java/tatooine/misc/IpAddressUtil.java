package tatooine.misc;

import inet.ipaddr.IPAddress;
import inet.ipaddr.IPAddressString;

import java.math.BigInteger;
import java.util.Iterator;

public class IpAddressUtil {

  static void convertNonSequentialBlock(String string) {
    final IPAddressString addrString = new IPAddressString(string);
    final IPAddress addr = addrString.getAddress();
    System.out.println("Initial range block is " + addr);

    System.out.println("Address is " + addr.toInetAddress().getHostAddress());

    final BigInteger sequentialCount = addr.getCount();
    System.out.println("Sequential is " + sequentialCount);

    final Iterable<? extends IPAddress> itr = addr.getIterable();
    itr.forEach(System.out::println);

    final IPAddress lowerAddr = addr.getLower();
    System.out.println("Lower Address is " + lowerAddr.toInetAddress().getHostAddress());
    final IPAddress upperAddr = addr.getUpper();
    System.out.println("Upper Address is " + upperAddr.toInetAddress().getHostAddress());
  }
}

