package scratch;

import java.util.ArrayList;
import java.util.List;

import com.beust.jcommander.JCommander;
import com.beust.jcommander.Parameter;
import com.beust.jcommander.ParameterException;
import lombok.Builder;

/**
 * Scratch code for the JCommmander https://github.com/cbeust/jcommander
 *
 * NOTE:
 *   This sample is using lombok.builder.
 *   In Java IDE (e.g. Eclipse), you need install the Lombok first.
 *   (Run java -jar lombok.jar)
 *
 * @author lexiongjia@gmail.com
 */
@Builder
public class JCommanderScratch {
	private final Commander cmd = new Commander();
	private final JCommander jcmd = new JCommander(cmd);
	private String programName = "JScratch";

	public class Commander {
		@Parameter(names = {"--verbose", "-v" }, description = "verbose mode")
		public boolean verbose = false;

		@Parameter(names = {"--number", "-n" }, description = "Number parameter")
		public Integer num = 0;

		@Parameter(names = {"--str", "-s" }, description = "String parameter")
		public String str = "default";

		@Parameter(names = "--listStr", description = "List Strings")
		public List<String> listStr = new ArrayList<>();
	}

	public String usage() {
		jcmd.setProgramName(this.programName);
		StringBuilder usage = new StringBuilder();
		jcmd.usage(usage);
		return usage.toString();
	}

	public Commander parser(String... args) throws ParameterException {
		jcmd.parse(args);
		return this.cmd;
	}
}
