package hun.core;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import hun.aop.ProxyFactoryBean;
import org.apache.commons.beanutils.BeanUtils;

public class ClassPathXmlApplicationContext implements BeanFactory {
  private final Map<String, Bean> config;
  private Map<String, Object> context = new HashMap<>();

  public ClassPathXmlApplicationContext(String path) {
    config = ConfigurationManager.getBeanConfig(path);
    if (config == null) {
      return;
    }
    for (Map.Entry<String, Bean> entry : config.entrySet()) {
      final String beanName = entry.getKey();
      final Bean bean = entry.getValue();
      if (bean.getScope().equals(Bean.SINGLETON)) {
        Object beanObj = createBeanByConfig(bean);
        context.put(beanName, beanObj);
      }
    }
  }

  private Object createBeanByConfig(Bean bean) {
    try {
      final Class clazz = Class.forName(bean.getClassName());
      final Object beanObj = clazz.newInstance();

      final List<Property> properties = bean.getProperties();
      for (Property prop : properties) {
        final Map<String, Object> params = new HashMap<>();
        if (prop.getValue() != null) {
          params.put(prop.getName(), prop.getValue());
          BeanUtils.populate(beanObj, params);
        } else if (prop.getRef() != null) {
          Object ref = context.get(prop.getRef());
          if (ref == null) {
            ref = createBeanByConfig(config.get(prop.getRef()));
          }
          params.put(prop.getName(), ref);
          BeanUtils.populate(beanObj, params);
        }
      }
      if (clazz.equals(ProxyFactoryBean.class)) {
        final ProxyFactoryBean factoryBean = (ProxyFactoryBean) beanObj;
        return factoryBean.createProxy();
      }
      return beanObj;
    } catch (Exception error) {
      throw new RuntimeException("Creating " + bean.getClassName() + " failed");
    }
  }

  @Override
  public Object getBean(String name) {
    Bean bean = config.get(name);
    Object beanObj = null;
    if (bean.getScope().equals(Bean.SINGLETON)) {
      beanObj = context.get(name);
    } else if (bean.getScope().equals(Bean.PROTOTYPE)) {
      beanObj = createBeanByConfig(bean);
    }
    return beanObj;
  }
}
