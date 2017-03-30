package scratch.server;

import lombok.Data;

import java.io.Serializable;

@Data
public class ServletBean implements Serializable {
  private static final long serialVersionUID = 1L;
  private String testData;
}
