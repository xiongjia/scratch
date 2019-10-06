package snow.debugger;

import com.google.common.base.Strings;
import java.util.Arrays;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;
import org.apache.commons.lang3.StringUtils;
import org.springframework.core.env.AbstractEnvironment;
import org.springframework.core.env.EnumerablePropertySource;
import org.springframework.core.env.Environment;
import org.springframework.core.env.MutablePropertySources;
import org.springframework.stereotype.Service;
import snow.misc.ContextProvider;

@Service
public class DumpProperties {
  public String dumpAll() {
    final Environment env = ContextProvider.getContext().getEnvironment();
    final MutablePropertySources sources = ((AbstractEnvironment) env).getPropertySources();
    return StreamSupport.stream(sources.spliterator(), false)
        .filter(ps -> ps instanceof EnumerablePropertySource)
        .map(ps -> ((EnumerablePropertySource) ps).getPropertyNames())
        .flatMap(Arrays::stream)
        .distinct()
        .filter(prop -> !StringUtils.containsIgnoreCase(prop, "password"))
        .map(prop -> String.format("%s = %s", prop, Strings.nullToEmpty(env.getProperty(prop))))
        .collect(Collectors.joining("\n"));
  }
}
