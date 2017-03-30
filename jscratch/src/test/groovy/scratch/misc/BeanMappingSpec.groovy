package scratch.misc

import spock.lang.Specification

class BeanMappingSpec extends Specification {
  def "source to target (ModelMapper tests)"() {
    expect:
      beans.getTarget().getTargetStr() == targetStr
      beans.getTarget().getTargetNum() == targetNum

    where:
      beans                    || targetStr || targetNum
      new Beans("src",    10)  || "src"     || 10
      new Beans("test2", 100)  || "test2"   || 100
  }

  static class Beans {
    BeanSource src
    BeanTarget target

    Beans(String str, int num) {
      src = new BeanSource(str, num)
      target = BeanMapping.toTarget(src)
    }
  }
}
