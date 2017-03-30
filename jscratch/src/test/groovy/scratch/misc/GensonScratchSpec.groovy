package scratch.misc

import com.owlike.genson.JsonBindingException

import scratch.misc.GensonScratch.JsonObj
import scratch.misc.GensonScratch.JsonRootObj

import spock.lang.Specification

class GensonScratchSpec extends Specification {
  def "string to json object"() {
    setup: "creating json from string"
      def jsonRoot = GensonScratch.buildRootFromStr(" { "
                                                  + "   \"str\": \"test\", "
                                                  + "   \"num\": 100, "
                                                  + "   \"jsonObj\": { \"str\": \"abc\"}"
                                                  + " }")

    when: "json values"
      def str = jsonRoot.getStr()
      def num = jsonRoot.getNum()
      def subStr = jsonRoot.getJsonObj().getStr()

    then:
      str == "test"
      num == 100
      subStr == "abc"
  }

  def "json object to string"() {
    setup: "creating json object"
      def jsonRoot =  new JsonRootObj()
      jsonRoot.setNum(99)
      jsonRoot.setStr("rootStr")
      jsonRoot.setJsonObj(new JsonObj())
      jsonRoot.getJsonObj().setStr("subStr")

    when: "json values"
      def str = jsonRoot.getStr()
      def num = jsonRoot.getNum()
      def subStr = jsonRoot.getJsonObj().getStr()

    then:
      str == "rootStr"
      num == 99
      subStr == "subStr"
  }

  def "invalid json string"() {
    when:
      GensonScratch.buildRootFromStr("{ \"str\" bad json fmt }")

    then:
      JsonBindingException err = thrown()
  }
}
