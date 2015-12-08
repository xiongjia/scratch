/**
 *
 */
package xiongjia.scratch

object Scratch {
  def main(args: Array[String]) = {
    var msg = ""
    for (i <- 0 until args.length) {
      msg += (" [ " + args(i) + " ] ")
    }
    println(msg)
  }
}

