package scratch;

import java.io.IOException;
import java.util.concurrent.TimeUnit;
import java.util.HashMap;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import lombok.Data;
import okhttp3.Call;
import okhttp3.Callback;
import okhttp3.Headers;
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
    private final static Logger log = LoggerFactory.getLogger(Scratch.class);
    private final OkHttpClient client;

    private OkHttpScratch(OkHttpClient client) {
        this.client = client;
    }

    public void shutdownClient() {
        this.client.dispatcher().executorService().shutdown();
    }

    public static class Builder {
        private int timeoutRd = 30;
        private int timeoutWr = 30;
        private int timeoutConn = 30;

        public void setTimeoutRd(int timeoutRd) { this.timeoutRd = timeoutRd; }
        public void setTimeoutWr(int timeoutWr) { this.timeoutWr = timeoutWr; }
        public void setTimeoutConn(int timeoutConn) { this.timeoutConn = timeoutConn; }

        public OkHttpScratch build() {
            OkHttpClient client = new OkHttpClient.Builder()
                    .writeTimeout(timeoutWr, TimeUnit.SECONDS)
                    .readTimeout(timeoutRd, TimeUnit.SECONDS)
                    .connectTimeout(timeoutConn, TimeUnit.SECONDS)
                    .build();
            return new OkHttpScratch(client);
        }
    }

    @Data
    public class ResponseData {
        private int code = 200;
        private String body = "";
        private HashMap<String, String> headers = new HashMap<String, String>();
    }

    private ResponseData toResponseData(Response rep) throws IOException {
        if (!rep.isSuccessful()) {
            log.debug("Invalid response");
            throw new IOException("Invalid response");
        }

        ResponseData ret = new ResponseData();
        ret.setCode(rep.code());
        ret.setBody(rep.body().string());
        final HashMap<String, String> retHdrs = ret.getHeaders();
        final Headers hdrs = rep.headers();
        hdrs.names().forEach(name -> retHdrs.put(name, hdrs.get(name)));
        return ret;
    }

    public interface AsyncCallback {
        void onFailure(IOException err);
        void onResponse(ResponseData response);
    }

    public void getAsync(String url, AsyncCallback act)  {
        Request req = new Request.Builder().url(url).build();
        client.newCall(req).enqueue(new Callback() {
            @Override
            public void onFailure(Call call, IOException err) { act.onFailure(err); }

            @Override
            public void onResponse(Call call, Response rep) {
                try {
                    act.onResponse(toResponseData(rep));
                }
                catch (IOException err) {
                    act.onFailure(err);
                }
            }
        });
    }

    public ResponseData get(String url) throws IOException {
        final Request req = new Request.Builder().url(url).build();
        final Response rep = client.newCall(req).execute();
        return toResponseData(rep);
    }
}
