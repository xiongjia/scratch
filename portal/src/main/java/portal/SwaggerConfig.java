package portal;

import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import springfox.documentation.builders.ApiInfoBuilder;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger.web.UiConfigurationBuilder;
import springfox.documentation.swagger2.annotations.EnableSwagger2;
import springfox.documentation.swagger.web.UiConfiguration;


@Configuration
@ConditionalOnProperty(prefix = "portal.swagger", value = {"enable"}, havingValue = "true")
@EnableSwagger2
public class SwaggerConfig {
  @Bean
  public Docket apiConf() {
    final ApiInfo apiInfo = new ApiInfoBuilder().title("API Docs").build();
    return new Docket(DocumentationType.SWAGGER_2)
      .apiInfo(apiInfo)
      .select()
      .apis(RequestHandlerSelectors.basePackage("portal"))
      .paths(PathSelectors.any())
      .build();
  }

  @Bean
  UiConfiguration uiConfig() {
    return UiConfigurationBuilder.builder().build();
  }
}
