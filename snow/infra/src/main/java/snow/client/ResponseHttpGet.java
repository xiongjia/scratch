package snow.client;

import java.util.Map;

public class ResponseHttpGet {
  private Map<String, Object> args;
  private Map<String, String> headers;
  private String origin;
  private String url;

  public Map<String, Object> getArgs() {
    return args;
  }

  public void setArgs(Map<String, Object> args) {
    this.args = args;
  }

  public Map<String, String> getHeaders() {
    return headers;
  }

  public void setHeaders(Map<String, String> headers) {
    this.headers = headers;
  }

  public String getOrigin() {
    return origin;
  }

  public void setOrigin(String origin) {
    this.origin = origin;
  }

  public String getUrl() {
    return url;
  }

  public void setUrl(String url) {
    this.url = url;
  }

  @Override
  public String toString() {
    return "ResponseHttpGet{" + "args=" + args
        + ", headers=" + headers
        + ", origin='" + origin + '\''
        + ", url='" + url + '\'' + '}';
  }
}
