package scratch;

import com.owlike.genson.Genson;
import com.owlike.genson.GensonBuilder;

import lombok.Data;
import lombok.RequiredArgsConstructor;

/**
 * Scratch code for the Genson https://github.com/owlike/genson
 * NOTE:
 *   This sample is using lombok.builder.
 *   In Java IDE (e.g. Eclipse), you need install the Lombok first.
 *   (Run java -jar lombok.jar)
 *
 * @author lexiongjia@gmail.com
 */
class GensonScratch {
  private static final Genson genson = new GensonBuilder()
      .useClassMetadata(false).useRuntimeType(true).create();

  @Data
  @RequiredArgsConstructor
  public static class JsonObj {
    private String str;
  }

  @Data
  @RequiredArgsConstructor
  public static class JsonRootObj {
    private String str;
    private int num;
    private JsonObj jsonObj;
  }

  public static JsonRootObj buildRootFromStr(String json) {
    return genson.deserialize(json, JsonRootObj.class);
  }

  public static String buildJsonStr(JsonRootObj obj) {
    return genson.serialize(obj);
  }
}
