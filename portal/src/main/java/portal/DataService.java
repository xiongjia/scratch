package portal;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class DataService {
  private static final Logger logger = LoggerFactory.getLogger(DataService.class);

  @Autowired
  private DataRepository dataRepository;

  public void saveData(Data data) {
    dataRepository.save(data);
  }

  public List<String> getAllData() {
    final List<String> result = new ArrayList<String>();
    dataRepository.findAll().forEach(data -> result.add(data.toString()));
    dataRepository.findByName("adding").forEach(data -> {
      logger.debug("SQL Query: {}", data.toString());
    });
    return result;
  }
}
