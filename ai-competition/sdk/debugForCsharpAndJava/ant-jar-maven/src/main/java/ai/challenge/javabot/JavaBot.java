package ai.challenge.javabot;

import java.io.IOException;
import java.io.PrintStream;
import java.net.ServerSocket;
import java.net.Socket;

/**
 * Starter bot implementation.
 */
public class JavaBot extends Bot {
    /**
     * Main method executed by the game engine for starting the bot.
     * 
     * @param args command line arguments
     * 
     * @throws IOException if an I/O error occurs
     */
    public static void main(String[] args) throws IOException {
        if(args.length > 0){
            int port=Integer.parseInt(args[0]);
            try(final ServerSocket serverSocket=new ServerSocket(port)){
                do {
                    final Socket socket = serverSocket.accept();
                    Logger.log("|||:new connection build.");
                    new Thread(new Runnable() {
                        public void run() {
                            try {
                                JavaBot bot = new JavaBot();
                                bot.in = socket.getInputStream();
                                bot.out = new PrintStream(socket.getOutputStream());
                                bot.readSystemInput();
                            } catch (IOException e) {
                                e.printStackTrace();
                            } finally {
                                try {
                                    socket.close();
                                } catch (IOException e) {
                                    //
                                }
                            }
                        }
                    }).run();
                }while (true);
            }

        }
        JavaBot bot=new JavaBot();
        bot.in=System.in;
        bot.out=System.out;
        bot.readSystemInput();
    }
    
    /**
     * For every ant check every direction in fixed order (N, E, S, W) and move it if the tile is
     * passable.
     */
    @Override
    public void doTurn() {
        Ants ants = getAnts();
        for (Tile myAnt : ants.getMyAnts()) {
            for (Aim direction : Aim.values()) {
                if (ants.getIlk(myAnt, direction).isPassable()) {
                    ants.issueOrder(myAnt, direction);
                    break;
                }
            }
        }
    }
}
