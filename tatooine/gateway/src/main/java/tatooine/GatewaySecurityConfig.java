package tatooine;

import org.springframework.context.annotation.Bean;
import org.springframework.security.config.annotation.method.configuration.EnableReactiveMethodSecurity;
import org.springframework.security.config.annotation.web.reactive.EnableWebFluxSecurity;
import org.springframework.security.config.web.server.ServerHttpSecurity;
import org.springframework.security.web.server.SecurityWebFilterChain;

@EnableWebFluxSecurity
@EnableReactiveMethodSecurity
public class GatewaySecurityConfig {
  @Bean
  public SecurityWebFilterChain filterChain(ServerHttpSecurity http) {
    http
      .authorizeExchange()
      .pathMatchers("/api/v1/**").permitAll();

    return http.build();
  }
}
