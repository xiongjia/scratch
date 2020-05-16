package hun.core;

import java.util.ArrayList;
import java.util.List;

public class Bean {
  public static final String SINGLETON = "singleton";
  public static final String PROTOTYPE = "prototype";

  private String name;
  private String className;
  private String scope = SINGLETON;
  private List<Property> properties = new ArrayList<Property>();

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getClassName() {
    return className;
  }

  public void setClassName(String className) {
    this.className = className;
  }

  public String getScope() {
    return scope;
  }

  public void setScope(String scope) {
    this.scope = scope;
  }

  public List<Property> getProperties() {
    return properties;
  }

  public void setProperties(List<Property> properties) {
    this.properties = properties;
  }
}
