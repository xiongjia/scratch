using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading;

namespace Ants {

	class MyBot : Bot {

        #region Make this bot use socket communicate with SockerForward.jar through port 18889 ʹ��SocketServer�ķ�ʽ����18889�˿ڣ���SocketForwardͨ���Ա��ڵ���
        private static Socket serverSocket;
        private static Thread listenThread;

        MyBot(Socket sock)
        {
            serverSocket = sock;
        }

        public void IssueOrder(Location loc, Direction direction)
        {
            var buffers = Encoding.UTF8.GetBytes(String.Format("o {0} {1} {2}\n", loc.Row, loc.Col, direction.ToChar()));
            try
            {
                serverSocket.Send(buffers);
            }
            catch (Exception ex)
            {
                System.Console.Out.WriteLine(ex.ToString());
            }

            System.Console.Out.WriteLine("o {0} {1} {2}\n", loc.Row, loc.Col, direction.ToChar());
        }

        public static void Main(string[] args)
        {

            IPAddress ip = IPAddress.Parse("127.0.0.1");
            IPEndPoint point = new IPEndPoint(ip, 18889);
            serverSocket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
            serverSocket.Bind(point);

            serverSocket.Listen(10);



            listenThread = new Thread(SocketListen);
            listenThread.Start();

        }

        static void SocketListen()
        {
            while (true)
            {
                //�ȴ����Ӳ��Ҵ���һ������ͨѶ��socket
                var send = serverSocket.Accept();
                //��ȡ���ӵ�IP��ַ
                var sendIpoint = send.RemoteEndPoint.ToString();
                Console.WriteLine($"{sendIpoint}Connection");
                //����һ�����̲߳�ͣ������Ϣ
                Thread thread = new Thread(Recive);
                thread.IsBackground = true;
                thread.Start(send);
            }
        }

        static void Recive(object clientSocket)
        {
            Socket send = (Socket)clientSocket;
            MyBot mybot = new MyBot(send);
            Ants antsEntry = new Ants();
            antsEntry.setClientSocket(send);
            antsEntry.PlayGame(mybot);
        }
        #endregion

        #region Where to implement game logic, same as before ��Ϸ�߼������ĵط�����NoDebug��һ��
        // DoTurn is run once per turn
        public override void DoTurn (IGameState state) {

			// loop through all my ants and try to give them orders
			foreach (Ant ant in state.MyAnts) {
				
				// try all the directions
				foreach (Direction direction in Ants.Aim.Keys) {

					// GetDestination will wrap around the map properly
					// and give us a new location
					Location newLoc = state.GetDestination(ant, direction);

					// GetIsPassable returns true if the location is land
					if (state.GetIsPassable(newLoc)) {
						IssueOrder(ant, direction);
						// stop now, don't give 1 and multiple orders
						break;
					}
				}
				
				// check if we have time left to calculate more orders
				if (state.TimeRemaining < 10) break;
			}
			
		}
        #endregion

    }
	
}