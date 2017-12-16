package scratch.misc;

import java.util.concurrent.TimeUnit;

import org.openjdk.jmh.annotations.Benchmark;
import org.openjdk.jmh.annotations.BenchmarkMode;
import org.openjdk.jmh.annotations.Level;
import org.openjdk.jmh.annotations.Mode;
import org.openjdk.jmh.annotations.OutputTimeUnit;
import org.openjdk.jmh.annotations.Scope;
import org.openjdk.jmh.annotations.Setup;
import org.openjdk.jmh.annotations.State;
import org.openjdk.jmh.annotations.TearDown;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Microbenchmark {
  private static final Logger log = LoggerFactory.getLogger(Microbenchmark.class);

  @State(Scope.Thread)
  public static class BenchmarkState {
    public int ctx = 1;
    public int sum = 0;

    @Setup(Level.Trial)
    public void doSetup() {
      log.debug("benchmark setup: sum = {}, ctx = {}", this.sum, this.ctx);
    }

    @TearDown(Level.Trial)
    public void doTearDown() {
      log.debug("benchmark teardown: sum = {}, ctx = {}", this.sum, this.ctx);
    }
  }

  /** test throughput. */
  @Benchmark
  @BenchmarkMode(Mode.Throughput)
  @OutputTimeUnit(TimeUnit.MICROSECONDS)
  public void measureThroughput(BenchmarkState state) throws InterruptedException {
    state.sum += state.ctx++;
    TimeUnit.MILLISECONDS.sleep(50);
  }

  /** test AverageTime. */
  @Benchmark
  @BenchmarkMode(Mode.AverageTime)
  @OutputTimeUnit(TimeUnit.MICROSECONDS)
  public void measureAvgTime(BenchmarkState state) throws InterruptedException {
    state.sum += state.ctx++;
    TimeUnit.MILLISECONDS.sleep(100);
  }

  /** run benchmark. */
  public static void run() throws RunnerException {
    final Options opt = new OptionsBuilder()
        .include(Microbenchmark.class.getSimpleName())
        .warmupIterations(2)
        .measurementIterations(5)
        .forks(1)
        .build();
    new Runner(opt).run();
  }
}
