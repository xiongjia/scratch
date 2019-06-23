package snow.data;

import io.lettuce.core.dynamic.annotation.Param;
import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;
import org.springframework.web.bind.annotation.CrossOrigin;

import java.util.List;
import java.util.Optional;

@CrossOrigin
@RepositoryRestResource(collectionResourceRel = "datum", path = "datum")
public interface DataItemRepository extends PagingAndSortingRepository<DataItem, String> {
  Optional<DataItem> findById(@Param("id") String id);
}
