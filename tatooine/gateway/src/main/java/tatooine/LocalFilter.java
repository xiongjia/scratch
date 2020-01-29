package tatooine;

import org.springframework.cloud.gateway.filter.GatewayFilter;
import org.springframework.cloud.gateway.filter.factory.AbstractGatewayFilterFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class LocalFilter extends AbstractGatewayFilterFactory<LocalFilter.Config> {
  public LocalFilter() {
    super(Config.class);
  }

  @Bean
  public LocalFilter localFilterFactory() {
    return new LocalFilter();
  }

  @Override
  public GatewayFilter apply(Config config) {
    return (exchange, chain) -> {
      exchange.getRequest().getQueryParams();
      return chain.filter(exchange);
    };
  }

  public static class Config {

  }
}
