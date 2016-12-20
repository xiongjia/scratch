package scratch;

import static org.junit.Assert.*;
import org.junit.Test;
import com.typesafe.config.ConfigException;

public class TypesafeConfScratchTest {
    @Test
    public void loadRead() {
        TypesafeConfScratch conf = new TypesafeConfScratch();
        assertEquals(conf.getConfName(), "typesafe conf");
        assertEquals(conf.getStrVal("jscratch.data.value"), "conf-value");
        assertEquals(conf.getStrVal("jscratch.data.base-value"), "base value");
    }

    @Test(expected = ConfigException.class)
    public void badConf() {
        TypesafeConfScratch conf = new TypesafeConfScratch();
        conf.getStrVal("jscratch.bad-path");
    }
}
