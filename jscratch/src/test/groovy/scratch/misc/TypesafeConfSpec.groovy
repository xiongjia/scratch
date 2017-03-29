package scratch.misc;

import com.typesafe.config.ConfigException;

import spock.lang.Specification

class TypesafeConfSpec extends Specification {
  def "simple java type safe conf"() {
    expect:
      conf.getConfName() == confName
      conf.getStrVal("jscratch.data.value") == dataVal
      conf.getStrVal("jscratch.data.base-value") == baseVal

    where:
      conf               || confName        || baseVal      || dataVal
      new TypesafeConf() || "typesafe conf" || "base value" || "conf-value"
  }

  def "bad conf"() {
    setup:
      def conf = new TypesafeConf()

    when:
      conf.getStrVal("jscratch.bad-path")

    then:
      ConfigException err = thrown()
  }
}
