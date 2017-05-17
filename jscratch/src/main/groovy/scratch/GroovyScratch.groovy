package scratch

import com.google.inject.AbstractModule
import com.google.inject.Guice
import com.google.inject.Injector
import com.google.inject.name.Names

import org.slf4j.Logger
import org.slf4j.LoggerFactory

import scratch.misc.GuiceSrc

class GroovyScratch {
  def get(String src) {
    return src
  }

  def getNum() {
    return 100
  }
}

class GroovyExtends extends Exception {
  def getErrorCode() {
    return 99
  }
}

class GroovyGuice {
  static final Logger log = LoggerFactory.getLogger(GroovyGuice.class)

  def test() {
    final Injector inj = Guice.createInjector(new AbstractModule() {
      @Override
      protected void configure() {
        bindConstant().annotatedWith(Names.named("testValue")).to(999)
      }
    })
    final GuiceSrc src = inj.getInstance(GuiceSrc.class)
    log.debug("testValue: {}", src.getTestValue());
  }
}
