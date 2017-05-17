package scratch.misc;

import lombok.Getter;

import javax.inject.Inject;
import javax.inject.Named;

public class GuiceSrc {
  @Getter @Inject @Named("testValue") public int testValue = 100;
}
