package portal;

import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface DataRepository extends CrudRepository<Data, Long> {

  @Query(value = "SELECT * FROM DATA WHERE name = ?1",
         countQuery = "SELECT count(*) FROM DATA WHERE name = ?1",
         nativeQuery = true)
  List<Data> findByName(String name);
}
