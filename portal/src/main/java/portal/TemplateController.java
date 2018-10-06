package portal;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller
public class TemplateController {

  /** Thymeleaf testing */
  @RequestMapping(value = "/", method = RequestMethod.GET)
  public String mainPage(Model model) {
    model.addAttribute("message", "test message");
    return "main-page";
  }
}
