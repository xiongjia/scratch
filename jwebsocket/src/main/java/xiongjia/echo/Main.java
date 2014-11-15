package xiongjia.echo;

import java.net.URI;
import java.net.URISyntaxException;
import java.util.logging.Logger;

import org.apache.commons.cli.Options;
import org.apache.commons.cli.CommandLineParser;
import org.apache.commons.cli.CommandLine;
import org.apache.commons.cli.BasicParser;
import org.apache.commons.cli.ParseException;
import org.apache.commons.cli.Option;
import org.apache.commons.cli.HelpFormatter;

import org.java_websocket.drafts.Draft_10;
import xiongjia.echo.EchoServer;
import xiongjia.echo.EchoClient;

public class Main {
    private final static Logger LOG = Logger.getLogger(Main.class.getName());
    private final static Options OPTS = createOptions();

    private final static int DEFAULT_SRV_PORT = 9005;
    private final static String DEFAULT_SRV_URI = "ws://localhost:9005";
    private final static String DEFAULT_SND_MSG = "Test Message";

    private static Options createOptions() {
        final Option optPort = new Option("port", true, "port number");
        optPort.setType(Number.class);
        final Option optTargetSrv = new Option("targetServ", true, "target server");
        optTargetSrv.setType(String.class);
        final Option optSndMsg = new Option("message", true, "Echo message");
        optSndMsg.setType(String.class);

        final Options opts = new Options();
        opts.addOption(optPort);
        opts.addOption(optTargetSrv);
        opts.addOption(optSndMsg);
        opts.addOption(new Option("help", "print this message"));
        opts.addOption(new Option("startServ", "start echo server"));
        return opts;
    }

    private static void printHelp() {
        final HelpFormatter formatter = new HelpFormatter();
        formatter.printHelp("jwebsocket", OPTS);
    }

    private static void startEchoServ(int port) {
        LOG.info("Creating Echo server. (Port = " + port + ")");
        try {
            final EchoServer srv = new EchoServer(port);
            srv.start();
        }
        catch (Exception ex) {
            LOG.warning("EchoServer err: " + ex.toString());
            ex.printStackTrace();
        }
    }

    private static void sendMsg(final String uri, final String msg) {
        LOG.info("Sending " + msg + " to " + uri);
        try {
            final EchoClient client = new EchoClient(new URI(uri), new Draft_10(), msg);
            client.connect();
        }
        catch (URISyntaxException uriErr) {
            LOG.warning("Invalid URI: " + uri + "; Err: " + uriErr.toString());
        }
    }

    public static void main(String[] args) {
        int servPort = DEFAULT_SRV_PORT;
        boolean startServ = false;
        String targetServ = DEFAULT_SRV_URI;
        String sndTxt = DEFAULT_SND_MSG;

        try {
            final CommandLineParser optParser = new BasicParser();
            CommandLine cmd = optParser.parse(OPTS, args);
            if (cmd.hasOption("help")) {
                printHelp();
                return;
            }
            if (cmd.hasOption("port")) {
                servPort = ((Number)cmd.getParsedOptionValue("port")).intValue();
            }
            if (cmd.hasOption("startServ")) {
                startServ = true;
            }
            if (cmd.hasOption("targetServ")) {
                targetServ = (String)cmd.getParsedOptionValue("targetServ");
            }
            if (cmd.hasOption("message")) {
                sndTxt = (String)cmd.getParsedOptionValue("message");
            }
        }
        catch (ParseException parseErr) {
            LOG.warning("Command line parse exception: " + parseErr.toString());
            printHelp();
            return;
        }

        if (startServ) {
            startEchoServ(servPort);
        }
        else {
            sendMsg(targetServ, sndTxt);
        }
    }
}
