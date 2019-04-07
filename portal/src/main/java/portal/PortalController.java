package portal;

import io.swagger.annotations.ApiOperation;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.Arrays;
import java.util.List;

@RestController
public class PortalController {
  private static final Logger logger = LoggerFactory.getLogger(PortalController.class);

  //  @Autowired
  //  private DataService dataService;

  /** List all data. */
  @ApiOperation("scratch")
  @RequestMapping(value = "/api/list", method = RequestMethod.GET)
  public List<String> listAllData() {
    // final List<String> data = dataService.getAllData();
    // logger.debug("data: {}", Arrays.toString(new List[]{data}));
    // return data;
    return Arrays.asList("disable", "data", "service");
  }

  /** Add data */
  @ApiOperation("scratch")
  @RequestMapping(value = "/api/add", method = RequestMethod.GET)
  public void addData() {
    // final Data data = new Data("newName", 100);
    // dataService.saveData(data);
  }
}
