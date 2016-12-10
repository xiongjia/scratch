package scratch;

import com.beust.jcommander.JCommander;
import com.beust.jcommander.Parameter;

/**
 * Scratch code for the JCommmander https://github.com/cbeust/jcommander
 *
 * @author lexiongjia@gmail.com
 */
public class JCommanderScratch {
	private final JCommander jcmd = new JCommander(new Commander());

	public class Commander {
		@Parameter(names = {"-verbose", "-v" }, description = "verbose mode")
		public boolean verbose = false;
	}

	public JCommanderScratch(Builder builder) {
		jcmd.setProgramName(builder.programName);
	}

	public void parser(String... args) {
		jcmd.parse(args);
	}

	public static class Builder {
		public String programName = "JScratch";

		public Builder withAppName(String programName) {
			this.programName = programName;
			return this;
		}

		public JCommanderScratch build() {
			return new JCommanderScratch(this);
		}
	 }
}
