/**
 * My scala Tests
 */
package xiongjia.scratch

object Scratch {
  def main(args: Array[String]) = {
    /* test command line args */
    var msg = ""
    for (i <- 0 until args.length) {
      msg += (" [ " + args(i) + " ] ")
    }
    println(msg)

    /* test list */
    val listData: List[String] = List("apples", "oranges", "pears")
    for (listItem <- listData) println(listItem)
  }
}

