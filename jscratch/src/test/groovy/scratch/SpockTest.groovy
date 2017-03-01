package scratch;

import spock.lang.Specification

class SpockTest extends Specification {
	def sayHello() {
        given: "given hello"
        def greeting = "Hello";

        expect: "expect hello"
        greeting == "Hello";
    }
}
