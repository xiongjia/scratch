package scratch.misc;

import com.google.inject.Inject;
import com.google.inject.name.Named;

import lombok.Getter;

public class GuiceSrc {
  @Getter @Inject @Named("testValue") public int testValue = 100;
}
