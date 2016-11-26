package scratch;

import java.util.logging.Logger;
import java.util.logging.Level;

public class Scratch {
    private final static Logger LOG = Logger.getLogger(Scratch.class.getName());

    public static void main(String[] args) {
        LOG.setLevel(Level.ALL);
        LOG.info("Hello JScratch !");
    }
}

