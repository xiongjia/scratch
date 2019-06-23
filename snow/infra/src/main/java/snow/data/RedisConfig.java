package snow.data;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisStandaloneConfiguration;
import org.springframework.data.redis.connection.jedis.JedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.GenericToStringSerializer;

@Configuration
public class RedisConfig {
  @Value("${spring.redis.host}")
  private String redisHost = "localhost";
  @Value("${spring.redis.port}")
  private int redisPort = 6379;

  @Bean
  JedisConnectionFactory jedisConnectionFactory() {
    final RedisStandaloneConfiguration redisStandaloneConfiguration =
        new RedisStandaloneConfiguration(redisHost, redisPort);
    return new JedisConnectionFactory(redisStandaloneConfiguration);
  }

  @Bean
  public RedisTemplate<String, Object> redisTemplate() {
    final RedisTemplate<String, Object> template = new RedisTemplate<String, Object>();
    final JedisConnectionFactory connectionFactory = jedisConnectionFactory();
    final String hostname = connectionFactory.getHostName();
    final int port = connectionFactory.getPort();

    template.setConnectionFactory(connectionFactory);
    template.setValueSerializer(new GenericToStringSerializer<Object>(Object.class));
    return template;
  }
}
