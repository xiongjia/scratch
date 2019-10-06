package snow.controller;

import com.google.common.net.HttpHeaders;
import com.opencsv.CSVWriter;
import com.opencsv.bean.StatefulBeanToCsv;
import com.opencsv.bean.StatefulBeanToCsvBuilder;
import com.opencsv.exceptions.CsvException;
import java.io.IOException;
import java.util.List;
import javax.servlet.http.HttpServletResponse;
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

  @GetMapping("export")
  public void export(HttpServletResponse response) throws IOException, CsvException {
    response.setContentType("text/csv");
    response.setHeader(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=users.csv");
    final StatefulBeanToCsv<User> csvWriter = new StatefulBeanToCsvBuilder<User>(response.getWriter())
        .withQuotechar(CSVWriter.NO_QUOTE_CHARACTER)
        .withSeparator(CSVWriter.DEFAULT_SEPARATOR)
        .withOrderedResults(false)
        .build();
    csvWriter.write(accountService.getAllUsers());
  }
}
