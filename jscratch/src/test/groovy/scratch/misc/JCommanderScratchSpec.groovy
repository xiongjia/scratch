package scratch.misc

import com.beust.jcommander.ParameterException;

import scratch.misc.JCommanderScratch.Commander
import scratch.misc.JCommanderScratch.JCommanderScratchBuilder

import spock.lang.Specification

class JCommanderScratchSpec extends Specification {
  def "simple argv parser"() {
    setup:
      def cmd = new JCommanderScratchBuilder().programName("JScratchTest").build()
      String[] argv = [
        "--verbose",
        "--number", "100",
        "--str", "JCommanderTest",
        "--listStr", "str1",
        "--listStr", "str2"
      ]

    when:
      def cmdArgs = cmd.parser(argv)

    then:
      cmdArgs.num.intValue() == 100
      cmdArgs.str == "JCommanderTest"
      cmdArgs.verbose == true
      cmdArgs.listStr.size() == 2
      cmdArgs.listStr.contains("str1")
      cmdArgs.listStr.contains("str2")
  }

  def "cmd usage"() {
    setup:
      def cmd = new JCommanderScratchBuilder().programName("UsageTest").build();

    when:
      def usage = cmd.usage()

    then:
      usage.matches("^Usage:\\sUsageTest(.|\\s)*") == true
  }

  def "bad option"() {
    setup:
      def cmd = new JCommanderScratchBuilder().programName("JScratchTest").build()
      String[] argv = [ "--badOption" ]

    when:
      cmd.parser(argv)

    then:
      ParameterException err = thrown()
      err.message.matches("^Unknown option:\\s--badOption(.|\\s)*") == true
  }
}
