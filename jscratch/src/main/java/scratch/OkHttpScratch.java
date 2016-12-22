package scratch;

import java.io.IOException;

import lombok.Data;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

/**
 * Scratch code for the OkHttp https://github.com/square/okhttp
 * NOTE:
 *   In Java IDE (e.g. Eclipse), you need install the Lombok first.
 *   (Run java -jar lombok.jar)
 *
 * @author lexiongjia@gmail.com
 */
public class OkHttpScratch {
    private final OkHttpClient client = new OkHttpClient();

    @Data
    public class ResponseData {
        private int code = 200;
        private String body = "";
    }

    public ResponseData get(String url) throws IOException {
        final Request req = new Request.Builder().url(url).build();
        final Response rep = client.newCall(req).execute();
        final ResponseData ret = new ResponseData();
        ret.setCode(rep.code());
        ret.setBody(rep.body().string());
        return ret;
    }
}
