package scratch.database;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "users")
public class UserEntity {
  @Id
  @Column(name = "id")
  @Getter @Setter private Long id;

  @Column(name = "name")
  @Getter @Setter private String name;

  public UserEntity(Long id, String name) {
    this.id = id;
    this.name = name;
  }
}
