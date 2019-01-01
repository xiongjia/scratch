package owl;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.RequestMapping;

@FeignClient("elephant")
public interface Elephant {
  @RequestMapping("/")
  String getData();
}
