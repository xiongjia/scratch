package jinstrument;

import java.lang.instrument.Instrumentation;

public class Instrument {
  public static void premain(String agentArgs, Instrumentation instrument) {
    System.out.println("Agent main");
    instrument.addTransformer(new TestTransformer());
  }
}
