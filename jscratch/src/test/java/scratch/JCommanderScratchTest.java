package scratch;

import org.junit.Test;
import com.beust.jcommander.ParameterException;

public class JCommanderScratchTest {
    @Test
    public void cmdTest() {
    	JCommanderScratch cmd = new JCommanderScratch.Builder()
    			.withAppName("JScratchTest").build();

    	final String[] argv = { "-verbose", "3" };
    	try {
    		cmd.parser(argv);
    	}
    	catch (ParameterException e) { }
    }
}
