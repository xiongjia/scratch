package snow.controller;

import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import snow.account.AccountService;
import snow.account.User;

@RestController
@RequestMapping("${snow.apiPrefix}/account")
public class Account {
  @Autowired
  private AccountService accountService;

  @GetMapping("users")
  public List<User> getAllUsers() {
    return accountService.getAllUsers();
  }
}
