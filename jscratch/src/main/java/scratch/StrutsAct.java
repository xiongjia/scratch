package scratch;

import lombok.Getter;
import lombok.Setter;

import org.apache.struts2.dispatcher.DefaultActionSupport;

public class StrutsAct extends DefaultActionSupport {
  private static final long serialVersionUID = 1L;
  @Getter @Setter private String data;

  @Override
  public String execute() throws Exception {
    this.data = "scratch data";
    return "ScratchView";
  }
}
