package hun.core;

import com.google.common.base.Strings;
import java.io.InputStream;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import org.dom4j.Document;
import org.dom4j.DocumentException;
import org.dom4j.Element;
import org.dom4j.Node;
import org.dom4j.io.SAXReader;

public class ConfigurationManager {

  private static void parseProp(Bean bean, List<Element>  propNodes) {
    if (propNodes == null) {
      return;
    }
    for (final Element element : propNodes) {
      final Property prop = new Property();
      prop.setName(element.attributeValue("name"));
      prop.setValue(element.attributeValue("value"));
      prop.setRef(element.attributeValue("ref"));
      bean.getProperties().add(prop);
    }
  }

  public static Map<String, Bean> getBeanConfig(String path) {
    final Map<String, Bean> result = new HashMap<>();

    final SAXReader reader = new SAXReader();
    final InputStream is = ConfigurationManager.class.getResourceAsStream(path);
    Document doc = null;
    try {
      doc = reader.read(is);
    } catch (DocumentException e) {
      throw new RuntimeException("Cannot load config.");
    }
    final List<Node> beanNodes = doc.selectNodes("//bean");
    for (final Node beanNode : beanNodes) {
      final Element element = (Element)beanNode;
      final Bean bean = new Bean();
      bean.setName(element.attributeValue("name"));
      bean.setClassName(element.attributeValue("class"));
      final String scope = element.attributeValue("scope");
      bean.setScope(Strings.isNullOrEmpty(scope) ? Bean.SINGLETON : scope);
      parseProp(bean, element.elements("property"));
      result.put(bean.getName(), bean);
    }
    return result;
  }
}
