package scratch;

import static org.junit.Assert.*;
import org.junit.Test;
import com.beust.jcommander.ParameterException;
import scratch.JCommanderScratch.JCommanderScratchBuilder;
import scratch.JCommanderScratch.Commander;

public class JCommanderScratchTest {
    @Test
    public void cmdTypes() {
    	JCommanderScratch cmd = new JCommanderScratchBuilder()
    			.programName("JScratchTest").build();

    	final String[] argv = {
    			"--verbose",
    			"--number", "100",
    			"--str", "JCommanderTest",
    			"--listStr", "str1",
    			"--listStr", "str2" };
    	final Commander cmdArgs = cmd.parser(argv);
    	assertEquals(cmdArgs.num.intValue(), 100);
    	assertEquals(cmdArgs.str, "JCommanderTest");
    	assertTrue(cmdArgs.verbose);
    	assertEquals(cmdArgs.listStr.size(), 2);
    	assertTrue(cmdArgs.listStr.contains("str1"));
    	assertTrue(cmdArgs.listStr.contains("str2"));
    }

    @Test
    public void cmdUsage() {
    	JCommanderScratch cmd = new JCommanderScratchBuilder()
    			.programName("UsageTest").build();
    	final String usage = cmd.usage();
    	assertTrue(usage.matches("^Usage:\\sUsageTest(.|\\s)*"));
    }

    @Test
    public void cmdInvalidParam() {
    	final JCommanderScratch cmd = new JCommanderScratchBuilder().build();
    	final String[] argv = { "--badOption" };
    	try {
    		cmd.parser(argv);
    		fail("Should have thrown SomeException but did not!");
    	}
    	catch (ParameterException badOptErr) {
    		final String errMsg = badOptErr.getMessage();
    		final String pattern = "^Unknown option:\\s--badOption(.|\\s)*";
        	assertTrue(errMsg.matches(pattern));
    	}
    }
}
